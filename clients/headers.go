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
	"github.com/gofrs/uuid"
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
	signature := getSignature(header, keyId, algorithm, privateKey)

	header.Add("Signature", signature)
}

func RequestTarget(header http.Header, method string, endpoint string) {
	header.Add("(request-target)", strings.ToLower(method)+" "+endpoint)
}

func RequestId(header http.Header) {
	header.Add("X-Request-ID", uuid.Must(uuid.NewV4()).String())
}

func Bearer(header http.Header, bearerToken string) {
	header.Add("Authorization", "Bearer "+bearerToken)
}

func AuthorizationSignature(header http.Header, keyId string, algorithm crypto.Hash, privateKey *rsa.PrivateKey) {
	signature := getSignature(header, keyId, algorithm, privateKey)

	header.Add("Authorization", "Signature "+signature)
}

func ContentTypeForm(header http.Header) {
	header.Add("Content-Type", "application/x-www-form-urlencoded")
}

func ContentTypeJson(header http.Header) {
	header.Add("Content-Type", "application/json")
}

func AccepJson(header http.Header) {
	header.Add("Accept", "application/json")
}

func getSignature(header http.Header, keyId string, algorithm crypto.Hash, privateKey *rsa.PrivateKey) string {
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

	return signature
}
