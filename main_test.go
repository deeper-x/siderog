package main

import (
	"testing"

	"github.com/rafaeljusto/redigomock"
)

func TestSetName(t *testing.T) {
	mockConn := redigomock.NewConn()

	cmd := mockConn.Command("SET", "name", "justorius").Expect("OK")

	val := SetName(mockConn, "justorius")

	if mockConn.Stats(cmd) != 1 {
		t.Fatalf("Command %v never called...", cmd)
	}

	if val != "OK" {
		t.Fatal("OK error", val)
	}
}

func TestGetName(t *testing.T) {
	mockConn := redigomock.NewConn()
	expected := "justorius"

	cmd := mockConn.Command("GET", "name").Expect(expected)

	val := GetName(mockConn, "name")

	if mockConn.Stats(cmd) != 1 {
		t.Fatal("cmd not executed")
	}

	if val != expected {
		t.Error("retval not as expected", val)
	}
}
