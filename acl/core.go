package acl

import (
	"log"

	"github.com/casbin/casbin"
	redisadapter "github.com/casbin/redis-adapter"
)

// Authorizer is the acl interface
type Authorizer interface {
	SaveRule(string, string, string) bool
	IsAllowed(string, string, string) bool
	FlushPolicy(string, string, string) bool
}

// Auth is the ACL receiver
type Auth struct {
	Enforcer *casbin.Enforcer
}

// NewAuth return redis-adapter.Adapter pointer
func NewAuth(p string, redisURI, conf string) Authorizer {
	adp := redisadapter.NewAdapter(p, redisURI)

	return Auth{Enforcer: casbin.NewEnforcer(conf, adp)}
}

// SaveRule write policy to db
func (a Auth) SaveRule(who, what, method string) bool {
	enf := a.Enforcer

	ok := enf.AddPolicy(who, what, method)
	err := enf.SavePolicy()

	if err != nil {
		log.Println(err)
	}

	return ok
}

// IsAllowed returns is actor is allowed to access resource with given method
func (a Auth) IsAllowed(who, what, method string) bool {
	enf := a.Enforcer

	ok := enf.Enforce(who, what, method)

	return ok
}

// FlushPolicy delete policy from memory
func (a Auth) FlushPolicy(who, what, method string) bool {
	enf := a.Enforcer
	ok := enf.RemovePolicy(who, what, method)
	err := enf.SavePolicy()

	if err != nil {
		log.Println(err)
	}

	return ok
}

// // NewEnforcer returns Enforcer pointer
// func NewEnforcer(p string, redisURI, conf string) {
// 	adp := redisadapter.NewAdapter(p, redisURI)

// 	enf := casbin.NewEnforcer(conf, adp)

// 	// ## WRITE POLICY ON DB
// 	enf.AddPolicy("bob", "/token", "GET")
// 	enf.SavePolicy()

// 	// # CONSUME POLICY DATA
// 	ok := enf.Enforce("bob", "/token/foo/bar", "GET")

// 	if ok {
// 		log.Println("-->ok")
// 	} else {
// 		log.Println("-->ko")
// 	}

// 	// # DEL POLICY FROM DB
// 	enf.RemovePolicy("bob", "/token", "GET")
// 	enf.SavePolicy()
// }
