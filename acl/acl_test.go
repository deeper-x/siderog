package acl

import (
	"testing"
)

func TestSavePolicy(t *testing.T) {
	na := MockNewAdapter()
	dummyPath := "/somewhere/over/the/rainbow"
	enf := MockNewEnforcer(dummyPath, na)

	wa := NewWebAdapter(enf)
	ok := wa.StorePolicy("bob", "foo", "baz")

	if !ok {
		t.Errorf("Saving policy failed, got %v", ok)
	}
}

func TestCheckPolicy(t *testing.T) {
	na := MockNewAdapter()
	dummyPath := "/somewhere/over/the/rainbow"
	enf := MockNewEnforcer(dummyPath, na)

	wa := NewWebAdapter(enf)
	ok := wa.CheckPolicy("bob", "/admin", "GET")

	if !ok {
		t.Errorf("Check policy failed, got %v", ok)
	}
}
