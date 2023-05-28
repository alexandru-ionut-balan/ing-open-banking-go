package headers

import (
	"net/http"
	"time"

	"github.com/alexandru-ionut-balan/ing-open-banking-go/core"
)

type Headers struct {
	*http.Header
}

func (headers *Headers) AcceptJson() *Headers {
	headers.Add("Accept", "application/json")
	return headers
}

func (headers *Headers) ContentJson() *Headers {
	headers.Add("Content-Type", "application/json")
	return headers
}

func (headers *Headers) ContentFormEncoded() *Headers {
	headers.Add("Content-Type", "application/x-www-form-urlencoded")
	return headers
}

func (headers *Headers) Digest(payload string, algorithm core.HashingAlgorithm) *Headers {
	headers.Add("Digest", core.Digest(payload, algorithm))
	return headers
}

func (headers *Headers) Date() *Headers {
	location, _ := time.LoadLocation("GMT")
	date := time.Now().In(location).Format(time.RFC1123)

	headers.Add("Date", date)
	return headers
}
