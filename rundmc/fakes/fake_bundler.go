// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/goci"
	"github.com/cloudfoundry-incubator/guardian/gardener"
	"github.com/cloudfoundry-incubator/guardian/rundmc"
)

type FakeBundler struct {
	BundleStub        func(spec gardener.DesiredContainerSpec) *goci.Bndl
	bundleMutex       sync.RWMutex
	bundleArgsForCall []struct {
		spec gardener.DesiredContainerSpec
	}
	bundleReturns struct {
		result1 *goci.Bndl
	}
}

func (fake *FakeBundler) Bundle(spec gardener.DesiredContainerSpec) *goci.Bndl {
	fake.bundleMutex.Lock()
	fake.bundleArgsForCall = append(fake.bundleArgsForCall, struct {
		spec gardener.DesiredContainerSpec
	}{spec})
	fake.bundleMutex.Unlock()
	if fake.BundleStub != nil {
		return fake.BundleStub(spec)
	} else {
		return fake.bundleReturns.result1
	}
}

func (fake *FakeBundler) BundleCallCount() int {
	fake.bundleMutex.RLock()
	defer fake.bundleMutex.RUnlock()
	return len(fake.bundleArgsForCall)
}

func (fake *FakeBundler) BundleArgsForCall(i int) gardener.DesiredContainerSpec {
	fake.bundleMutex.RLock()
	defer fake.bundleMutex.RUnlock()
	return fake.bundleArgsForCall[i].spec
}

func (fake *FakeBundler) BundleReturns(result1 *goci.Bndl) {
	fake.BundleStub = nil
	fake.bundleReturns = struct {
		result1 *goci.Bndl
	}{result1}
}

var _ rundmc.Bundler = new(FakeBundler)