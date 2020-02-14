package memory

import (
	"github.com/gomodule/redigo/redis"
	"github.com/rafaeljusto/redigomock"
)

type MockToken struct {
	value string
}

// SetValue mocks a redis SET
func (m MockToken) SetValue(r *redigomock.Conn, name, value string) interface{} {
	return "OK"
}

// GetValue mocks a redis GET
func (m MockToken) GetValue(r *redigomock.Conn, name string) interface{} {
	return "justorius"
}

func (m MockToken) Close(r redis.Conn) {
	r.Close()
}
