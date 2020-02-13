package memory

import (
	"github.com/gomodule/redigo/redis"
)

type mockToken struct {
	value string
}

// SetValue mocks a redis SET
func (m mockToken) SetValue(r redis.Conn, name, value string) interface{} {
	return "OK"
}

// GetValue mocks a redis GET
func (m mockToken) GetValue(r redis.Conn, name string) interface{} {
	return "justorius"
}

func (m mockToken) Close(r redis.Conn) {
	r.Close()
}
