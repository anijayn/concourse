// Code generated by counterfeiter. DO NOT EDIT.
package workerfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc/worker"
)

type FakeArtifactSource struct {
	ExistsOnStub        func(lager.Logger, worker.Worker) (worker.Volume, bool, error)
	existsOnMutex       sync.RWMutex
	existsOnArgsForCall []struct {
		arg1 lager.Logger
		arg2 worker.Worker
	}
	existsOnReturns struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	existsOnReturnsOnCall map[int]struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	HandleStub        func() string
	handleMutex       sync.RWMutex
	handleArgsForCall []struct {
	}
	handleReturns struct {
		result1 string
	}
	handleReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeArtifactSource) ExistsOn(arg1 lager.Logger, arg2 worker.Worker) (worker.Volume, bool, error) {
	fake.existsOnMutex.Lock()
	ret, specificReturn := fake.existsOnReturnsOnCall[len(fake.existsOnArgsForCall)]
	fake.existsOnArgsForCall = append(fake.existsOnArgsForCall, struct {
		arg1 lager.Logger
		arg2 worker.Worker
	}{arg1, arg2})
	stub := fake.ExistsOnStub
	fakeReturns := fake.existsOnReturns
	fake.recordInvocation("ExistsOn", []interface{}{arg1, arg2})
	fake.existsOnMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeArtifactSource) ExistsOnCallCount() int {
	fake.existsOnMutex.RLock()
	defer fake.existsOnMutex.RUnlock()
	return len(fake.existsOnArgsForCall)
}

func (fake *FakeArtifactSource) ExistsOnCalls(stub func(lager.Logger, worker.Worker) (worker.Volume, bool, error)) {
	fake.existsOnMutex.Lock()
	defer fake.existsOnMutex.Unlock()
	fake.ExistsOnStub = stub
}

func (fake *FakeArtifactSource) ExistsOnArgsForCall(i int) (lager.Logger, worker.Worker) {
	fake.existsOnMutex.RLock()
	defer fake.existsOnMutex.RUnlock()
	argsForCall := fake.existsOnArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeArtifactSource) ExistsOnReturns(result1 worker.Volume, result2 bool, result3 error) {
	fake.existsOnMutex.Lock()
	defer fake.existsOnMutex.Unlock()
	fake.ExistsOnStub = nil
	fake.existsOnReturns = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeArtifactSource) ExistsOnReturnsOnCall(i int, result1 worker.Volume, result2 bool, result3 error) {
	fake.existsOnMutex.Lock()
	defer fake.existsOnMutex.Unlock()
	fake.ExistsOnStub = nil
	if fake.existsOnReturnsOnCall == nil {
		fake.existsOnReturnsOnCall = make(map[int]struct {
			result1 worker.Volume
			result2 bool
			result3 error
		})
	}
	fake.existsOnReturnsOnCall[i] = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeArtifactSource) Handle() string {
	fake.handleMutex.Lock()
	ret, specificReturn := fake.handleReturnsOnCall[len(fake.handleArgsForCall)]
	fake.handleArgsForCall = append(fake.handleArgsForCall, struct {
	}{})
	stub := fake.HandleStub
	fakeReturns := fake.handleReturns
	fake.recordInvocation("Handle", []interface{}{})
	fake.handleMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeArtifactSource) HandleCallCount() int {
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	return len(fake.handleArgsForCall)
}

func (fake *FakeArtifactSource) HandleCalls(stub func() string) {
	fake.handleMutex.Lock()
	defer fake.handleMutex.Unlock()
	fake.HandleStub = stub
}

func (fake *FakeArtifactSource) HandleReturns(result1 string) {
	fake.handleMutex.Lock()
	defer fake.handleMutex.Unlock()
	fake.HandleStub = nil
	fake.handleReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeArtifactSource) HandleReturnsOnCall(i int, result1 string) {
	fake.handleMutex.Lock()
	defer fake.handleMutex.Unlock()
	fake.HandleStub = nil
	if fake.handleReturnsOnCall == nil {
		fake.handleReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.handleReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeArtifactSource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.existsOnMutex.RLock()
	defer fake.existsOnMutex.RUnlock()
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeArtifactSource) recordInvocation(key string, args []interface{}) {
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

var _ worker.ArtifactSource = new(FakeArtifactSource)
