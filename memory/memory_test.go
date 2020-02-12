package memory

import (
	"testing"

	"github.com/gomodule/redigo/redis"
	"github.com/rafaeljusto/redigomock"
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

func TestSetValue(t *testing.T) {
	mockConn := redigomock.NewConn()
	mt := mockToken{}

	val := mt.SetValue(mockConn, "token", "098204982")

	if val != "OK" {
		t.Fatal("OK error", val)
	}
}

func TestGetValue(t *testing.T) {
	mockConn := redigomock.NewConn()
	mt := mockToken{}

	expected := "justorius"

	val := mt.GetValue(mockConn, "token")

	if val != expected {
		t.Error("retval not as expected", val)
	}
}
