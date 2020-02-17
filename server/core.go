package server

import (
	"io"
	"log"
	"net/http"

	"github.com/casbin/casbin"
	redisadapter "github.com/casbin/redis-adapter"
	"github.com/deeper-x/siderog/acl"
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
		hash := m.HashString(ID)

		// TODO - check if it's created, createToken should return a bool
		s.createToken(c, hash)
		_, err := io.WriteString(w, hash)

		if err != nil {
			log.Println(err)
		}

		// test acl instance here
		adapter := redisadapter.NewAdapter("tcp", "127.0.0.1:6379")
		enf := casbin.NewEnforcer("./acl/conf/rbac_model.conf", adapter)

		wa := acl.NewWebAdapter(enf)
		wa.StorePolicy("marta", "data3", "read")
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

		_, err := io.WriteString(w, retVal)

		if err != nil {
			log.Println(err)
		}
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
