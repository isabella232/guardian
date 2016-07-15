// This file was generated by counterfeiter
package runruncfakes

import (
	"os"
	"sync"

	"code.cloudfoundry.org/guardian/rundmc/runrunc"
)

type FakeMkdirer struct {
	MkdirAsStub        func(rootfsPath string, uid, gid int, mode os.FileMode, recreate bool, path ...string) error
	mkdirAsMutex       sync.RWMutex
	mkdirAsArgsForCall []struct {
		rootfsPath string
		uid        int
		gid        int
		mode       os.FileMode
		recreate   bool
		path       []string
	}
	mkdirAsReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMkdirer) MkdirAs(rootfsPath string, uid int, gid int, mode os.FileMode, recreate bool, path ...string) error {
	fake.mkdirAsMutex.Lock()
	fake.mkdirAsArgsForCall = append(fake.mkdirAsArgsForCall, struct {
		rootfsPath string
		uid        int
		gid        int
		mode       os.FileMode
		recreate   bool
		path       []string
	}{rootfsPath, uid, gid, mode, recreate, path})
	fake.recordInvocation("MkdirAs", []interface{}{rootfsPath, uid, gid, mode, recreate, path})
	fake.mkdirAsMutex.Unlock()
	if fake.MkdirAsStub != nil {
		return fake.MkdirAsStub(rootfsPath, uid, gid, mode, recreate, path...)
	} else {
		return fake.mkdirAsReturns.result1
	}
}

func (fake *FakeMkdirer) MkdirAsCallCount() int {
	fake.mkdirAsMutex.RLock()
	defer fake.mkdirAsMutex.RUnlock()
	return len(fake.mkdirAsArgsForCall)
}

func (fake *FakeMkdirer) MkdirAsArgsForCall(i int) (string, int, int, os.FileMode, bool, []string) {
	fake.mkdirAsMutex.RLock()
	defer fake.mkdirAsMutex.RUnlock()
	return fake.mkdirAsArgsForCall[i].rootfsPath, fake.mkdirAsArgsForCall[i].uid, fake.mkdirAsArgsForCall[i].gid, fake.mkdirAsArgsForCall[i].mode, fake.mkdirAsArgsForCall[i].recreate, fake.mkdirAsArgsForCall[i].path
}

func (fake *FakeMkdirer) MkdirAsReturns(result1 error) {
	fake.MkdirAsStub = nil
	fake.mkdirAsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMkdirer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.mkdirAsMutex.RLock()
	defer fake.mkdirAsMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeMkdirer) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ runrunc.Mkdirer = new(FakeMkdirer)