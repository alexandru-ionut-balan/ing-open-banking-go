package clients

import (
	"crypto"
	"crypto/rsa"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/alexandru-ionut-balan/ing-open-banking-go/core"
	"github.com/alexandru-ionut-balan/ing-open-banking-go/core/log"
	"github.com/alexandru-ionut-balan/ing-open-banking-go/encode"
)

func Digest(header *http.Header, grant core.Grant, algorithm crypto.Hash) {
	header.Add("Digest", encode.Base64(encode.Hash("grant-type="+string(grant), algorithm)))
}

func Date(header *http.Header) {
	location, err := time.LoadLocation("GMT")
	if err != nil {
		log.Error.Println("Cannot load location: GMT", err)
	}

	date := time.Now().In(location).Format(time.RFC1123)

	header.Add("Date", date)
}

func Signature(header http.Header, keyId string, algorithm crypto.Hash, privateKey *rsa.PrivateKey) {
	var headerNames string
	var signaturePayload string
	var signatureAlgorithm string

	for headerName, headerValue := range header {
		var headerValueString string
		for _, item := range headerValue {
			headerValueString += item + ","
		}
		headerValueString = strings.TrimSuffix(headerValueString, ",")

		lowerHeaderName := strings.ToLower(headerName)

		headerNames += lowerHeaderName + " "
		signaturePayload += lowerHeaderName + ": " + headerValueString + "\n"
	}

	headerNames = strings.TrimSuffix(headerNames, " ")
	signaturePayload = strings.TrimSuffix(signaturePayload, "\n")

	switch algorithm {
	case crypto.SHA512:
		signatureAlgorithm = "rsa-sha512"
	default:
		signatureAlgorithm = "rsa-sha256"
	}

	signedPayload := encode.Sign(signaturePayload, algorithm, privateKey)

	signature := fmt.Sprintf("keyId=\"SN=%s\",algorithm=\"%s\",header\"%s\",signature=\"%s\"", keyId, signatureAlgorithm, headerNames, signedPayload)

	header.Add("Signature", signature)
}
