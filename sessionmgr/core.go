package sessionmgr

// SessionIsOn just check is token is a string
func SessionIsOn(token string) bool {
	if len(token) > 0 {
		return true
	}

	return false
}
