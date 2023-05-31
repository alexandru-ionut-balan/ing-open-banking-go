package encode

import (
	"crypto"
	"encoding/base64"

	"github.com/alexandru-ionut-balan/ing-open-banking-go/core/log"
)

func Base64(payload []byte) string {
	return base64.StdEncoding.EncodeToString(payload)
}

func Hash(payload string, algorithm crypto.Hash) []byte {
	hash := algorithm.New()

	if _, err := hash.Write([]byte(payload)); err != nil {
		log.Error.Println("Could not compute hash of "+payload, err)
	}

	return hash.Sum(nil)
}
