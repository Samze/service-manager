// Code generated by counterfeiter. DO NOT EDIT.
package securityfakes

import (
	"sync"

	"github.com/Peripli/service-manager/security"
)

type FakeKeyFetcher struct {
	GetEncryptionKeyStub        func() ([]byte, error)
	getEncryptionKeyMutex       sync.RWMutex
	getEncryptionKeyArgsForCall []struct{}
	getEncryptionKeyReturns     struct {
		result1 []byte
		result2 error
	}
	getEncryptionKeyReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeKeyFetcher) GetEncryptionKey() ([]byte, error) {
	fake.getEncryptionKeyMutex.Lock()
	ret, specificReturn := fake.getEncryptionKeyReturnsOnCall[len(fake.getEncryptionKeyArgsForCall)]
	fake.getEncryptionKeyArgsForCall = append(fake.getEncryptionKeyArgsForCall, struct{}{})
	fake.recordInvocation("GetEncryptionKey", []interface{}{})
	fake.getEncryptionKeyMutex.Unlock()
	if fake.GetEncryptionKeyStub != nil {
		return fake.GetEncryptionKeyStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getEncryptionKeyReturns.result1, fake.getEncryptionKeyReturns.result2
}

func (fake *FakeKeyFetcher) GetEncryptionKeyCallCount() int {
	fake.getEncryptionKeyMutex.RLock()
	defer fake.getEncryptionKeyMutex.RUnlock()
	return len(fake.getEncryptionKeyArgsForCall)
}

func (fake *FakeKeyFetcher) GetEncryptionKeyReturns(result1 []byte, result2 error) {
	fake.GetEncryptionKeyStub = nil
	fake.getEncryptionKeyReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeKeyFetcher) GetEncryptionKeyReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.GetEncryptionKeyStub = nil
	if fake.getEncryptionKeyReturnsOnCall == nil {
		fake.getEncryptionKeyReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getEncryptionKeyReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeKeyFetcher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getEncryptionKeyMutex.RLock()
	defer fake.getEncryptionKeyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeKeyFetcher) recordInvocation(key string, args []interface{}) {
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

var _ security.KeyFetcher = new(FakeKeyFetcher)