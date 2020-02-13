package token

//MockMachine is the twin object
type MockMachine struct {
	ID string
}

// GetID is the mock twin
func (m MockMachine) GetID() string {
	return "1812a91668bc47da8a5d734bd87fe8c5"
}

// HashString is the mock twin
func (m MockMachine) HashString(key, input string) string {
	return "12345678901234567890123456789012"
}
