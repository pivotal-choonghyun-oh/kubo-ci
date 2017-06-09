// Code generated by counterfeiter. DO NOT EDIT.
package vspherefakes

import (
	"sync"
	"vsphere-cleaner/vsphere"
)

type FakeVM struct {
	PowerOffStub        func() error
	powerOffMutex       sync.RWMutex
	powerOffArgsForCall []struct{}
	powerOffReturns     struct {
		result1 error
	}
	powerOffReturnsOnCall map[int]struct {
		result1 error
	}
	DestroyStub        func() error
	destroyMutex       sync.RWMutex
	destroyArgsForCall []struct{}
	destroyReturns     struct {
		result1 error
	}
	destroyReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVM) PowerOff() error {
	fake.powerOffMutex.Lock()
	ret, specificReturn := fake.powerOffReturnsOnCall[len(fake.powerOffArgsForCall)]
	fake.powerOffArgsForCall = append(fake.powerOffArgsForCall, struct{}{})
	fake.recordInvocation("PowerOff", []interface{}{})
	fake.powerOffMutex.Unlock()
	if fake.PowerOffStub != nil {
		return fake.PowerOffStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.powerOffReturns.result1
}

func (fake *FakeVM) PowerOffCallCount() int {
	fake.powerOffMutex.RLock()
	defer fake.powerOffMutex.RUnlock()
	return len(fake.powerOffArgsForCall)
}

func (fake *FakeVM) PowerOffReturns(result1 error) {
	fake.PowerOffStub = nil
	fake.powerOffReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVM) PowerOffReturnsOnCall(i int, result1 error) {
	fake.PowerOffStub = nil
	if fake.powerOffReturnsOnCall == nil {
		fake.powerOffReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.powerOffReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVM) Destroy() error {
	fake.destroyMutex.Lock()
	ret, specificReturn := fake.destroyReturnsOnCall[len(fake.destroyArgsForCall)]
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct{}{})
	fake.recordInvocation("Destroy", []interface{}{})
	fake.destroyMutex.Unlock()
	if fake.DestroyStub != nil {
		return fake.DestroyStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.destroyReturns.result1
}

func (fake *FakeVM) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeVM) DestroyReturns(result1 error) {
	fake.DestroyStub = nil
	fake.destroyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVM) DestroyReturnsOnCall(i int, result1 error) {
	fake.DestroyStub = nil
	if fake.destroyReturnsOnCall == nil {
		fake.destroyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVM) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.powerOffMutex.RLock()
	defer fake.powerOffMutex.RUnlock()
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVM) recordInvocation(key string, args []interface{}) {
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

var _ vsphere.VM = new(FakeVM)