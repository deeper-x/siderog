package token

import (
	"crypto/sha256"
	"fmt"
	"log"

	"github.com/denisbrodbeck/machineid"
)

// Tokenizer is the token generator
type Tokenizer interface {
	GetID() string
	HashString(key, input string) string
}

// Machine the running PC
type Machine struct {
	ID string
}

// GetID retrive the machine unique identifier
func (m Machine) GetID() string {
	id, err := machineid.ID()

	if err != nil {
		log.Fatal(err)
	}

	return id
}

// HashString a generic input string
func (m Machine) HashString(input string) string {
	hashed := sha256.Sum256([]byte(input))

	return fmt.Sprintf("%x", hashed)
}
