package clients

import (
	"crypto/tls"
	"net/http"

	"github.com/alexandru-ionut-balan/ing-open-banking-go/core/log"
)

type IngClient struct {
	*http.Client
}

type IngClientType uint8

const (
	Psd2 IngClientType = iota
	Premium
)

func MakeIngClient(environment Environment, clientType IngClientType) *IngClient {
	var cert tls.Certificate

	switch clientType {
	case Psd2:
		c, err := environment.Psd2TlsCertificates.LoadKeyPair()
		if err != nil {
			log.Error.Fatalln(err)
		}
		cert = c
	case Premium:
		c, err := environment.PremiumTlsCertificates.LoadKeyPair()
		if err != nil {
			log.Error.Fatalln(err)
		}
		cert = c
	}

	return &IngClient{
		Client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					Certificates: []tls.Certificate{cert},
				},
			},
		},
	}
}
