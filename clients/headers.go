package clients

import (
	"crypto"
	"net/http"
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
