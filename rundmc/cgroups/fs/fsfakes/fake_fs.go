// Code generated by counterfeiter. DO NOT EDIT.
package fsfakes

import (
	"sync"

	"code.cloudfoundry.org/guardian/rundmc/cgroups/fs"
)

type FakeFS struct {
	ChownStub        func(string, int, int) error
	chownMutex       sync.RWMutex
	chownArgsForCall []struct {
		arg1 string
		arg2 int
		arg3 int
	}
	chownReturns struct {
		result1 error
	}
	chownReturnsOnCall map[int]struct {
		result1 error
	}
	MountStub        func(string, string, string, uintptr, string) error
	mountMutex       sync.RWMutex
	mountArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 uintptr
		arg5 string
	}
	mountReturns struct {
		result1 error
	}
	mountReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFS) Chown(arg1 string, arg2 int, arg3 int) error {
	fake.chownMutex.Lock()
	ret, specificReturn := fake.chownReturnsOnCall[len(fake.chownArgsForCall)]
	fake.chownArgsForCall = append(fake.chownArgsForCall, struct {
		arg1 string
		arg2 int
		arg3 int
	}{arg1, arg2, arg3})
	fake.recordInvocation("Chown", []interface{}{arg1, arg2, arg3})
	fake.chownMutex.Unlock()
	if fake.ChownStub != nil {
		return fake.ChownStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.chownReturns.result1
}

func (fake *FakeFS) ChownCallCount() int {
	fake.chownMutex.RLock()
	defer fake.chownMutex.RUnlock()
	return len(fake.chownArgsForCall)
}

func (fake *FakeFS) ChownArgsForCall(i int) (string, int, int) {
	fake.chownMutex.RLock()
	defer fake.chownMutex.RUnlock()
	return fake.chownArgsForCall[i].arg1, fake.chownArgsForCall[i].arg2, fake.chownArgsForCall[i].arg3
}

func (fake *FakeFS) ChownReturns(result1 error) {
	fake.ChownStub = nil
	fake.chownReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFS) ChownReturnsOnCall(i int, result1 error) {
	fake.ChownStub = nil
	if fake.chownReturnsOnCall == nil {
		fake.chownReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.chownReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFS) Mount(arg1 string, arg2 string, arg3 string, arg4 uintptr, arg5 string) error {
	fake.mountMutex.Lock()
	ret, specificReturn := fake.mountReturnsOnCall[len(fake.mountArgsForCall)]
	fake.mountArgsForCall = append(fake.mountArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 uintptr
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("Mount", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.mountMutex.Unlock()
	if fake.MountStub != nil {
		return fake.MountStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.mountReturns.result1
}

func (fake *FakeFS) MountCallCount() int {
	fake.mountMutex.RLock()
	defer fake.mountMutex.RUnlock()
	return len(fake.mountArgsForCall)
}

func (fake *FakeFS) MountArgsForCall(i int) (string, string, string, uintptr, string) {
	fake.mountMutex.RLock()
	defer fake.mountMutex.RUnlock()
	return fake.mountArgsForCall[i].arg1, fake.mountArgsForCall[i].arg2, fake.mountArgsForCall[i].arg3, fake.mountArgsForCall[i].arg4, fake.mountArgsForCall[i].arg5
}

func (fake *FakeFS) MountReturns(result1 error) {
	fake.MountStub = nil
	fake.mountReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFS) MountReturnsOnCall(i int, result1 error) {
	fake.MountStub = nil
	if fake.mountReturnsOnCall == nil {
		fake.mountReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.mountReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFS) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.chownMutex.RLock()
	defer fake.chownMutex.RUnlock()
	fake.mountMutex.RLock()
	defer fake.mountMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFS) recordInvocation(key string, args []interface{}) {
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

var _ fs.FS = new(FakeFS)
