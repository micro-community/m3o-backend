// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	customers "github.com/m3o/services/customers/proto"
	"github.com/micro/micro/v3/service/client"
)

type FakeCustomersService struct {
	BanStub        func(context.Context, *customers.BanRequest, ...client.CallOption) (*customers.BanResponse, error)
	banMutex       sync.RWMutex
	banArgsForCall []struct {
		arg1 context.Context
		arg2 *customers.BanRequest
		arg3 []client.CallOption
	}
	banReturns struct {
		result1 *customers.BanResponse
		result2 error
	}
	banReturnsOnCall map[int]struct {
		result1 *customers.BanResponse
		result2 error
	}
	CreateStub        func(context.Context, *customers.CreateRequest, ...client.CallOption) (*customers.CreateResponse, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 context.Context
		arg2 *customers.CreateRequest
		arg3 []client.CallOption
	}
	createReturns struct {
		result1 *customers.CreateResponse
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 *customers.CreateResponse
		result2 error
	}
	DeleteStub        func(context.Context, *customers.DeleteRequest, ...client.CallOption) (*customers.DeleteResponse, error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 context.Context
		arg2 *customers.DeleteRequest
		arg3 []client.CallOption
	}
	deleteReturns struct {
		result1 *customers.DeleteResponse
		result2 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 *customers.DeleteResponse
		result2 error
	}
	ListStub        func(context.Context, *customers.ListRequest, ...client.CallOption) (*customers.ListResponse, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 context.Context
		arg2 *customers.ListRequest
		arg3 []client.CallOption
	}
	listReturns struct {
		result1 *customers.ListResponse
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 *customers.ListResponse
		result2 error
	}
	LoginStub        func(context.Context, *customers.LoginRequest, ...client.CallOption) (*customers.LoginResponse, error)
	loginMutex       sync.RWMutex
	loginArgsForCall []struct {
		arg1 context.Context
		arg2 *customers.LoginRequest
		arg3 []client.CallOption
	}
	loginReturns struct {
		result1 *customers.LoginResponse
		result2 error
	}
	loginReturnsOnCall map[int]struct {
		result1 *customers.LoginResponse
		result2 error
	}
	LogoutStub        func(context.Context, *customers.LogoutRequest, ...client.CallOption) (*customers.LogoutResponse, error)
	logoutMutex       sync.RWMutex
	logoutArgsForCall []struct {
		arg1 context.Context
		arg2 *customers.LogoutRequest
		arg3 []client.CallOption
	}
	logoutReturns struct {
		result1 *customers.LogoutResponse
		result2 error
	}
	logoutReturnsOnCall map[int]struct {
		result1 *customers.LogoutResponse
		result2 error
	}
	MarkVerifiedStub        func(context.Context, *customers.MarkVerifiedRequest, ...client.CallOption) (*customers.MarkVerifiedResponse, error)
	markVerifiedMutex       sync.RWMutex
	markVerifiedArgsForCall []struct {
		arg1 context.Context
		arg2 *customers.MarkVerifiedRequest
		arg3 []client.CallOption
	}
	markVerifiedReturns struct {
		result1 *customers.MarkVerifiedResponse
		result2 error
	}
	markVerifiedReturnsOnCall map[int]struct {
		result1 *customers.MarkVerifiedResponse
		result2 error
	}
	ReadStub        func(context.Context, *customers.ReadRequest, ...client.CallOption) (*customers.ReadResponse, error)
	readMutex       sync.RWMutex
	readArgsForCall []struct {
		arg1 context.Context
		arg2 *customers.ReadRequest
		arg3 []client.CallOption
	}
	readReturns struct {
		result1 *customers.ReadResponse
		result2 error
	}
	readReturnsOnCall map[int]struct {
		result1 *customers.ReadResponse
		result2 error
	}
	UnbanStub        func(context.Context, *customers.UnbanRequest, ...client.CallOption) (*customers.UnbanResponse, error)
	unbanMutex       sync.RWMutex
	unbanArgsForCall []struct {
		arg1 context.Context
		arg2 *customers.UnbanRequest
		arg3 []client.CallOption
	}
	unbanReturns struct {
		result1 *customers.UnbanResponse
		result2 error
	}
	unbanReturnsOnCall map[int]struct {
		result1 *customers.UnbanResponse
		result2 error
	}
	UpdateStub        func(context.Context, *customers.UpdateRequest, ...client.CallOption) (*customers.UpdateResponse, error)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 context.Context
		arg2 *customers.UpdateRequest
		arg3 []client.CallOption
	}
	updateReturns struct {
		result1 *customers.UpdateResponse
		result2 error
	}
	updateReturnsOnCall map[int]struct {
		result1 *customers.UpdateResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCustomersService) Ban(arg1 context.Context, arg2 *customers.BanRequest, arg3 ...client.CallOption) (*customers.BanResponse, error) {
	fake.banMutex.Lock()
	ret, specificReturn := fake.banReturnsOnCall[len(fake.banArgsForCall)]
	fake.banArgsForCall = append(fake.banArgsForCall, struct {
		arg1 context.Context
		arg2 *customers.BanRequest
		arg3 []client.CallOption
	}{arg1, arg2, arg3})
	stub := fake.BanStub
	fakeReturns := fake.banReturns
	fake.recordInvocation("Ban", []interface{}{arg1, arg2, arg3})
	fake.banMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCustomersService) BanCallCount() int {
	fake.banMutex.RLock()
	defer fake.banMutex.RUnlock()
	return len(fake.banArgsForCall)
}

func (fake *FakeCustomersService) BanCalls(stub func(context.Context, *customers.BanRequest, ...client.CallOption) (*customers.BanResponse, error)) {
	fake.banMutex.Lock()
	defer fake.banMutex.Unlock()
	fake.BanStub = stub
}

func (fake *FakeCustomersService) BanArgsForCall(i int) (context.Context, *customers.BanRequest, []client.CallOption) {
	fake.banMutex.RLock()
	defer fake.banMutex.RUnlock()
	argsForCall := fake.banArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCustomersService) BanReturns(result1 *customers.BanResponse, result2 error) {
	fake.banMutex.Lock()
	defer fake.banMutex.Unlock()
	fake.BanStub = nil
	fake.banReturns = struct {
		result1 *customers.BanResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomersService) BanReturnsOnCall(i int, result1 *customers.BanResponse, result2 error) {
	fake.banMutex.Lock()
	defer fake.banMutex.Unlock()
	fake.BanStub = nil
	if fake.banReturnsOnCall == nil {
		fake.banReturnsOnCall = make(map[int]struct {
			result1 *customers.BanResponse
			result2 error
		})
	}
	fake.banReturnsOnCall[i] = struct {
		result1 *customers.BanResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomersService) Create(arg1 context.Context, arg2 *customers.CreateRequest, arg3 ...client.CallOption) (*customers.CreateResponse, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 context.Context
		arg2 *customers.CreateRequest
		arg3 []client.CallOption
	}{arg1, arg2, arg3})
	stub := fake.CreateStub
	fakeReturns := fake.createReturns
	fake.recordInvocation("Create", []interface{}{arg1, arg2, arg3})
	fake.createMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCustomersService) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeCustomersService) CreateCalls(stub func(context.Context, *customers.CreateRequest, ...client.CallOption) (*customers.CreateResponse, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeCustomersService) CreateArgsForCall(i int) (context.Context, *customers.CreateRequest, []client.CallOption) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCustomersService) CreateReturns(result1 *customers.CreateResponse, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 *customers.CreateResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomersService) CreateReturnsOnCall(i int, result1 *customers.CreateResponse, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 *customers.CreateResponse
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 *customers.CreateResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomersService) Delete(arg1 context.Context, arg2 *customers.DeleteRequest, arg3 ...client.CallOption) (*customers.DeleteResponse, error) {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 context.Context
		arg2 *customers.DeleteRequest
		arg3 []client.CallOption
	}{arg1, arg2, arg3})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1, arg2, arg3})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCustomersService) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeCustomersService) DeleteCalls(stub func(context.Context, *customers.DeleteRequest, ...client.CallOption) (*customers.DeleteResponse, error)) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeCustomersService) DeleteArgsForCall(i int) (context.Context, *customers.DeleteRequest, []client.CallOption) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCustomersService) DeleteReturns(result1 *customers.DeleteResponse, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 *customers.DeleteResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomersService) DeleteReturnsOnCall(i int, result1 *customers.DeleteResponse, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 *customers.DeleteResponse
			result2 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 *customers.DeleteResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomersService) List(arg1 context.Context, arg2 *customers.ListRequest, arg3 ...client.CallOption) (*customers.ListResponse, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 context.Context
		arg2 *customers.ListRequest
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

func (fake *FakeCustomersService) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeCustomersService) ListCalls(stub func(context.Context, *customers.ListRequest, ...client.CallOption) (*customers.ListResponse, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakeCustomersService) ListArgsForCall(i int) (context.Context, *customers.ListRequest, []client.CallOption) {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	argsForCall := fake.listArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCustomersService) ListReturns(result1 *customers.ListResponse, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 *customers.ListResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomersService) ListReturnsOnCall(i int, result1 *customers.ListResponse, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 *customers.ListResponse
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 *customers.ListResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomersService) Login(arg1 context.Context, arg2 *customers.LoginRequest, arg3 ...client.CallOption) (*customers.LoginResponse, error) {
	fake.loginMutex.Lock()
	ret, specificReturn := fake.loginReturnsOnCall[len(fake.loginArgsForCall)]
	fake.loginArgsForCall = append(fake.loginArgsForCall, struct {
		arg1 context.Context
		arg2 *customers.LoginRequest
		arg3 []client.CallOption
	}{arg1, arg2, arg3})
	stub := fake.LoginStub
	fakeReturns := fake.loginReturns
	fake.recordInvocation("Login", []interface{}{arg1, arg2, arg3})
	fake.loginMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCustomersService) LoginCallCount() int {
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	return len(fake.loginArgsForCall)
}

func (fake *FakeCustomersService) LoginCalls(stub func(context.Context, *customers.LoginRequest, ...client.CallOption) (*customers.LoginResponse, error)) {
	fake.loginMutex.Lock()
	defer fake.loginMutex.Unlock()
	fake.LoginStub = stub
}

func (fake *FakeCustomersService) LoginArgsForCall(i int) (context.Context, *customers.LoginRequest, []client.CallOption) {
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	argsForCall := fake.loginArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCustomersService) LoginReturns(result1 *customers.LoginResponse, result2 error) {
	fake.loginMutex.Lock()
	defer fake.loginMutex.Unlock()
	fake.LoginStub = nil
	fake.loginReturns = struct {
		result1 *customers.LoginResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomersService) LoginReturnsOnCall(i int, result1 *customers.LoginResponse, result2 error) {
	fake.loginMutex.Lock()
	defer fake.loginMutex.Unlock()
	fake.LoginStub = nil
	if fake.loginReturnsOnCall == nil {
		fake.loginReturnsOnCall = make(map[int]struct {
			result1 *customers.LoginResponse
			result2 error
		})
	}
	fake.loginReturnsOnCall[i] = struct {
		result1 *customers.LoginResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomersService) Logout(arg1 context.Context, arg2 *customers.LogoutRequest, arg3 ...client.CallOption) (*customers.LogoutResponse, error) {
	fake.logoutMutex.Lock()
	ret, specificReturn := fake.logoutReturnsOnCall[len(fake.logoutArgsForCall)]
	fake.logoutArgsForCall = append(fake.logoutArgsForCall, struct {
		arg1 context.Context
		arg2 *customers.LogoutRequest
		arg3 []client.CallOption
	}{arg1, arg2, arg3})
	stub := fake.LogoutStub
	fakeReturns := fake.logoutReturns
	fake.recordInvocation("Logout", []interface{}{arg1, arg2, arg3})
	fake.logoutMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCustomersService) LogoutCallCount() int {
	fake.logoutMutex.RLock()
	defer fake.logoutMutex.RUnlock()
	return len(fake.logoutArgsForCall)
}

func (fake *FakeCustomersService) LogoutCalls(stub func(context.Context, *customers.LogoutRequest, ...client.CallOption) (*customers.LogoutResponse, error)) {
	fake.logoutMutex.Lock()
	defer fake.logoutMutex.Unlock()
	fake.LogoutStub = stub
}

func (fake *FakeCustomersService) LogoutArgsForCall(i int) (context.Context, *customers.LogoutRequest, []client.CallOption) {
	fake.logoutMutex.RLock()
	defer fake.logoutMutex.RUnlock()
	argsForCall := fake.logoutArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCustomersService) LogoutReturns(result1 *customers.LogoutResponse, result2 error) {
	fake.logoutMutex.Lock()
	defer fake.logoutMutex.Unlock()
	fake.LogoutStub = nil
	fake.logoutReturns = struct {
		result1 *customers.LogoutResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomersService) LogoutReturnsOnCall(i int, result1 *customers.LogoutResponse, result2 error) {
	fake.logoutMutex.Lock()
	defer fake.logoutMutex.Unlock()
	fake.LogoutStub = nil
	if fake.logoutReturnsOnCall == nil {
		fake.logoutReturnsOnCall = make(map[int]struct {
			result1 *customers.LogoutResponse
			result2 error
		})
	}
	fake.logoutReturnsOnCall[i] = struct {
		result1 *customers.LogoutResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomersService) MarkVerified(arg1 context.Context, arg2 *customers.MarkVerifiedRequest, arg3 ...client.CallOption) (*customers.MarkVerifiedResponse, error) {
	fake.markVerifiedMutex.Lock()
	ret, specificReturn := fake.markVerifiedReturnsOnCall[len(fake.markVerifiedArgsForCall)]
	fake.markVerifiedArgsForCall = append(fake.markVerifiedArgsForCall, struct {
		arg1 context.Context
		arg2 *customers.MarkVerifiedRequest
		arg3 []client.CallOption
	}{arg1, arg2, arg3})
	stub := fake.MarkVerifiedStub
	fakeReturns := fake.markVerifiedReturns
	fake.recordInvocation("MarkVerified", []interface{}{arg1, arg2, arg3})
	fake.markVerifiedMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCustomersService) MarkVerifiedCallCount() int {
	fake.markVerifiedMutex.RLock()
	defer fake.markVerifiedMutex.RUnlock()
	return len(fake.markVerifiedArgsForCall)
}

func (fake *FakeCustomersService) MarkVerifiedCalls(stub func(context.Context, *customers.MarkVerifiedRequest, ...client.CallOption) (*customers.MarkVerifiedResponse, error)) {
	fake.markVerifiedMutex.Lock()
	defer fake.markVerifiedMutex.Unlock()
	fake.MarkVerifiedStub = stub
}

func (fake *FakeCustomersService) MarkVerifiedArgsForCall(i int) (context.Context, *customers.MarkVerifiedRequest, []client.CallOption) {
	fake.markVerifiedMutex.RLock()
	defer fake.markVerifiedMutex.RUnlock()
	argsForCall := fake.markVerifiedArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCustomersService) MarkVerifiedReturns(result1 *customers.MarkVerifiedResponse, result2 error) {
	fake.markVerifiedMutex.Lock()
	defer fake.markVerifiedMutex.Unlock()
	fake.MarkVerifiedStub = nil
	fake.markVerifiedReturns = struct {
		result1 *customers.MarkVerifiedResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomersService) MarkVerifiedReturnsOnCall(i int, result1 *customers.MarkVerifiedResponse, result2 error) {
	fake.markVerifiedMutex.Lock()
	defer fake.markVerifiedMutex.Unlock()
	fake.MarkVerifiedStub = nil
	if fake.markVerifiedReturnsOnCall == nil {
		fake.markVerifiedReturnsOnCall = make(map[int]struct {
			result1 *customers.MarkVerifiedResponse
			result2 error
		})
	}
	fake.markVerifiedReturnsOnCall[i] = struct {
		result1 *customers.MarkVerifiedResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomersService) Read(arg1 context.Context, arg2 *customers.ReadRequest, arg3 ...client.CallOption) (*customers.ReadResponse, error) {
	fake.readMutex.Lock()
	ret, specificReturn := fake.readReturnsOnCall[len(fake.readArgsForCall)]
	fake.readArgsForCall = append(fake.readArgsForCall, struct {
		arg1 context.Context
		arg2 *customers.ReadRequest
		arg3 []client.CallOption
	}{arg1, arg2, arg3})
	stub := fake.ReadStub
	fakeReturns := fake.readReturns
	fake.recordInvocation("Read", []interface{}{arg1, arg2, arg3})
	fake.readMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCustomersService) ReadCallCount() int {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return len(fake.readArgsForCall)
}

func (fake *FakeCustomersService) ReadCalls(stub func(context.Context, *customers.ReadRequest, ...client.CallOption) (*customers.ReadResponse, error)) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = stub
}

func (fake *FakeCustomersService) ReadArgsForCall(i int) (context.Context, *customers.ReadRequest, []client.CallOption) {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	argsForCall := fake.readArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCustomersService) ReadReturns(result1 *customers.ReadResponse, result2 error) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = nil
	fake.readReturns = struct {
		result1 *customers.ReadResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomersService) ReadReturnsOnCall(i int, result1 *customers.ReadResponse, result2 error) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = nil
	if fake.readReturnsOnCall == nil {
		fake.readReturnsOnCall = make(map[int]struct {
			result1 *customers.ReadResponse
			result2 error
		})
	}
	fake.readReturnsOnCall[i] = struct {
		result1 *customers.ReadResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomersService) Unban(arg1 context.Context, arg2 *customers.UnbanRequest, arg3 ...client.CallOption) (*customers.UnbanResponse, error) {
	fake.unbanMutex.Lock()
	ret, specificReturn := fake.unbanReturnsOnCall[len(fake.unbanArgsForCall)]
	fake.unbanArgsForCall = append(fake.unbanArgsForCall, struct {
		arg1 context.Context
		arg2 *customers.UnbanRequest
		arg3 []client.CallOption
	}{arg1, arg2, arg3})
	stub := fake.UnbanStub
	fakeReturns := fake.unbanReturns
	fake.recordInvocation("Unban", []interface{}{arg1, arg2, arg3})
	fake.unbanMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCustomersService) UnbanCallCount() int {
	fake.unbanMutex.RLock()
	defer fake.unbanMutex.RUnlock()
	return len(fake.unbanArgsForCall)
}

func (fake *FakeCustomersService) UnbanCalls(stub func(context.Context, *customers.UnbanRequest, ...client.CallOption) (*customers.UnbanResponse, error)) {
	fake.unbanMutex.Lock()
	defer fake.unbanMutex.Unlock()
	fake.UnbanStub = stub
}

func (fake *FakeCustomersService) UnbanArgsForCall(i int) (context.Context, *customers.UnbanRequest, []client.CallOption) {
	fake.unbanMutex.RLock()
	defer fake.unbanMutex.RUnlock()
	argsForCall := fake.unbanArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCustomersService) UnbanReturns(result1 *customers.UnbanResponse, result2 error) {
	fake.unbanMutex.Lock()
	defer fake.unbanMutex.Unlock()
	fake.UnbanStub = nil
	fake.unbanReturns = struct {
		result1 *customers.UnbanResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomersService) UnbanReturnsOnCall(i int, result1 *customers.UnbanResponse, result2 error) {
	fake.unbanMutex.Lock()
	defer fake.unbanMutex.Unlock()
	fake.UnbanStub = nil
	if fake.unbanReturnsOnCall == nil {
		fake.unbanReturnsOnCall = make(map[int]struct {
			result1 *customers.UnbanResponse
			result2 error
		})
	}
	fake.unbanReturnsOnCall[i] = struct {
		result1 *customers.UnbanResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomersService) Update(arg1 context.Context, arg2 *customers.UpdateRequest, arg3 ...client.CallOption) (*customers.UpdateResponse, error) {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 context.Context
		arg2 *customers.UpdateRequest
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

func (fake *FakeCustomersService) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeCustomersService) UpdateCalls(stub func(context.Context, *customers.UpdateRequest, ...client.CallOption) (*customers.UpdateResponse, error)) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakeCustomersService) UpdateArgsForCall(i int) (context.Context, *customers.UpdateRequest, []client.CallOption) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCustomersService) UpdateReturns(result1 *customers.UpdateResponse, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 *customers.UpdateResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomersService) UpdateReturnsOnCall(i int, result1 *customers.UpdateResponse, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 *customers.UpdateResponse
			result2 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 *customers.UpdateResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomersService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.banMutex.RLock()
	defer fake.banMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	fake.logoutMutex.RLock()
	defer fake.logoutMutex.RUnlock()
	fake.markVerifiedMutex.RLock()
	defer fake.markVerifiedMutex.RUnlock()
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	fake.unbanMutex.RLock()
	defer fake.unbanMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCustomersService) recordInvocation(key string, args []interface{}) {
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

var _ customers.CustomersService = new(FakeCustomersService)
