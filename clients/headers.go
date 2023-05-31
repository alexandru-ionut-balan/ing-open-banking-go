package clients

import (
	"crypto"
	"net/http"

	"github.com/alexandru-ionut-balan/ing-open-banking-go/core"
	"github.com/alexandru-ionut-balan/ing-open-banking-go/encode"
)

func Digest(header *http.Header, grant core.Grant, algorithm crypto.Hash) {
	header.Add("Digest", encode.Base64(encode.Hash("grant-type="+string(grant), algorithm)))
}
