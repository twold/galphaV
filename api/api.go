package api

import (
	"github.com/twold/galphaV/client"
)

type Service struct {
	*client.Client
	Error error
}

func New(key *string) *Service {
	return &Service{
		Client: client.New().Auth(key),
	}
}

func (c *Service) Daily() (Fetcher, error) {
	c.Client, c.Error = c.Function("TIME_SERIES_DAILY")
	if c.Error != nil {
		return nil, c.Error
	}
	return &Daily{
		Service: c,
	}, nil
}

type Fetcher interface {
	Get(opts ...string) ([]byte, error)
}

type Daily struct {
	*Service
}

func (c *Daily) Get(opts ...string) ([]byte, error) {
	resp, err := c.Symbols(opts...).Do("GET")
	if err != nil {
		return nil, err
	}
	return read(resp.Body)
}
