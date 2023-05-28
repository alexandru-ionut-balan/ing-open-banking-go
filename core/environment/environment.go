package environment

import "net/url"

type Environment struct {
	BaseUrl url.URL
}

var (
	Sandbox Environment = Environment{
		BaseUrl: url.URL{
			Scheme: "https",
			Host:   "api.sandbox.ing.com",
		},
	}

	Production Environment = Environment{
		BaseUrl: url.URL{
			Scheme: "https",
			Host:   "api.ing.com",
		},
	}
)
