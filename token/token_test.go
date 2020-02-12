package token

import (
	"testing"
)

const machineIDLen = 32
const hashLength = 32

type mockMachine struct {
	ID string
}

func (m mockMachine) GetID() string {
	return "12345678901234567890123456789012"
}

func (m mockMachine) HashString(key, input string) string {
	return "12345678901234567890123456789012"
}

func TestGetID(t *testing.T) {
	mt := mockMachine{ID: "1"}
	res := mt.GetID()

	if len(res) != machineIDLen {
		t.Errorf("Token %v not generated correctly", res)
	}
}

func TestHashString(t *testing.T) {
	aString := "2039840283420834"
	key := "a random key..."

	mt := mockMachine{ID: "1"}

	res := mt.HashString(key, aString)

	if len(res) != hashLength {
		t.Errorf("Result %d not generated correctly", len(res))
	}
}
