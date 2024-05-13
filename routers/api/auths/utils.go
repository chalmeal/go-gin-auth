/**
* auth support utils
 */
package auths

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func createState() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

/*
* hasshing password sha-256
 */
func hashPW(password string) string {
	pw := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", pw)
}

/*
* hasshing accessToken sha-256
 */
func hashToken(token string) string {
	t := sha256.Sum256([]byte(token))
	return fmt.Sprintf("%x", t)
}
