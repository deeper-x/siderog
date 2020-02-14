package memory

import (
	"testing"

	"github.com/rafaeljusto/redigomock"
)

func TestSetValue(t *testing.T) {
	mockConn := redigomock.NewConn()
	mt := MockToken{}

	val := mt.SetValue(mockConn, "token", "098204982")

	if val != "OK" {
		t.Fatal("OK error", val)
	}
}

func TestGetValue(t *testing.T) {
	mockConn := redigomock.NewConn()
	mt := MockToken{}

	expected := "justorius"

	val := mt.GetValue(mockConn, "token")

	if val != expected {
		t.Error("retval not as expected", val)
	}
}
