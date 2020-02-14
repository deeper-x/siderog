package memory

import (
	"log"
	"strings"

	"github.com/gomodule/redigo/redis"
)

// Valuer defines methods of value CRUD
type Valuer interface {
	SetValue(redis.Conn, string, string) interface{}
	GetValue(redis.Conn, string, string) interface{}
	Close(redis.Conn)
}

// Token is the session key
type Token struct{}

// NewConn create redis connection
func NewConn() redis.Conn {
	conn, err := redis.Dial("tcp", ":6379")

	if err != nil {
		log.Println(err)
	}

	return conn
}

// SetValue store value in memory
func (t Token) SetValue(conn redis.Conn, name, value string) interface{} {
	val, err := conn.Do("SET", name, value)
	// defer conn.Close()

	if err != nil {
		log.Println(err)
	}

	return val
}

// GetValue retrieves value from memory
func (t Token) GetValue(conn redis.Conn, value string) bool {
	val, err := conn.Do("GET", value)
	// defer conn.Close()
	var retVal bool

	if err != nil {
		log.Println(err)
	}

	if val == nil {
		return false
	}

	inSlice := val.([]uint8)

	letters := []string{}
	for _, v := range inSlice {
		letters = append(letters, string(v))
	}

	retStr := strings.Join(letters, "")

	if retStr != "true" {
		retVal = true
	}

	return retVal
}

// Close redis
func (t Token) Close(conn redis.Conn) {
	conn.Close()
}
