package acl

import (
	"log"

	"github.com/casbin/casbin/model"
)

// AdapterInterface todo doc
type AdapterInterface interface {
	AddPolicy(string, string, []string) error
	SavePolicy(model model.Model) error

	RemovePolicy(string, string, []string) error
	LoadPolicy(model model.Model) error
}

// EnforcerInterface todo doc
type EnforcerInterface interface {
	AddPolicy(...interface{}) bool
	SavePolicy() error
	Enforce(...interface{}) bool
}

// WebAdapter todo doc
type WebAdapter struct {
	Enforcer EnforcerInterface
}

// NewWebAdapter return WebAdapter obj
func NewWebAdapter(enf EnforcerInterface) WebAdapter {
	return WebAdapter{Enforcer: enf}
}

// StorePolicy save data to redis db
func (wa WebAdapter) StorePolicy(who, what, how string) bool {
	var retVal bool

	wa.Enforcer.AddPolicy(who, what, how)
	err := wa.Enforcer.SavePolicy()

	if err != nil {
		log.Println(err)
		return retVal
	}

	retVal = true

	return retVal
}

// CheckPolicy check policy exit status
func (wa WebAdapter) CheckPolicy(who, what, how string) bool {
	ok := wa.Enforcer.Enforce(who, what, how)

	return ok
}
