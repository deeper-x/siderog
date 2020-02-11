package memory

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

// Token is the session key
type Token struct {
	value string `redis:"value"`
}

// NewConn create redis connection
func NewConn() redis.Conn {
	conn, err := redis.Dial("tcp", ":6379")

	if err != nil {
		log.Println(err)
	}

	return conn
}

// SetValue store value in memory
func SetValue(conn redis.Conn, name, value string) interface{} {
	val, err := conn.Do("SET", name, value)
	defer conn.Close()

	if err != nil {
		log.Println(err)
	}

	return val
}

// GetValue retrieves value from memory
func GetValue(conn redis.Conn, name string) interface{} {
	val, err := conn.Do("GET", name)
	defer conn.Close()

	if err != nil {
		log.Println(err)
	}

	return val
}
