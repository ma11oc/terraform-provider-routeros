package routeros

import (
	rc "github.com/ma11oc/go-routerosclient"
)

// Config ...
// FIXME: write comment
type Config struct {
	address  string
	username string
	password string
	async    bool
}

// Client ...
// FIXME: write comment
func (c *Config) Client() (*rc.Client, error) {

	conf := &rc.Config{
		Address:  c.address,
		Username: c.username,
		Password: c.password,
		Async:    c.async,
	}

	if RouterOSClient == nil {
		var err error

		if RouterOSClient, err = rc.NewClient(conf); err != nil {
			return nil, err
		}
	}

	return RouterOSClient, nil
}
