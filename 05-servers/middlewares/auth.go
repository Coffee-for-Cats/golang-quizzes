package mw

import (
	db "coffee-server/database"
	"encoding/base64"
	"net/http"
	"net/url"
	"strings"
)

func Access(req *http.Request) bool {
	header := req.Header.Get("Authorization")

	raw_token, basic := strings.CutPrefix(header, "Basic ")
	if !basic {
		return false
	}
	// token is base64 encoded for some reason =)
	decoded_token, err := base64.StdEncoding.DecodeString(raw_token)
	if err != nil {
		return false
	}
	// user must base64-encode it just *after* url-escaping.
	token, err := url.QueryUnescape(string(decoded_token))
	if err != nil {
		return false
	}

	split_token := strings.Split(token, ":")
	name := split_token[0]
	pass := split_token[1]

	return db.VerifyUser(name, pass)
}
