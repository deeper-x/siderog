package server

import (
	"io"
	"net/http"

	"github.com/deeper-x/siderog/memory"
	"github.com/deeper-x/siderog/token"
	"github.com/rafaeljusto/redigomock"
)

// MockSession define mock receiver
type MockSession struct {
	isActive bool
}

// Start mocks machine ID and hashing http response
func (ms MockSession) Start(mc *redigomock.Conn) http.HandlerFunc {
	sFunc := func(w http.ResponseWriter, r *http.Request) {
		m := token.MockMachine{}

		// TODO - MachineID should not be passed publicly - please hash it
		ID := m.GetID()
		hash := m.HashString(ID)

		// TODO - check if it's created, createToken should return a bool
		ms.CreateToken(mc, hash)
		io.WriteString(w, hash)
	}

	return sFunc
}

// CreateToken mocks token generation
func (ms MockSession) CreateToken(conn *redigomock.Conn, inputVal string) {
	token := memory.MockToken{}

	token.SetValue(conn, inputVal, "true")
}

// Check if session is active
func (ms MockSession) Check(c *redigomock.Conn) http.HandlerFunc {
	sFunc := func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()
		token := values["token"][0]

		retVal := "false"

		ok := ms.CheckToken(c, token)

		if ok {
			retVal = "true"
		}

		io.WriteString(w, retVal)
	}

	return sFunc
}

// CheckToken mocks the CheckToken method on Session receiver
func (ms MockSession) CheckToken(conn *redigomock.Conn, inputVal string) bool {
	token := memory.MockToken{}

	val := token.GetValue(conn, inputVal)

	return val
}
