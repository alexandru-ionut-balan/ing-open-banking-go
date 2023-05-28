package psd2

import (
	"github.com/alexandru-ionut-balan/ing-open-banking-go/clients"
	"github.com/alexandru-ionut-balan/ing-open-banking-go/core/environment"
)

type AccountInformationClient struct {
	*clients.IngClient
}

func MakeAisClient(certificatePath string, privateKeyPath string, environment environment.Environment) *AccountInformationClient {
	return &AccountInformationClient{
		IngClient: clients.MakeIngClient(certificatePath, privateKeyPath, environment),
	}
}
