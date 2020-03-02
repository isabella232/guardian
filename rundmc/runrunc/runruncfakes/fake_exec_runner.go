// Code generated by counterfeiter. DO NOT EDIT.
package runruncfakes

import (
	"io"
	"sync"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/guardian/rundmc/runrunc"
	"code.cloudfoundry.org/lager"
)

type FakeExecRunner struct {
	AttachStub        func(lager.Logger, string, string, garden.ProcessIO) (garden.Process, error)
	attachMutex       sync.RWMutex
	attachArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
		arg4 garden.ProcessIO
	}
	attachReturns struct {
		result1 garden.Process
		result2 error
	}
	attachReturnsOnCall map[int]struct {
		result1 garden.Process
		result2 error
	}
	RunStub        func(lager.Logger, string, string, garden.ProcessIO, bool, io.Reader, func() error) (garden.Process, error)
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
		arg4 garden.ProcessIO
		arg5 bool
		arg6 io.Reader
		arg7 func() error
	}
	runReturns struct {
		result1 garden.Process
		result2 error
	}
	runReturnsOnCall map[int]struct {
		result1 garden.Process
		result2 error
	}
	RunPeaStub        func(lager.Logger, string, string, garden.ProcessIO, bool, io.Reader, func() error) (garden.Process, error)
	runPeaMutex       sync.RWMutex
	runPeaArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
		arg4 garden.ProcessIO
		arg5 bool
		arg6 io.Reader
		arg7 func() error
	}
	runPeaReturns struct {
		result1 garden.Process
		result2 error
	}
	runPeaReturnsOnCall map[int]struct {
		result1 garden.Process
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeExecRunner) Attach(arg1 lager.Logger, arg2 string, arg3 string, arg4 garden.ProcessIO) (garden.Process, error) {
	fake.attachMutex.Lock()
	ret, specificReturn := fake.attachReturnsOnCall[len(fake.attachArgsForCall)]
	fake.attachArgsForCall = append(fake.attachArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
		arg4 garden.ProcessIO
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Attach", []interface{}{arg1, arg2, arg3, arg4})
	fake.attachMutex.Unlock()
	if fake.AttachStub != nil {
		return fake.AttachStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.attachReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeExecRunner) AttachCallCount() int {
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	return len(fake.attachArgsForCall)
}

func (fake *FakeExecRunner) AttachCalls(stub func(lager.Logger, string, string, garden.ProcessIO) (garden.Process, error)) {
	fake.attachMutex.Lock()
	defer fake.attachMutex.Unlock()
	fake.AttachStub = stub
}

func (fake *FakeExecRunner) AttachArgsForCall(i int) (lager.Logger, string, string, garden.ProcessIO) {
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	argsForCall := fake.attachArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeExecRunner) AttachReturns(result1 garden.Process, result2 error) {
	fake.attachMutex.Lock()
	defer fake.attachMutex.Unlock()
	fake.AttachStub = nil
	fake.attachReturns = struct {
		result1 garden.Process
		result2 error
	}{result1, result2}
}

func (fake *FakeExecRunner) AttachReturnsOnCall(i int, result1 garden.Process, result2 error) {
	fake.attachMutex.Lock()
	defer fake.attachMutex.Unlock()
	fake.AttachStub = nil
	if fake.attachReturnsOnCall == nil {
		fake.attachReturnsOnCall = make(map[int]struct {
			result1 garden.Process
			result2 error
		})
	}
	fake.attachReturnsOnCall[i] = struct {
		result1 garden.Process
		result2 error
	}{result1, result2}
}

func (fake *FakeExecRunner) Run(arg1 lager.Logger, arg2 string, arg3 string, arg4 garden.ProcessIO, arg5 bool, arg6 io.Reader, arg7 func() error) (garden.Process, error) {
	fake.runMutex.Lock()
	ret, specificReturn := fake.runReturnsOnCall[len(fake.runArgsForCall)]
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
		arg4 garden.ProcessIO
		arg5 bool
		arg6 io.Reader
		arg7 func() error
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.recordInvocation("Run", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.runReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeExecRunner) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeExecRunner) RunCalls(stub func(lager.Logger, string, string, garden.ProcessIO, bool, io.Reader, func() error) (garden.Process, error)) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = stub
}

func (fake *FakeExecRunner) RunArgsForCall(i int) (lager.Logger, string, string, garden.ProcessIO, bool, io.Reader, func() error) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	argsForCall := fake.runArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7
}

func (fake *FakeExecRunner) RunReturns(result1 garden.Process, result2 error) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 garden.Process
		result2 error
	}{result1, result2}
}

func (fake *FakeExecRunner) RunReturnsOnCall(i int, result1 garden.Process, result2 error) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = nil
	if fake.runReturnsOnCall == nil {
		fake.runReturnsOnCall = make(map[int]struct {
			result1 garden.Process
			result2 error
		})
	}
	fake.runReturnsOnCall[i] = struct {
		result1 garden.Process
		result2 error
	}{result1, result2}
}

func (fake *FakeExecRunner) RunPea(arg1 lager.Logger, arg2 string, arg3 string, arg4 garden.ProcessIO, arg5 bool, arg6 io.Reader, arg7 func() error) (garden.Process, error) {
	fake.runPeaMutex.Lock()
	ret, specificReturn := fake.runPeaReturnsOnCall[len(fake.runPeaArgsForCall)]
	fake.runPeaArgsForCall = append(fake.runPeaArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
		arg4 garden.ProcessIO
		arg5 bool
		arg6 io.Reader
		arg7 func() error
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.recordInvocation("RunPea", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.runPeaMutex.Unlock()
	if fake.RunPeaStub != nil {
		return fake.RunPeaStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.runPeaReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeExecRunner) RunPeaCallCount() int {
	fake.runPeaMutex.RLock()
	defer fake.runPeaMutex.RUnlock()
	return len(fake.runPeaArgsForCall)
}

func (fake *FakeExecRunner) RunPeaCalls(stub func(lager.Logger, string, string, garden.ProcessIO, bool, io.Reader, func() error) (garden.Process, error)) {
	fake.runPeaMutex.Lock()
	defer fake.runPeaMutex.Unlock()
	fake.RunPeaStub = stub
}

func (fake *FakeExecRunner) RunPeaArgsForCall(i int) (lager.Logger, string, string, garden.ProcessIO, bool, io.Reader, func() error) {
	fake.runPeaMutex.RLock()
	defer fake.runPeaMutex.RUnlock()
	argsForCall := fake.runPeaArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7
}

func (fake *FakeExecRunner) RunPeaReturns(result1 garden.Process, result2 error) {
	fake.runPeaMutex.Lock()
	defer fake.runPeaMutex.Unlock()
	fake.RunPeaStub = nil
	fake.runPeaReturns = struct {
		result1 garden.Process
		result2 error
	}{result1, result2}
}

func (fake *FakeExecRunner) RunPeaReturnsOnCall(i int, result1 garden.Process, result2 error) {
	fake.runPeaMutex.Lock()
	defer fake.runPeaMutex.Unlock()
	fake.RunPeaStub = nil
	if fake.runPeaReturnsOnCall == nil {
		fake.runPeaReturnsOnCall = make(map[int]struct {
			result1 garden.Process
			result2 error
		})
	}
	fake.runPeaReturnsOnCall[i] = struct {
		result1 garden.Process
		result2 error
	}{result1, result2}
}

func (fake *FakeExecRunner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	fake.runPeaMutex.RLock()
	defer fake.runPeaMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeExecRunner) recordInvocation(key string, args []interface{}) {
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

var _ runrunc.ExecRunner = new(FakeExecRunner)
