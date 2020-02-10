package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

type User struct {
	name     string `redis:"name"`
	password string `redis:"password"`
}

func main() {
	conn, err := redis.Dial("tcp", ":6379")

	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	SetName(conn, "justorius")
	res := GetName(conn, "name")

	log.Println(res)
}

func SetName(conn redis.Conn, name string) interface{} {
	val, err := conn.Do("SET", "name", name)

	if err != nil {
		log.Println(err)
	}

	return val
}

func GetName(conn redis.Conn, name string) interface{} {
	val, err := conn.Do("GET", "name")
	if err != nil {
		log.Println(err)
	}

	return val
}
