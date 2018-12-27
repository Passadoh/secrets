package secrets

import (
	"io/ioutil"
	"log"
	"strings"
)

// Handle takes a Docker Secret path (e.g. /run/secrets/<secret>) as an argument and
// returns the secret as a string.
func Handle(scrt string) (secret string) {
	secretBytes, err := ioutil.ReadFile(scrt) // You must create a secret ahead of time
	if err != nil {
		log.Fatal("Secret init error", err)
	}
	secret = strings.TrimSpace(string(secretBytes))
	return
}
