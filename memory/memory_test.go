package memory

import (
	"testing"

	"github.com/rafaeljusto/redigomock"
)

func TestSetValue(t *testing.T) {
	mockConn := redigomock.NewConn()

	cmd := mockConn.Command("SET", "token", "098204982").Expect("OK")

	val := SetValue(mockConn, "token", "098204982")

	if mockConn.Stats(cmd) != 1 {
		t.Fatalf("Command %v never called...", cmd)
	}

	if val != "OK" {
		t.Fatal("OK error", val)
	}
}

func TestGetValue(t *testing.T) {
	mockConn := redigomock.NewConn()
	expected := "justorius"

	cmd := mockConn.Command("GET", "token").Expect(expected)

	val := GetValue(mockConn, "token")

	if mockConn.Stats(cmd) != 1 {
		t.Fatal("cmd not executed")
	}

	if val != expected {
		t.Error("retval not as expected", val)
	}
}
