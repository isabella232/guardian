// Code generated by counterfeiter. DO NOT EDIT.
package runruncfakes

import (
	"sync"

	"code.cloudfoundry.org/guardian/rundmc/runrunc"
	"code.cloudfoundry.org/lager"
)

type FakeRuncStater struct {
	StateStub        func(lager.Logger, string) (runrunc.State, error)
	stateMutex       sync.RWMutex
	stateArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	stateReturns struct {
		result1 runrunc.State
		result2 error
	}
	stateReturnsOnCall map[int]struct {
		result1 runrunc.State
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRuncStater) State(arg1 lager.Logger, arg2 string) (runrunc.State, error) {
	fake.stateMutex.Lock()
	ret, specificReturn := fake.stateReturnsOnCall[len(fake.stateArgsForCall)]
	fake.stateArgsForCall = append(fake.stateArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("State", []interface{}{arg1, arg2})
	fake.stateMutex.Unlock()
	if fake.StateStub != nil {
		return fake.StateStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.stateReturns.result1, fake.stateReturns.result2
}

func (fake *FakeRuncStater) StateCallCount() int {
	fake.stateMutex.RLock()
	defer fake.stateMutex.RUnlock()
	return len(fake.stateArgsForCall)
}

func (fake *FakeRuncStater) StateArgsForCall(i int) (lager.Logger, string) {
	fake.stateMutex.RLock()
	defer fake.stateMutex.RUnlock()
	return fake.stateArgsForCall[i].arg1, fake.stateArgsForCall[i].arg2
}

func (fake *FakeRuncStater) StateReturns(result1 runrunc.State, result2 error) {
	fake.StateStub = nil
	fake.stateReturns = struct {
		result1 runrunc.State
		result2 error
	}{result1, result2}
}

func (fake *FakeRuncStater) StateReturnsOnCall(i int, result1 runrunc.State, result2 error) {
	fake.StateStub = nil
	if fake.stateReturnsOnCall == nil {
		fake.stateReturnsOnCall = make(map[int]struct {
			result1 runrunc.State
			result2 error
		})
	}
	fake.stateReturnsOnCall[i] = struct {
		result1 runrunc.State
		result2 error
	}{result1, result2}
}

func (fake *FakeRuncStater) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.stateMutex.RLock()
	defer fake.stateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRuncStater) recordInvocation(key string, args []interface{}) {
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

var _ runrunc.RuncStater = new(FakeRuncStater)