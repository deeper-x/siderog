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
func (m MockMachine) HashString(input string) string {
	return "9b1c40d6add1b6a57e30f9aaddb21d1e39d7093a9ed3336170f146c558134077"
}
