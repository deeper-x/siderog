package sessionmgr

// SessionIsOn just check is token is a string
func SessionIsOn(token string) bool {
	var retVal bool

	if len(token) > 0 {
		retVal = true
	}

	return retVal
}
