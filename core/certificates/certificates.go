package certificates

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"strings"

	"github.com/alexandru-ionut-balan/ing-open-banking-go/core"
	"github.com/alexandru-ionut-balan/ing-open-banking-go/logger"
)

func ReadPemCertificate(path string) (*x509.Certificate, error) {
	bytes, err := core.ReadFileAsBytes(path)
	if err != nil {
		return nil, err
	}

	derBytes, undecodable := pem.Decode(bytes)
	if len(undecodable) > 0 {
		return nil, fmt.Errorf("%s", "The certificate bytes could not be decoded from PEM decoding")
	}

	certificate, err := x509.ParseCertificate(derBytes.Bytes)
	if err != nil {
		return nil, err
	}

	return certificate, nil
}

func ReadPemCertificateAsString(path string) string {
	certificate, err := core.ReadFileAsString(path)
	if err != nil {
		logger.Error(err)
	}

	return strings.ReplaceAll(certificate, "\n", "")
}

// ReadPrivateKey will try to read a PKCS1 private key from {path}. If the reading
// is not successful, then it will try to read it as a PKCS8 key and if this fails as well
// it will return nil.
func ReadPrivateKey(path string) *rsa.PrivateKey {
	bytes, err := core.ReadFileAsBytes(path)
	if err != nil {
		logger.Error(err)
		return nil
	}

	derBytes, undecodable := pem.Decode(bytes)
	if len(undecodable) > 0 {
		logger.Error("Could not PEM decode the private key at " + path)
		return nil
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(derBytes.Bytes)
	if err != nil {
		logger.Error(err)
		logger.Info("Trying to parse as PKCS8 Key...")

		anyKey, err := x509.ParsePKCS8PrivateKey(derBytes.Bytes)
		if err != nil {
			logger.Error(err)
			return nil
		}

		rsaKey, _ := anyKey.(*rsa.PrivateKey)
		return rsaKey
	}

	return privateKey
}
