// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/gorouter/registry"
	"code.cloudfoundry.org/gorouter/route"
)

type FakeRegistry struct {
	RegisterStub        func(uri route.Uri, endpoint *route.Endpoint)
	registerMutex       sync.RWMutex
	registerArgsForCall []struct {
		uri      route.Uri
		endpoint *route.Endpoint
	}
	UnregisterStub        func(uri route.Uri, endpoint *route.Endpoint)
	unregisterMutex       sync.RWMutex
	unregisterArgsForCall []struct {
		uri      route.Uri
		endpoint *route.Endpoint
	}
	LookupStub        func(uri route.Uri) *route.EndpointPool
	lookupMutex       sync.RWMutex
	lookupArgsForCall []struct {
		uri route.Uri
	}
	lookupReturns struct {
		result1 *route.EndpointPool
	}
	lookupReturnsOnCall map[int]struct {
		result1 *route.EndpointPool
	}
	LookupWithInstanceStub        func(uri route.Uri, appID, appIndex string) *route.EndpointPool
	lookupWithInstanceMutex       sync.RWMutex
	lookupWithInstanceArgsForCall []struct {
		uri      route.Uri
		appID    string
		appIndex string
	}
	lookupWithInstanceReturns struct {
		result1 *route.EndpointPool
	}
	lookupWithInstanceReturnsOnCall map[int]struct {
		result1 *route.EndpointPool
	}
	StartPruningCycleStub        func()
	startPruningCycleMutex       sync.RWMutex
	startPruningCycleArgsForCall []struct{}
	StopPruningCycleStub         func()
	stopPruningCycleMutex        sync.RWMutex
	stopPruningCycleArgsForCall  []struct{}
	NumUrisStub                  func() int
	numUrisMutex                 sync.RWMutex
	numUrisArgsForCall           []struct{}
	numUrisReturns               struct {
		result1 int
	}
	numUrisReturnsOnCall map[int]struct {
		result1 int
	}
	NumEndpointsStub        func() int
	numEndpointsMutex       sync.RWMutex
	numEndpointsArgsForCall []struct{}
	numEndpointsReturns     struct {
		result1 int
	}
	numEndpointsReturnsOnCall map[int]struct {
		result1 int
	}
	MarshalJSONStub        func() ([]byte, error)
	marshalJSONMutex       sync.RWMutex
	marshalJSONArgsForCall []struct{}
	marshalJSONReturns     struct {
		result1 []byte
		result2 error
	}
	marshalJSONReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRegistry) Register(uri route.Uri, endpoint *route.Endpoint) {
	fake.registerMutex.Lock()
	fake.registerArgsForCall = append(fake.registerArgsForCall, struct {
		uri      route.Uri
		endpoint *route.Endpoint
	}{uri, endpoint})
	fake.recordInvocation("Register", []interface{}{uri, endpoint})
	fake.registerMutex.Unlock()
	if fake.RegisterStub != nil {
		fake.RegisterStub(uri, endpoint)
	}
}

func (fake *FakeRegistry) RegisterCallCount() int {
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	return len(fake.registerArgsForCall)
}

func (fake *FakeRegistry) RegisterArgsForCall(i int) (route.Uri, *route.Endpoint) {
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	return fake.registerArgsForCall[i].uri, fake.registerArgsForCall[i].endpoint
}

func (fake *FakeRegistry) Unregister(uri route.Uri, endpoint *route.Endpoint) {
	fake.unregisterMutex.Lock()
	fake.unregisterArgsForCall = append(fake.unregisterArgsForCall, struct {
		uri      route.Uri
		endpoint *route.Endpoint
	}{uri, endpoint})
	fake.recordInvocation("Unregister", []interface{}{uri, endpoint})
	fake.unregisterMutex.Unlock()
	if fake.UnregisterStub != nil {
		fake.UnregisterStub(uri, endpoint)
	}
}

func (fake *FakeRegistry) UnregisterCallCount() int {
	fake.unregisterMutex.RLock()
	defer fake.unregisterMutex.RUnlock()
	return len(fake.unregisterArgsForCall)
}

func (fake *FakeRegistry) UnregisterArgsForCall(i int) (route.Uri, *route.Endpoint) {
	fake.unregisterMutex.RLock()
	defer fake.unregisterMutex.RUnlock()
	return fake.unregisterArgsForCall[i].uri, fake.unregisterArgsForCall[i].endpoint
}

func (fake *FakeRegistry) Lookup(uri route.Uri) *route.EndpointPool {
	fake.lookupMutex.Lock()
	ret, specificReturn := fake.lookupReturnsOnCall[len(fake.lookupArgsForCall)]
	fake.lookupArgsForCall = append(fake.lookupArgsForCall, struct {
		uri route.Uri
	}{uri})
	fake.recordInvocation("Lookup", []interface{}{uri})
	fake.lookupMutex.Unlock()
	if fake.LookupStub != nil {
		return fake.LookupStub(uri)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.lookupReturns.result1
}

func (fake *FakeRegistry) LookupCallCount() int {
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	return len(fake.lookupArgsForCall)
}

func (fake *FakeRegistry) LookupArgsForCall(i int) route.Uri {
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	return fake.lookupArgsForCall[i].uri
}

func (fake *FakeRegistry) LookupReturns(result1 *route.EndpointPool) {
	fake.LookupStub = nil
	fake.lookupReturns = struct {
		result1 *route.EndpointPool
	}{result1}
}

func (fake *FakeRegistry) LookupReturnsOnCall(i int, result1 *route.EndpointPool) {
	fake.LookupStub = nil
	if fake.lookupReturnsOnCall == nil {
		fake.lookupReturnsOnCall = make(map[int]struct {
			result1 *route.EndpointPool
		})
	}
	fake.lookupReturnsOnCall[i] = struct {
		result1 *route.EndpointPool
	}{result1}
}

func (fake *FakeRegistry) LookupWithInstance(uri route.Uri, appID string, appIndex string) *route.EndpointPool {
	fake.lookupWithInstanceMutex.Lock()
	ret, specificReturn := fake.lookupWithInstanceReturnsOnCall[len(fake.lookupWithInstanceArgsForCall)]
	fake.lookupWithInstanceArgsForCall = append(fake.lookupWithInstanceArgsForCall, struct {
		uri      route.Uri
		appID    string
		appIndex string
	}{uri, appID, appIndex})
	fake.recordInvocation("LookupWithInstance", []interface{}{uri, appID, appIndex})
	fake.lookupWithInstanceMutex.Unlock()
	if fake.LookupWithInstanceStub != nil {
		return fake.LookupWithInstanceStub(uri, appID, appIndex)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.lookupWithInstanceReturns.result1
}

func (fake *FakeRegistry) LookupWithInstanceCallCount() int {
	fake.lookupWithInstanceMutex.RLock()
	defer fake.lookupWithInstanceMutex.RUnlock()
	return len(fake.lookupWithInstanceArgsForCall)
}

func (fake *FakeRegistry) LookupWithInstanceArgsForCall(i int) (route.Uri, string, string) {
	fake.lookupWithInstanceMutex.RLock()
	defer fake.lookupWithInstanceMutex.RUnlock()
	return fake.lookupWithInstanceArgsForCall[i].uri, fake.lookupWithInstanceArgsForCall[i].appID, fake.lookupWithInstanceArgsForCall[i].appIndex
}

func (fake *FakeRegistry) LookupWithInstanceReturns(result1 *route.EndpointPool) {
	fake.LookupWithInstanceStub = nil
	fake.lookupWithInstanceReturns = struct {
		result1 *route.EndpointPool
	}{result1}
}

func (fake *FakeRegistry) LookupWithInstanceReturnsOnCall(i int, result1 *route.EndpointPool) {
	fake.LookupWithInstanceStub = nil
	if fake.lookupWithInstanceReturnsOnCall == nil {
		fake.lookupWithInstanceReturnsOnCall = make(map[int]struct {
			result1 *route.EndpointPool
		})
	}
	fake.lookupWithInstanceReturnsOnCall[i] = struct {
		result1 *route.EndpointPool
	}{result1}
}

func (fake *FakeRegistry) StartPruningCycle() {
	fake.startPruningCycleMutex.Lock()
	fake.startPruningCycleArgsForCall = append(fake.startPruningCycleArgsForCall, struct{}{})
	fake.recordInvocation("StartPruningCycle", []interface{}{})
	fake.startPruningCycleMutex.Unlock()
	if fake.StartPruningCycleStub != nil {
		fake.StartPruningCycleStub()
	}
}

func (fake *FakeRegistry) StartPruningCycleCallCount() int {
	fake.startPruningCycleMutex.RLock()
	defer fake.startPruningCycleMutex.RUnlock()
	return len(fake.startPruningCycleArgsForCall)
}

func (fake *FakeRegistry) StopPruningCycle() {
	fake.stopPruningCycleMutex.Lock()
	fake.stopPruningCycleArgsForCall = append(fake.stopPruningCycleArgsForCall, struct{}{})
	fake.recordInvocation("StopPruningCycle", []interface{}{})
	fake.stopPruningCycleMutex.Unlock()
	if fake.StopPruningCycleStub != nil {
		fake.StopPruningCycleStub()
	}
}

func (fake *FakeRegistry) StopPruningCycleCallCount() int {
	fake.stopPruningCycleMutex.RLock()
	defer fake.stopPruningCycleMutex.RUnlock()
	return len(fake.stopPruningCycleArgsForCall)
}

func (fake *FakeRegistry) NumUris() int {
	fake.numUrisMutex.Lock()
	ret, specificReturn := fake.numUrisReturnsOnCall[len(fake.numUrisArgsForCall)]
	fake.numUrisArgsForCall = append(fake.numUrisArgsForCall, struct{}{})
	fake.recordInvocation("NumUris", []interface{}{})
	fake.numUrisMutex.Unlock()
	if fake.NumUrisStub != nil {
		return fake.NumUrisStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.numUrisReturns.result1
}

func (fake *FakeRegistry) NumUrisCallCount() int {
	fake.numUrisMutex.RLock()
	defer fake.numUrisMutex.RUnlock()
	return len(fake.numUrisArgsForCall)
}

func (fake *FakeRegistry) NumUrisReturns(result1 int) {
	fake.NumUrisStub = nil
	fake.numUrisReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeRegistry) NumUrisReturnsOnCall(i int, result1 int) {
	fake.NumUrisStub = nil
	if fake.numUrisReturnsOnCall == nil {
		fake.numUrisReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.numUrisReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeRegistry) NumEndpoints() int {
	fake.numEndpointsMutex.Lock()
	ret, specificReturn := fake.numEndpointsReturnsOnCall[len(fake.numEndpointsArgsForCall)]
	fake.numEndpointsArgsForCall = append(fake.numEndpointsArgsForCall, struct{}{})
	fake.recordInvocation("NumEndpoints", []interface{}{})
	fake.numEndpointsMutex.Unlock()
	if fake.NumEndpointsStub != nil {
		return fake.NumEndpointsStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.numEndpointsReturns.result1
}

func (fake *FakeRegistry) NumEndpointsCallCount() int {
	fake.numEndpointsMutex.RLock()
	defer fake.numEndpointsMutex.RUnlock()
	return len(fake.numEndpointsArgsForCall)
}

func (fake *FakeRegistry) NumEndpointsReturns(result1 int) {
	fake.NumEndpointsStub = nil
	fake.numEndpointsReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeRegistry) NumEndpointsReturnsOnCall(i int, result1 int) {
	fake.NumEndpointsStub = nil
	if fake.numEndpointsReturnsOnCall == nil {
		fake.numEndpointsReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.numEndpointsReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeRegistry) MarshalJSON() ([]byte, error) {
	fake.marshalJSONMutex.Lock()
	ret, specificReturn := fake.marshalJSONReturnsOnCall[len(fake.marshalJSONArgsForCall)]
	fake.marshalJSONArgsForCall = append(fake.marshalJSONArgsForCall, struct{}{})
	fake.recordInvocation("MarshalJSON", []interface{}{})
	fake.marshalJSONMutex.Unlock()
	if fake.MarshalJSONStub != nil {
		return fake.MarshalJSONStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.marshalJSONReturns.result1, fake.marshalJSONReturns.result2
}

func (fake *FakeRegistry) MarshalJSONCallCount() int {
	fake.marshalJSONMutex.RLock()
	defer fake.marshalJSONMutex.RUnlock()
	return len(fake.marshalJSONArgsForCall)
}

func (fake *FakeRegistry) MarshalJSONReturns(result1 []byte, result2 error) {
	fake.MarshalJSONStub = nil
	fake.marshalJSONReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeRegistry) MarshalJSONReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.MarshalJSONStub = nil
	if fake.marshalJSONReturnsOnCall == nil {
		fake.marshalJSONReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.marshalJSONReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeRegistry) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	fake.unregisterMutex.RLock()
	defer fake.unregisterMutex.RUnlock()
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	fake.lookupWithInstanceMutex.RLock()
	defer fake.lookupWithInstanceMutex.RUnlock()
	fake.startPruningCycleMutex.RLock()
	defer fake.startPruningCycleMutex.RUnlock()
	fake.stopPruningCycleMutex.RLock()
	defer fake.stopPruningCycleMutex.RUnlock()
	fake.numUrisMutex.RLock()
	defer fake.numUrisMutex.RUnlock()
	fake.numEndpointsMutex.RLock()
	defer fake.numEndpointsMutex.RUnlock()
	fake.marshalJSONMutex.RLock()
	defer fake.marshalJSONMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRegistry) recordInvocation(key string, args []interface{}) {
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

var _ registry.Registry = new(FakeRegistry)
