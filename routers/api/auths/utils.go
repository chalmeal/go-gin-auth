/**
* auth support utils
 */
package auths

import (
	"crypto/sha256"
	"fmt"
)

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
