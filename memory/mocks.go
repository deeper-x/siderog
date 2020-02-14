package memory

import (
	"github.com/rafaeljusto/redigomock"
)

// MockToken is the Tocken moct struct
type MockToken struct{}

// SetValue mocks a redis SET
func (m MockToken) SetValue(r *redigomock.Conn, name, value string) interface{} {
	return "OK"
}

// GetValue mocks a redis GET
func (m MockToken) GetValue(r *redigomock.Conn, name string) bool {
	return true
}

// Close is the redis connection wrapper
func (m MockToken) Close(r *redigomock.Conn) {
	r.Close()
}
