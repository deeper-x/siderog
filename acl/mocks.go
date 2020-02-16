package acl

import (
	"log"

	"github.com/casbin/casbin"
	redisadapter "github.com/casbin/redis-adapter"
)

// MockAuth is the Auth mock
type MockAuth struct {
	Enforcer *casbin.Enforcer
}

// MockNewAuth return redis-adapter.Adapter pointer
func MockNewAuth(p string, redisURI, conf string) Authorizer {
	adp := redisadapter.NewAdapter(p, redisURI)

	return MockAuth{Enforcer: casbin.NewEnforcer(conf, adp)}
}

// SaveRule write policy to db
func (a MockAuth) SaveRule(who, what, method string) bool {
	enf := a.Enforcer

	ok := enf.AddPolicy(who, what, method)
	err := enf.SavePolicy()

	if err != nil {
		log.Println(err)
	}

	return ok
}

// IsAllowed returns is actor is allowed to access resource with given method
func (a MockAuth) IsAllowed(who, what, method string) bool {
	enf := a.Enforcer
	ok := enf.Enforce(who, what, method)

	return ok
}

// FlushPolicy delete policy from memory
func (a MockAuth) FlushPolicy(who, what, method string) bool {
	enf := a.Enforcer
	ok := enf.RemovePolicy(who, what, method)
	err := enf.SavePolicy()

	if err != nil {
		log.Println(err)
	}

	return ok
}
