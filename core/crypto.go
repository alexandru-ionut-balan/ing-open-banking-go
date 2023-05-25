package core

import (
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"

	"github.com/alexandru-ionut-balan/ing-open-banking-go/logger"
)

const base16 = 16

func ExtractSerialNumber(certificate *x509.Certificate) string {
	return certificate.SerialNumber.Text(base16)
}

func Base64(payload []byte) string {
	return base64.StdEncoding.EncodeToString(payload)
}

func Sha256(payload []byte) []byte {
	hash := sha256.New()

	if _, err := hash.Write(payload); err != nil {
		logger.Error("Cannot compute SHA-256 hash of payload", err)
	}

	return hash.Sum(nil)
}

func Sha512(payload []byte) []byte {
	hash := sha512.New()

	if _, err := hash.Write(payload); err != nil {
		logger.Error("Cannot compute SHA-512 hash of payload", err)
	}

	return hash.Sum(nil)
}
