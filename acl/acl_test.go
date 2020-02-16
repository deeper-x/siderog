package acl

import (
	"testing"
)

var redisURI = "127.0.0.1:6379"
var confPath = "./assets/rbac_model.conf"

func TestNewAuth(t *testing.T) {
	authr := MockNewAuth("tcp", redisURI, confPath)

	if authr == nil {
		t.Error("New Auth error")
	}
}

func TestSaveFlushPolicy(t *testing.T) {
	ma := MockNewAuth("tcp", redisURI, confPath)

	ok := ma.SaveRule("foo", "/test", "GET")

	if !ok {
		t.Error("SavePolicy error")
	}

	ok = ma.FlushPolicy("foo", "/test", "GET")

	if !ok {
		t.Errorf("FlushPolicy error")
	}
}

func TestSaveIsAllowedFlushPolicy(t *testing.T) {
	ma := MockNewAuth("tcp", redisURI, confPath)

	ok := ma.SaveRule("foo", "/test", "GET")

	if !ok {
		t.Error("SavePolicy error")
	}

	ok = ma.IsAllowed("foo", "/test", "GET")

	if !ok {
		t.Error("IsAllowed error")
	}

	ok = ma.IsAllowed("foo", "/admin", "GET")

	if ok {
		t.Error("IsAllowed error: admin area should be denied")
	}

	ok = ma.FlushPolicy("foo", "/test", "GET")

	if !ok {
		t.Errorf("FlushPolicy error")
	}
}
