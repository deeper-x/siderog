package token

import (
	"testing"
)

const machineIDLen = 32
const hashLength = 32

func TestGetID(t *testing.T) {
	mt := MockMachine{ID: "1"}
	res := mt.GetID()

	if len(res) != machineIDLen {
		t.Errorf("Token %v not generated correctly", res)
	}
}

func TestHashString(t *testing.T) {
	aString := "2039840283420834"
	key := "a random key..."

	mt := MockMachine{ID: "1"}

	res := mt.HashString(key, aString)

	if len(res) != hashLength {
		t.Errorf("Result %d not generated correctly", len(res))
	}
}
