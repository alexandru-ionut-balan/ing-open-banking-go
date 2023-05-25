package core

import (
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"strings"

	"github.com/alexandru-ionut-balan/ing-open-banking-go/logger"
)

type HashingAlgorithm string

const (
	AlgoSha256 HashingAlgorithm = "SHA-256"
	AlgoSha512 HashingAlgorithm = "SHA-512"
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

// Digest takes as input a string payload and a hashing algorithm,
// returning the following output: {HashingAlgorithm}=Base64({HashingAlgorithm}(payload)).
//
// The output value is the value of the HTTP "Digest" header. This value is used to verify the integrity of the message
// body that you're sending to ING in any request. It is mandatory in all requests.
func Digest(payload string, algorithm HashingAlgorithm) string {
	formattedPayload := strings.TrimRight(payload, "\n")

	switch algorithm {
	case AlgoSha512:
		return string(AlgoSha512) + "=" + Base64(Sha512([]byte(formattedPayload)))
	default:
		return string(AlgoSha256) + "=" + Base64(Sha256([]byte(formattedPayload)))
	}
}
