package token

import (
	"crypto/hmac"
	"crypto/sha256"
	"log"

	"github.com/denisbrodbeck/machineid"
)

// GetMachineID retrive the machine unique identifier
func GetMachineID() string {
	id, err := machineid.ID()

	if err != nil {
		log.Fatal(err)
	}

	return id
}

// HashString a generic input string
func HashString(key, input string) string {
	mac := hmac.New(sha256.New, []byte(input))

	mac.Write([]byte(key))

	// return fmt.Sprintf("%s", mac.Sum(nil))
	return string(mac.Sum(nil))
}
