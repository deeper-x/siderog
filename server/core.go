package server

import (
	"io"
	"log"
	"net/http"

	"github.com/deeper-x/siderog/memory"
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
		values := r.URL.Query()
		token := values["token"][0]

		s.createToken(c, token)
		io.WriteString(w, token)
	}

	return sFunc
}

func (s Session) createToken(conn redis.Conn, inputVal string) {
	token := memory.Token{}

	token.SetValue(conn, "token", inputVal)
}

// RunServer is the http listener/server
func RunServer() {
	conn := memory.NewConn()
	sess := Session{isActive: false}

	http.HandleFunc("/start_session", sess.Start(conn))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
