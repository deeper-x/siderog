package server

import (
	"io"
	"log"
	"net/http"

	"github.com/deeper-x/siderog/memory"
	"github.com/deeper-x/siderog/token"
	"github.com/gomodule/redigo/redis"
)

// Runner is the main server interface
type Runner interface {
	Start() http.HandlerFunc
}

// Session management object
type Session struct {
	isActive bool
}

// Start creates the memory token
func (s Session) Start(c redis.Conn) http.HandlerFunc {
	sFunc := func(w http.ResponseWriter, r *http.Request) {
		m := token.Machine{}

		// TODO - MachineID should not be passed publicly - please hash it
		ID := m.GetID()

		// TODO - check if it's created, createToken should return a bool
		s.createToken(c, ID)
		io.WriteString(w, ID)
	}

	return sFunc
}

// Check if session is active
func (s Session) Check(c redis.Conn) http.HandlerFunc {
	sFunc := func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()
		token := values["token"][0]

		retVal := "false"

		ok := s.checkToken(c, token)

		if ok {
			retVal = "true"
		}

		io.WriteString(w, retVal)
	}

	return sFunc
}

func (s Session) checkToken(conn redis.Conn, inputVal string) bool {
	token := memory.Token{}

	val := token.GetValue(conn, inputVal)

	return val
}

func (s Session) createToken(conn redis.Conn, inputVal string) {
	token := memory.Token{}

	token.SetValue(conn, inputVal, "true")
}

// RunServer is the http listener/server
func RunServer() {
	conn := memory.NewConn()
	sess := Session{isActive: false}

	http.HandleFunc("/start_session", sess.Start(conn))
	http.HandleFunc("/check_session", sess.Check(conn))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
