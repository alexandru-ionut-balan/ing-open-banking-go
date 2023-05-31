package encode

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
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

func Sign(payload string, algorithm crypto.Hash, privateKey *rsa.PrivateKey) string {
	hashedPayload := Hash(payload, algorithm)
	signedPayload, err := rsa.SignPKCS1v15(rand.Reader, privateKey, algorithm, hashedPayload)
	if err != nil {
		log.Error.Println("Cannot sign payload "+payload, err)
		return ""
	}

	return Base64(signedPayload)
}
