// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/guardian/rundmc/runrunc"
)

type FakeUserLookupper struct {
	LookupStub        func(rootFsPath string, user string) (uint32, uint32, error)
	lookupMutex       sync.RWMutex
	lookupArgsForCall []struct {
		rootFsPath string
		user       string
	}
	lookupReturns struct {
		result1 uint32
		result2 uint32
		result3 error
	}
}

func (fake *FakeUserLookupper) Lookup(rootFsPath string, user string) (uint32, uint32, error) {
	fake.lookupMutex.Lock()
	fake.lookupArgsForCall = append(fake.lookupArgsForCall, struct {
		rootFsPath string
		user       string
	}{rootFsPath, user})
	fake.lookupMutex.Unlock()
	if fake.LookupStub != nil {
		return fake.LookupStub(rootFsPath, user)
	} else {
		return fake.lookupReturns.result1, fake.lookupReturns.result2, fake.lookupReturns.result3
	}
}

func (fake *FakeUserLookupper) LookupCallCount() int {
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	return len(fake.lookupArgsForCall)
}

func (fake *FakeUserLookupper) LookupArgsForCall(i int) (string, string) {
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	return fake.lookupArgsForCall[i].rootFsPath, fake.lookupArgsForCall[i].user
}

func (fake *FakeUserLookupper) LookupReturns(result1 uint32, result2 uint32, result3 error) {
	fake.LookupStub = nil
	fake.lookupReturns = struct {
		result1 uint32
		result2 uint32
		result3 error
	}{result1, result2, result3}
}

var _ runrunc.UserLookupper = new(FakeUserLookupper)
