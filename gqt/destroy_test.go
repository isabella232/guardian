package gqt_test

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/cloudfoundry-incubator/garden"
	"github.com/cloudfoundry-incubator/guardian/gqt/runner"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Destroying a Container", func() {
	var (
		client    *runner.RunningGarden
		container garden.Container
	)

	BeforeEach(func() {
		client = startGarden()
	})

	AfterEach(func() {
		Expect(client.DestroyAndStop()).To(Succeed())
	})

	Context("when destroying the container", func() {
		var (
			process         garden.Process
			initProcPid     int
			containerRootfs string
		)

		BeforeEach(func() {
			var err error

			container, err = client.Create(garden.ContainerSpec{})
			Expect(err).NotTo(HaveOccurred())

			initProcPid = initProcessPID(container.Handle())

			process, err = container.Run(garden.ProcessSpec{
				Path: "/bin/sh",
				Args: []string{
					"-c", "read x",
				},
			}, ginkgoIO)
			Expect(err).NotTo(HaveOccurred())

			info, err := container.Info()
			Expect(err).NotTo(HaveOccurred())
			containerRootfs = info.ContainerPath
		})

		JustBeforeEach(func() {
			Expect(client.Destroy(container.Handle())).To(Succeed())
		})

		It("should kill the containers init process", func() {
			var killExitCode = func() int {
				sess, err := gexec.Start(exec.Command("kill", "-0", fmt.Sprintf("%d", initProcPid)), GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				sess.Wait(1 * time.Second)
				return sess.ExitCode()
			}

			Eventually(killExitCode).Should(Equal(1))
		})

		It("should destroy the container's depot directory", func() {
			Expect(filepath.Join(client.DepotDir, container.Handle())).NotTo(BeAnExistingFile())
		})

		It("should destroy the container rootfs", func() {
			Expect(containerRootfs).NotTo(BeAnExistingFile())
		})
	})

	Context("cleaning up networking after destroy", func() {
		var (
			contIfaceName     string
			contHandle        string
			existingContainer garden.Container
		)

		BeforeEach(func() {
			var err error

			container, err = client.Create(garden.ContainerSpec{
				Network: "177.100.10.30/24",
			})
			Expect(err).NotTo(HaveOccurred())
			contIfaceName = ethInterfaceName(container)
			contHandle = container.Handle()

			existingContainer, err = client.Create(garden.ContainerSpec{
				Network: "168.100.20.10/24",
			})
			Expect(err).NotTo(HaveOccurred())
		})

		AfterEach(func() {
			err := client.Destroy(existingContainer.Handle())
			if err != nil {
				Expect(err).To(MatchError(garden.ContainerNotFoundError{Handle: existingContainer.Handle()}))
			}
		})

		var itCleansUpTheNetwork = func() {
			It("should remove iptable entries", func() {
				out, err := exec.Command("iptables", "-w", "-S", "-t", "filter").CombinedOutput()
				Expect(err).NotTo(HaveOccurred())
				Expect(string(out)).NotTo(MatchRegexp("w-%d-instance.* 177.100.10.0/24", GinkgoParallelNode()))
				Expect(string(out)).To(ContainSubstring("168.100.20.0/24"))
			})

			It("should remove virtual ethernet cards", func() {
				ifconfigExits := func() int {
					session, err := gexec.Start(exec.Command("ifconfig", contIfaceName), GinkgoWriter, GinkgoWriter)
					Expect(err).NotTo(HaveOccurred())

					return session.Wait().ExitCode()
				}
				Eventually(ifconfigExits).ShouldNot(Equal(0))

				ifaceName := ethInterfaceName(existingContainer)
				session, err := gexec.Start(exec.Command("ifconfig", ifaceName), GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				Eventually(session).Should(gexec.Exit(0))
			})

			It("should remove the network bridge", func() {
				Expect(client.Destroy(existingContainer.Handle())).To(Succeed())

				// ifconfig can be flakey when it runs while parallel tests are destroying
				// interfaces, so run it a few times to minimize chance of problems
				var session *gexec.Session
				for i := 0; i < 30; i++ {
					var err error
					session, err = gexec.Start(
						exec.Command("ifconfig"),
						GinkgoWriter, GinkgoWriter,
					)
					Expect(err).NotTo(HaveOccurred())

					session.Wait()
					if session.ExitCode() == 0 {
						break
					}

					time.Sleep(1 * time.Second)
				}

				Expect(session.ExitCode()).To(Equal(0))
				Expect(session).NotTo(gbytes.Say("w%dbrdg-", GinkgoParallelNode()))
			})
		}

		Context("after Destroy is called", func() {
			JustBeforeEach(func() {
				Expect(client.Destroy(container.Handle())).To(Succeed())
			})

			itCleansUpTheNetwork()
		})

		Context("when runc kill is called directly", func() {
			JustBeforeEach(func() {
				sess, err := gexec.Start(exec.Command("runc", "kill", container.Handle(), "KILL"), GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				Eventually(sess).Should(gexec.Exit(0))

				Eventually(func() int {
					sess, err := gexec.Start(exec.Command("runc", "state", container.Handle()), GinkgoWriter, GinkgoWriter)
					Expect(err).NotTo(HaveOccurred())
					sess.Wait()

					return sess.ExitCode()
				}).ShouldNot(Equal(0))

				// runc deletes state BEFORE running post-stop hooks, so we have no choice
				// but to wait a bit longer for the deletes to have happened
				time.Sleep(3 * time.Second)
			})

			itCleansUpTheNetwork()
		})
	})

	It("doesn't leak goroutines", func() {
		handle := fmt.Sprintf("goroutine-leak-test-%d", GinkgoParallelNode())

		numGoroutinesBefore, err := client.NumGoroutines()
		Expect(err).NotTo(HaveOccurred())

		_, err = client.Create(garden.ContainerSpec{
			Handle: handle,
		})
		Expect(err).NotTo(HaveOccurred())

		client.Destroy(handle)

		Eventually(func() int {
			numGoroutinesAfter, err := client.NumGoroutines()
			Expect(err).NotTo(HaveOccurred())
			return numGoroutinesAfter
		}).Should(Equal(numGoroutinesBefore))
	})
})

func ethInterfaceName(container garden.Container) string {
	buffer := gbytes.NewBuffer()
	proc, err := container.Run(
		garden.ProcessSpec{
			Path: "sh",
			Args: []string{"-c", "ifconfig | grep 'Ethernet' | cut -f 1 -d ' '"},
			User: "root",
		},
		garden.ProcessIO{
			Stdout: buffer,
			Stderr: GinkgoWriter,
		},
	)
	Expect(err).NotTo(HaveOccurred())
	Expect(proc.Wait()).To(Equal(0))

	contIfaceName := string(buffer.Contents()) // g3-abc-1

	return contIfaceName[:len(contIfaceName)-2] + "0" // g3-abc-0
}
