package main

import (
	"log"

	"github.com/deeper-x/siderog/memory"
)

func main() {
	conn := memory.NewConn()
	memory.SetValue(conn, "val", "demo")

	conn = memory.NewConn()
	res := memory.GetValue(conn, "val")

	log.Println(res)
}
