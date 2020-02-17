package server

import (
	"log"
	"net/http"

	"github.com/deeper-x/siderog/memory"
)

// RunServer is the http listener/server
func RunServer() {
	conn := memory.NewConn()
	sess := Session{isActive: false}

	http.HandleFunc("/start_session", sess.Start(conn))
	http.HandleFunc("/check_session", sess.Check(conn))
	http.HandleFunc("/admin/new_role", sess.NewRole(conn))

	log.Fatal(http.ListenAndServeTLS(
		":8080",
		"./tls/cert/server.crt",
		"./tls/key/server.key",
		nil),
	)
}
