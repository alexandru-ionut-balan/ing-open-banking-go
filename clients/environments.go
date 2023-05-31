package clients

import (
	"crypto/tls"
	"net/url"
)

type Environment struct {
	BaseUrl                   url.URL
	Psd2TlsCertificates       CertificateBundle
	Psd2SigningCertificates   CertificateBundle
	PremiumTlsCertificates    CertificateBundle
	PremiumSigningCertificate CertificateBundle
}

type CertificateBundle struct {
	CertificatePath string
	PrivateKeyPath  string
}

var (
	Sandbox Environment = Environment{
		BaseUrl: url.URL{
			Scheme: "https",
			Host:   "api.sandbox.ing.com",
		},
		Psd2TlsCertificates: CertificateBundle{
			CertificatePath: "certs/psd2/example_client_tls.cer",
			PrivateKeyPath:  "certs/psd2/example_client_tls.key",
		},
		Psd2SigningCertificates: CertificateBundle{
			CertificatePath: "certs/psd2/example_client_signing.cer",
			PrivateKeyPath:  "certs/psd2/example_client_signing.cer",
		},
	}

	Production Environment = Environment{
		BaseUrl: url.URL{
			Scheme: "https",
			Host:   "api.ing.com",
		},
	}
)

func (bundle *CertificateBundle) LoadKeyPair() (tls.Certificate, error) {
	return tls.LoadX509KeyPair(bundle.CertificatePath, bundle.PrivateKeyPath)
}
