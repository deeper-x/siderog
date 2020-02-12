package sessionmgr

import "testing"

func TestSessionIsOn(t *testing.T) {
	var ok bool
	ok = SessionIsOn("dummyToken")

	if !ok {
		t.Errorf("%v should be true", ok)
	}

	ok = SessionIsOn("")

	if ok {
		t.Errorf("%v should be false", ok)
	}
}
