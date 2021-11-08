// Code generated by counterfeiter. DO NOT EDIT.
package volumefakes

import (
	"io"
	"sync"

	"github.com/concourse/concourse/worker/baggageclaim/volume"
)

type FakeStreamer struct {
	InStub        func(io.Reader, string, bool) (bool, error)
	inMutex       sync.RWMutex
	inArgsForCall []struct {
		arg1 io.Reader
		arg2 string
		arg3 bool
	}
	inReturns struct {
		result1 bool
		result2 error
	}
	inReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	OutStub        func(io.Writer, string, bool) error
	outMutex       sync.RWMutex
	outArgsForCall []struct {
		arg1 io.Writer
		arg2 string
		arg3 bool
	}
	outReturns struct {
		result1 error
	}
	outReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStreamer) In(arg1 io.Reader, arg2 string, arg3 bool) (bool, error) {
	fake.inMutex.Lock()
	ret, specificReturn := fake.inReturnsOnCall[len(fake.inArgsForCall)]
	fake.inArgsForCall = append(fake.inArgsForCall, struct {
		arg1 io.Reader
		arg2 string
		arg3 bool
	}{arg1, arg2, arg3})
	stub := fake.InStub
	fakeReturns := fake.inReturns
	fake.recordInvocation("In", []interface{}{arg1, arg2, arg3})
	fake.inMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStreamer) InCallCount() int {
	fake.inMutex.RLock()
	defer fake.inMutex.RUnlock()
	return len(fake.inArgsForCall)
}

func (fake *FakeStreamer) InCalls(stub func(io.Reader, string, bool) (bool, error)) {
	fake.inMutex.Lock()
	defer fake.inMutex.Unlock()
	fake.InStub = stub
}

func (fake *FakeStreamer) InArgsForCall(i int) (io.Reader, string, bool) {
	fake.inMutex.RLock()
	defer fake.inMutex.RUnlock()
	argsForCall := fake.inArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeStreamer) InReturns(result1 bool, result2 error) {
	fake.inMutex.Lock()
	defer fake.inMutex.Unlock()
	fake.InStub = nil
	fake.inReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeStreamer) InReturnsOnCall(i int, result1 bool, result2 error) {
	fake.inMutex.Lock()
	defer fake.inMutex.Unlock()
	fake.InStub = nil
	if fake.inReturnsOnCall == nil {
		fake.inReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.inReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeStreamer) Out(arg1 io.Writer, arg2 string, arg3 bool) error {
	fake.outMutex.Lock()
	ret, specificReturn := fake.outReturnsOnCall[len(fake.outArgsForCall)]
	fake.outArgsForCall = append(fake.outArgsForCall, struct {
		arg1 io.Writer
		arg2 string
		arg3 bool
	}{arg1, arg2, arg3})
	stub := fake.OutStub
	fakeReturns := fake.outReturns
	fake.recordInvocation("Out", []interface{}{arg1, arg2, arg3})
	fake.outMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStreamer) OutCallCount() int {
	fake.outMutex.RLock()
	defer fake.outMutex.RUnlock()
	return len(fake.outArgsForCall)
}

func (fake *FakeStreamer) OutCalls(stub func(io.Writer, string, bool) error) {
	fake.outMutex.Lock()
	defer fake.outMutex.Unlock()
	fake.OutStub = stub
}

func (fake *FakeStreamer) OutArgsForCall(i int) (io.Writer, string, bool) {
	fake.outMutex.RLock()
	defer fake.outMutex.RUnlock()
	argsForCall := fake.outArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeStreamer) OutReturns(result1 error) {
	fake.outMutex.Lock()
	defer fake.outMutex.Unlock()
	fake.OutStub = nil
	fake.outReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStreamer) OutReturnsOnCall(i int, result1 error) {
	fake.outMutex.Lock()
	defer fake.outMutex.Unlock()
	fake.OutStub = nil
	if fake.outReturnsOnCall == nil {
		fake.outReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.outReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStreamer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.inMutex.RLock()
	defer fake.inMutex.RUnlock()
	fake.outMutex.RLock()
	defer fake.outMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStreamer) recordInvocation(key string, args []interface{}) {
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

var _ volume.Streamer = new(FakeStreamer)