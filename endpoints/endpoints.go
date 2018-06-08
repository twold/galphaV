// package endpoints builds endpoint using default inputs and parameters from client and request pacakges.
package endpoints

import (
	"errors"
	"fmt"
	"net/url"
)

var (
	errInvalidService = errors.New("Invalid service request.")
)

const (
	defaultHost   = "www.alphavantage.co"
	defaultPath   = "query"
	defaultScheme = "https"
)

func defaultEndpoint() string {
	return fmt.Sprintf("%s://%s/%s", defaultScheme, defaultHost, defaultPath)
}

func New(values map[string][]string) (*url.URL, error) {
	// build default endpoint to parse into url.URL in next step
	e := defaultEndpoint()

	// parse into URL struct
	u, err := url.Parse(e)
	if err != nil {
		return nil, err
	}

	// set
	u.Scheme = defaultScheme
	u.Host = defaultHost
	u.Path = defaultPath

	v := url.Values{}
	for key, vals := range values {
		for _, val := range vals {
			v.Add(key, val)
		}
	}
	u.RawQuery = v.Encode()

	return u, nil
}
