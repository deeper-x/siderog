package main

import (
	"log"

	"github.com/deeper-x/siderog/memory"
)

func main() {
	conn := memory.NewConn()
	token := memory.Token{}

	token.SetValue(conn, "val", "demo")

	conn = memory.NewConn()
	res := token.GetValue(conn, "val")

	log.Println(res)
}
