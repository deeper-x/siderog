package token

import (
	"crypto/hmac"
	"crypto/sha256"
	"log"
	"strings"

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
func (m Machine) HashString(key, input string) string {
	mac := hmac.New(sha256.New, []byte(input))

	mac.Write([]byte(key))

	// return fmt.Sprintf("%s", mac.Sum(nil))
	sliceHash := string(mac.Sum(nil))

	var container []string
	for _, v := range sliceHash {
		container = append(container, string(v))
	}

	return strings.Join(container, "")
}
