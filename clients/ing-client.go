package clients

import (
	"crypto/tls"
	"net/http"
	"net/url"

	"github.com/alexandru-ionut-balan/ing-open-banking-go/core/environment"
	"github.com/alexandru-ionut-balan/ing-open-banking-go/logger"
)

type IngClient struct {
	*http.Client
	BaseUrl url.URL
}

// MakeIngClient creates a new base Ing Client with TLS enabled. TLS configuration is mandatory.
func MakeIngClient(certificatePath string, privateKeyPath string, environment environment.Environment) *IngClient {
	clientCertificate, err := tls.LoadX509KeyPair(certificatePath, privateKeyPath)
	if err != nil {
		logger.Error(err)
		return nil
	}

	tlsTransport := http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates: []tls.Certificate{clientCertificate},
		},
	}

	return &IngClient{
		Client: &http.Client{
			Transport: &tlsTransport,
		},
		BaseUrl: environment.BaseUrl,
	}
}
