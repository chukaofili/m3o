// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	"m3o.dev/platform/service/client"
	publicapi "m3o.dev/api/publicapi/proto"
)

type FakePublicapiService struct {
	GetStub        func(context.Context, *publicapi.GetRequest, ...client.CallOption) (*publicapi.GetResponse, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 context.Context
		arg2 *publicapi.GetRequest
		arg3 []client.CallOption
	}
	getReturns struct {
		result1 *publicapi.GetResponse
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 *publicapi.GetResponse
		result2 error
	}
	ListStub        func(context.Context, *publicapi.ListRequest, ...client.CallOption) (*publicapi.ListResponse, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 context.Context
		arg2 *publicapi.ListRequest
		arg3 []client.CallOption
	}
	listReturns struct {
		result1 *publicapi.ListResponse
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 *publicapi.ListResponse
		result2 error
	}
	PublishStub        func(context.Context, *publicapi.PublishRequest, ...client.CallOption) (*publicapi.PublishResponse, error)
	publishMutex       sync.RWMutex
	publishArgsForCall []struct {
		arg1 context.Context
		arg2 *publicapi.PublishRequest
		arg3 []client.CallOption
	}
	publishReturns struct {
		result1 *publicapi.PublishResponse
		result2 error
	}
	publishReturnsOnCall map[int]struct {
		result1 *publicapi.PublishResponse
		result2 error
	}
	RemoveStub        func(context.Context, *publicapi.RemoveRequest, ...client.CallOption) (*publicapi.RemoveResponse, error)
	removeMutex       sync.RWMutex
	removeArgsForCall []struct {
		arg1 context.Context
		arg2 *publicapi.RemoveRequest
		arg3 []client.CallOption
	}
	removeReturns struct {
		result1 *publicapi.RemoveResponse
		result2 error
	}
	removeReturnsOnCall map[int]struct {
		result1 *publicapi.RemoveResponse
		result2 error
	}
	UpdateStub        func(context.Context, *publicapi.UpdateRequest, ...client.CallOption) (*publicapi.UpdateResponse, error)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 context.Context
		arg2 *publicapi.UpdateRequest
		arg3 []client.CallOption
	}
	updateReturns struct {
		result1 *publicapi.UpdateResponse
		result2 error
	}
	updateReturnsOnCall map[int]struct {
		result1 *publicapi.UpdateResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePublicapiService) Get(arg1 context.Context, arg2 *publicapi.GetRequest, arg3 ...client.CallOption) (*publicapi.GetResponse, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 context.Context
		arg2 *publicapi.GetRequest
		arg3 []client.CallOption
	}{arg1, arg2, arg3})
	stub := fake.GetStub
	fakeReturns := fake.getReturns
	fake.recordInvocation("Get", []interface{}{arg1, arg2, arg3})
	fake.getMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePublicapiService) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakePublicapiService) GetCalls(stub func(context.Context, *publicapi.GetRequest, ...client.CallOption) (*publicapi.GetResponse, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakePublicapiService) GetArgsForCall(i int) (context.Context, *publicapi.GetRequest, []client.CallOption) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePublicapiService) GetReturns(result1 *publicapi.GetResponse, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *publicapi.GetResponse
		result2 error
	}{result1, result2}
}

func (fake *FakePublicapiService) GetReturnsOnCall(i int, result1 *publicapi.GetResponse, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 *publicapi.GetResponse
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 *publicapi.GetResponse
		result2 error
	}{result1, result2}
}

func (fake *FakePublicapiService) List(arg1 context.Context, arg2 *publicapi.ListRequest, arg3 ...client.CallOption) (*publicapi.ListResponse, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 context.Context
		arg2 *publicapi.ListRequest
		arg3 []client.CallOption
	}{arg1, arg2, arg3})
	stub := fake.ListStub
	fakeReturns := fake.listReturns
	fake.recordInvocation("List", []interface{}{arg1, arg2, arg3})
	fake.listMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePublicapiService) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakePublicapiService) ListCalls(stub func(context.Context, *publicapi.ListRequest, ...client.CallOption) (*publicapi.ListResponse, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakePublicapiService) ListArgsForCall(i int) (context.Context, *publicapi.ListRequest, []client.CallOption) {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	argsForCall := fake.listArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePublicapiService) ListReturns(result1 *publicapi.ListResponse, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 *publicapi.ListResponse
		result2 error
	}{result1, result2}
}

func (fake *FakePublicapiService) ListReturnsOnCall(i int, result1 *publicapi.ListResponse, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 *publicapi.ListResponse
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 *publicapi.ListResponse
		result2 error
	}{result1, result2}
}

func (fake *FakePublicapiService) Publish(arg1 context.Context, arg2 *publicapi.PublishRequest, arg3 ...client.CallOption) (*publicapi.PublishResponse, error) {
	fake.publishMutex.Lock()
	ret, specificReturn := fake.publishReturnsOnCall[len(fake.publishArgsForCall)]
	fake.publishArgsForCall = append(fake.publishArgsForCall, struct {
		arg1 context.Context
		arg2 *publicapi.PublishRequest
		arg3 []client.CallOption
	}{arg1, arg2, arg3})
	stub := fake.PublishStub
	fakeReturns := fake.publishReturns
	fake.recordInvocation("Publish", []interface{}{arg1, arg2, arg3})
	fake.publishMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePublicapiService) PublishCallCount() int {
	fake.publishMutex.RLock()
	defer fake.publishMutex.RUnlock()
	return len(fake.publishArgsForCall)
}

func (fake *FakePublicapiService) PublishCalls(stub func(context.Context, *publicapi.PublishRequest, ...client.CallOption) (*publicapi.PublishResponse, error)) {
	fake.publishMutex.Lock()
	defer fake.publishMutex.Unlock()
	fake.PublishStub = stub
}

func (fake *FakePublicapiService) PublishArgsForCall(i int) (context.Context, *publicapi.PublishRequest, []client.CallOption) {
	fake.publishMutex.RLock()
	defer fake.publishMutex.RUnlock()
	argsForCall := fake.publishArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePublicapiService) PublishReturns(result1 *publicapi.PublishResponse, result2 error) {
	fake.publishMutex.Lock()
	defer fake.publishMutex.Unlock()
	fake.PublishStub = nil
	fake.publishReturns = struct {
		result1 *publicapi.PublishResponse
		result2 error
	}{result1, result2}
}

func (fake *FakePublicapiService) PublishReturnsOnCall(i int, result1 *publicapi.PublishResponse, result2 error) {
	fake.publishMutex.Lock()
	defer fake.publishMutex.Unlock()
	fake.PublishStub = nil
	if fake.publishReturnsOnCall == nil {
		fake.publishReturnsOnCall = make(map[int]struct {
			result1 *publicapi.PublishResponse
			result2 error
		})
	}
	fake.publishReturnsOnCall[i] = struct {
		result1 *publicapi.PublishResponse
		result2 error
	}{result1, result2}
}

func (fake *FakePublicapiService) Remove(arg1 context.Context, arg2 *publicapi.RemoveRequest, arg3 ...client.CallOption) (*publicapi.RemoveResponse, error) {
	fake.removeMutex.Lock()
	ret, specificReturn := fake.removeReturnsOnCall[len(fake.removeArgsForCall)]
	fake.removeArgsForCall = append(fake.removeArgsForCall, struct {
		arg1 context.Context
		arg2 *publicapi.RemoveRequest
		arg3 []client.CallOption
	}{arg1, arg2, arg3})
	stub := fake.RemoveStub
	fakeReturns := fake.removeReturns
	fake.recordInvocation("Remove", []interface{}{arg1, arg2, arg3})
	fake.removeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePublicapiService) RemoveCallCount() int {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return len(fake.removeArgsForCall)
}

func (fake *FakePublicapiService) RemoveCalls(stub func(context.Context, *publicapi.RemoveRequest, ...client.CallOption) (*publicapi.RemoveResponse, error)) {
	fake.removeMutex.Lock()
	defer fake.removeMutex.Unlock()
	fake.RemoveStub = stub
}

func (fake *FakePublicapiService) RemoveArgsForCall(i int) (context.Context, *publicapi.RemoveRequest, []client.CallOption) {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	argsForCall := fake.removeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePublicapiService) RemoveReturns(result1 *publicapi.RemoveResponse, result2 error) {
	fake.removeMutex.Lock()
	defer fake.removeMutex.Unlock()
	fake.RemoveStub = nil
	fake.removeReturns = struct {
		result1 *publicapi.RemoveResponse
		result2 error
	}{result1, result2}
}

func (fake *FakePublicapiService) RemoveReturnsOnCall(i int, result1 *publicapi.RemoveResponse, result2 error) {
	fake.removeMutex.Lock()
	defer fake.removeMutex.Unlock()
	fake.RemoveStub = nil
	if fake.removeReturnsOnCall == nil {
		fake.removeReturnsOnCall = make(map[int]struct {
			result1 *publicapi.RemoveResponse
			result2 error
		})
	}
	fake.removeReturnsOnCall[i] = struct {
		result1 *publicapi.RemoveResponse
		result2 error
	}{result1, result2}
}

func (fake *FakePublicapiService) Update(arg1 context.Context, arg2 *publicapi.UpdateRequest, arg3 ...client.CallOption) (*publicapi.UpdateResponse, error) {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 context.Context
		arg2 *publicapi.UpdateRequest
		arg3 []client.CallOption
	}{arg1, arg2, arg3})
	stub := fake.UpdateStub
	fakeReturns := fake.updateReturns
	fake.recordInvocation("Update", []interface{}{arg1, arg2, arg3})
	fake.updateMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePublicapiService) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakePublicapiService) UpdateCalls(stub func(context.Context, *publicapi.UpdateRequest, ...client.CallOption) (*publicapi.UpdateResponse, error)) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakePublicapiService) UpdateArgsForCall(i int) (context.Context, *publicapi.UpdateRequest, []client.CallOption) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePublicapiService) UpdateReturns(result1 *publicapi.UpdateResponse, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 *publicapi.UpdateResponse
		result2 error
	}{result1, result2}
}

func (fake *FakePublicapiService) UpdateReturnsOnCall(i int, result1 *publicapi.UpdateResponse, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 *publicapi.UpdateResponse
			result2 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 *publicapi.UpdateResponse
		result2 error
	}{result1, result2}
}

func (fake *FakePublicapiService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.publishMutex.RLock()
	defer fake.publishMutex.RUnlock()
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePublicapiService) recordInvocation(key string, args []interface{}) {
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

var _ publicapi.PublicapiService = new(FakePublicapiService)
