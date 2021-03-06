// Code generated by counterfeiter. DO NOT EDIT.
package recordsfakes

import (
	"bosh-dns/dns/server/records"
	"sync"
)

type FakeFileReader struct {
	GetStub        func() ([]byte, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct{}
	getReturns     struct {
		result1 []byte
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	SubscribeStub        func() <-chan bool
	subscribeMutex       sync.RWMutex
	subscribeArgsForCall []struct{}
	subscribeReturns     struct {
		result1 <-chan bool
	}
	subscribeReturnsOnCall map[int]struct {
		result1 <-chan bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFileReader) Get() ([]byte, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct{}{})
	fake.recordInvocation("Get", []interface{}{})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getReturns.result1, fake.getReturns.result2
}

func (fake *FakeFileReader) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeFileReader) GetReturns(result1 []byte, result2 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeFileReader) GetReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeFileReader) Subscribe() <-chan bool {
	fake.subscribeMutex.Lock()
	ret, specificReturn := fake.subscribeReturnsOnCall[len(fake.subscribeArgsForCall)]
	fake.subscribeArgsForCall = append(fake.subscribeArgsForCall, struct{}{})
	fake.recordInvocation("Subscribe", []interface{}{})
	fake.subscribeMutex.Unlock()
	if fake.SubscribeStub != nil {
		return fake.SubscribeStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.subscribeReturns.result1
}

func (fake *FakeFileReader) SubscribeCallCount() int {
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	return len(fake.subscribeArgsForCall)
}

func (fake *FakeFileReader) SubscribeReturns(result1 <-chan bool) {
	fake.SubscribeStub = nil
	fake.subscribeReturns = struct {
		result1 <-chan bool
	}{result1}
}

func (fake *FakeFileReader) SubscribeReturnsOnCall(i int, result1 <-chan bool) {
	fake.SubscribeStub = nil
	if fake.subscribeReturnsOnCall == nil {
		fake.subscribeReturnsOnCall = make(map[int]struct {
			result1 <-chan bool
		})
	}
	fake.subscribeReturnsOnCall[i] = struct {
		result1 <-chan bool
	}{result1}
}

func (fake *FakeFileReader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeFileReader) recordInvocation(key string, args []interface{}) {
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

var _ records.FileReader = new(FakeFileReader)
