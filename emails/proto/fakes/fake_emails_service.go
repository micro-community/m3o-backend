// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	emails "github.com/m3o/services/emails/proto"
	"github.com/micro/micro/v3/service/client"
)

type FakeEmailsService struct {
	SendStub        func(context.Context, *emails.SendRequest, ...client.CallOption) (*emails.SendResponse, error)
	sendMutex       sync.RWMutex
	sendArgsForCall []struct {
		arg1 context.Context
		arg2 *emails.SendRequest
		arg3 []client.CallOption
	}
	sendReturns struct {
		result1 *emails.SendResponse
		result2 error
	}
	sendReturnsOnCall map[int]struct {
		result1 *emails.SendResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEmailsService) Send(arg1 context.Context, arg2 *emails.SendRequest, arg3 ...client.CallOption) (*emails.SendResponse, error) {
	fake.sendMutex.Lock()
	ret, specificReturn := fake.sendReturnsOnCall[len(fake.sendArgsForCall)]
	fake.sendArgsForCall = append(fake.sendArgsForCall, struct {
		arg1 context.Context
		arg2 *emails.SendRequest
		arg3 []client.CallOption
	}{arg1, arg2, arg3})
	stub := fake.SendStub
	fakeReturns := fake.sendReturns
	fake.recordInvocation("Send", []interface{}{arg1, arg2, arg3})
	fake.sendMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEmailsService) SendCallCount() int {
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	return len(fake.sendArgsForCall)
}

func (fake *FakeEmailsService) SendCalls(stub func(context.Context, *emails.SendRequest, ...client.CallOption) (*emails.SendResponse, error)) {
	fake.sendMutex.Lock()
	defer fake.sendMutex.Unlock()
	fake.SendStub = stub
}

func (fake *FakeEmailsService) SendArgsForCall(i int) (context.Context, *emails.SendRequest, []client.CallOption) {
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	argsForCall := fake.sendArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeEmailsService) SendReturns(result1 *emails.SendResponse, result2 error) {
	fake.sendMutex.Lock()
	defer fake.sendMutex.Unlock()
	fake.SendStub = nil
	fake.sendReturns = struct {
		result1 *emails.SendResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeEmailsService) SendReturnsOnCall(i int, result1 *emails.SendResponse, result2 error) {
	fake.sendMutex.Lock()
	defer fake.sendMutex.Unlock()
	fake.SendStub = nil
	if fake.sendReturnsOnCall == nil {
		fake.sendReturnsOnCall = make(map[int]struct {
			result1 *emails.SendResponse
			result2 error
		})
	}
	fake.sendReturnsOnCall[i] = struct {
		result1 *emails.SendResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeEmailsService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEmailsService) recordInvocation(key string, args []interface{}) {
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

var _ emails.EmailsService = new(FakeEmailsService)