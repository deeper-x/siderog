package server

import (
	"io"
	"net/http"

	"github.com/deeper-x/siderog/memory"
	"github.com/deeper-x/siderog/token"
	"github.com/rafaeljusto/redigomock"
)

type MockSession struct {
	isActive bool
}

func (ms MockSession) Start(mc *redigomock.Conn) http.HandlerFunc {
	sFunc := func(w http.ResponseWriter, r *http.Request) {
		m := token.MockMachine{}

		// TODO - MachineID should not be passed publicly - please hash it
		ID := m.GetID()
		hash := m.HashString(ID)

		// TODO - check if it's created, createToken should return a bool
		ms.createToken(mc, hash)
		io.WriteString(w, hash)
	}

	return sFunc
}

func (ms MockSession) createToken(conn *redigomock.Conn, inputVal string) {
	token := memory.MockToken{}

	token.SetValue(conn, inputVal, "true")
}
