package acl

import (
	"errors"

	"github.com/casbin/casbin/model"
)

// MockAdapter is the mocking receiver
type MockAdapter struct{}

// MockEnforcer is mocking receiver
type MockEnforcer struct{}

// MockNewAdapter returns mock adapter
func MockNewAdapter() *MockAdapter {
	return &MockAdapter{}
}

// MockNewEnforcer returns mock enforcer
func MockNewEnforcer(path string, ad *MockAdapter) *MockEnforcer {
	return &MockEnforcer{}
}

// AddPolicy Mocks new policy
func (me *MockEnforcer) AddPolicy(...interface{}) bool {
	return true
}

// SavePolicy mocks policy saving
func (me *MockEnforcer) SavePolicy() error {
	return nil
}

// LoadPolicy get policy rule to the storage.
func (a *MockAdapter) LoadPolicy(m model.Model) error {
	return errors.New("not implemented")
}

// SavePolicy save policy to storage
func (a *MockAdapter) SavePolicy(m model.Model) error {
	return nil
}

// AddPolicy adds a policy rule to the storage.
func (a *MockAdapter) AddPolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemovePolicy removes a policy rule from the storage.
func (a *MockAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (a *MockAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return nil
}
