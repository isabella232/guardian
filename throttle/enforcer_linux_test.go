package throttle_test

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	gardencgroups "code.cloudfoundry.org/guardian/rundmc/cgroups"
	"code.cloudfoundry.org/guardian/throttle"
	"code.cloudfoundry.org/lager/lagertest"
	uuid "github.com/nu7hatch/gouuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opencontainers/runc/libcontainer/cgroups"
)

var _ = Describe("Enforcer", func() {
	var (
		logger *lagertest.TestLogger
	)

	BeforeEach(func() {
		logger = lagertest.NewTestLogger("container-metrics-test")
	})

	Describe("Punish", func() {

		var (
			handle        string
			punishErr     error
			cgroupRoot    string
			cpuCgroupPath string
			command       *exec.Cmd
		)

		BeforeEach(func() {
			uuid, err := uuid.NewV4()
			Expect(err).NotTo(HaveOccurred())
			handle = uuid.String()

			cgroupRoot, err = ioutil.TempDir("", "cgroups")
			Expect(err).NotTo(HaveOccurred())

			mountCPUcgroup(cgroupRoot)
			cpuCgroupPath = filepath.Join(cgroupRoot, "cpu")
		})

		JustBeforeEach(func() {
			enforcer := throttle.NewEnforcer(cpuCgroupPath)
			punishErr = enforcer.Punish(logger, handle)
		})

		AfterEach(func() {
			Expect(command.Process.Kill()).To(Succeed())
			_, err := command.Process.Wait()
			Expect(err).NotTo(HaveOccurred())
			umountCgroups(cgroupRoot)
		})

		Context("containers that have been created after cpu throttling enablement", func() {
			var (
				goodCgroup          string
				goodContainerCgroup string
				badCgroup           string
				badContainerCgroup  string
			)

			BeforeEach(func() {
				goodCgroup = filepath.Join(cpuCgroupPath, gardencgroups.GoodCgroupName)
				goodContainerCgroup = filepath.Join(goodCgroup, handle)
				Expect(os.MkdirAll(goodContainerCgroup, 0755)).To(Succeed())
				Expect(ioutil.WriteFile(filepath.Join(goodContainerCgroup, "cpu.shares"), []byte("3456"), 0755)).To(Succeed())

				badCgroup = filepath.Join(cpuCgroupPath, gardencgroups.BadCgroupName)
				badContainerCgroup = filepath.Join(badCgroup, handle)
				Expect(os.MkdirAll(badContainerCgroup, 0755)).To(Succeed())

				command = exec.Command("sleep", "360")
				Expect(command.Start()).To(Succeed())

				Expect(cgroups.WriteCgroupProc(goodContainerCgroup, command.Process.Pid)).To(Succeed())
			})

			It("moves the process to the bad cgroup", func() {
				Expect(punishErr).NotTo(HaveOccurred())

				pids, err := cgroups.GetPids(goodContainerCgroup)
				Expect(err).NotTo(HaveOccurred())
				Expect(pids).To(BeEmpty())

				pids, err = cgroups.GetPids(badContainerCgroup)
				Expect(err).NotTo(HaveOccurred())
				Expect(pids).To(ContainElement(command.Process.Pid))
			})

			It("has low, non-default cpu.shares on the bad cgroup", func() {
				goodShares := readCPUShares(goodCgroup)
				badShares := readCPUShares(badCgroup)
				Expect(goodShares).To(BeNumerically(">", badShares))
			})

			It("copies CPU shares to the bad container cgroup", func() {
				badContainerShares := readCPUShares(badContainerCgroup)
				Expect(badContainerShares).To(Equal(3456))
			})
		})

		Context("containers that have been created before throttling feature was enabled", func() {
			var (
				containerCgroup string
			)

			BeforeEach(func() {
				containerCgroup = filepath.Join(cpuCgroupPath, handle)
				Expect(os.MkdirAll(containerCgroup, 0755)).To(Succeed())

				command = exec.Command("sleep", "360")
				Expect(command.Start()).To(Succeed())

				Expect(cgroups.WriteCgroupProc(containerCgroup, command.Process.Pid)).To(Succeed())
			})

			It("does not move the container to another cgroup", func() {
				Expect(punishErr).NotTo(HaveOccurred())
				pids, err := cgroups.GetPids(containerCgroup)
				Expect(err).NotTo(HaveOccurred())
				Expect(pids).To(ContainElement(command.Process.Pid))
			})
		})
	})
})

func readCPUShares(cgroupPath string) int {
	shareBytes, err := ioutil.ReadFile(filepath.Join(cgroupPath, "cpu.shares"))
	Expect(err).NotTo(HaveOccurred())
	shares, err := strconv.Atoi(strings.TrimSpace(string(shareBytes)))
	Expect(err).NotTo(HaveOccurred())
	return shares
}

func mountCPUcgroup(cgroupRoot string) {
	Expect(syscall.Mount("cgroup", cgroupRoot, "tmpfs", uintptr(0), "mode=0755")).To(Succeed())

	cpuCgroup := filepath.Join(cgroupRoot, "cpu")
	Expect(os.MkdirAll(cpuCgroup, 0755)).To(Succeed())

	Expect(syscall.Mount("cgroup", cpuCgroup, "cgroup", uintptr(0), "cpu,cpuacct")).To(Succeed())
}

func umountCgroups(cgroupRoot string) {
	cpuCgroup := filepath.Join(cgroupRoot, "cpu")
	Expect(os.RemoveAll(filepath.Join(cpuCgroup, gardencgroups.Garden))).To(Succeed())
	Expect(syscall.Unmount(cpuCgroup, 0)).To(Succeed())
	Expect(syscall.Unmount(cgroupRoot, 0)).To(Succeed())
}