// package client contains interface used to build and validate client requests.
package client

import (
	"errors"
	"fmt"

	"github.com/twold/galphaV/endpoints"
	"github.com/twold/galphaV/request"
)

var (
	defaultDatatype   = "json"
	defaultOutputsize = "compact"
)

type ClientInfo struct {
	APIKey     *string
	datatype   *string
	function   *string
	interval   *string
	outputsize *string
}

type Client struct {
	ClientInfo
	Inputs map[string][]string
	*request.Request
}

type Builder interface {
	Auth(key *string) *Client
	DataType(datatype string) (*Client, error)
	Function(function string) (*Client, error)
	OutputSize(outputsize string) (*Client, error)
	Symbols(symbols ...string) *Client
	Do(method string) (*request.Request, error)
}

// create new client with default settings
func New() *Client {
	return &Client{
		ClientInfo: ClientInfo{
			// api defaults
			datatype:   &defaultDatatype,
			outputsize: &defaultOutputsize,
		},
	}
}

// Required
// add API Key to receiver
func (c *Client) Auth(key *string) *Client {
	c.APIKey = key
	return c
}

// Optional
// include option if you would like to change default data type
func (c *Client) DataType(datatype string) (*Client, error) {
	// validate input
	switch datatype {
	// default
	case "json":
	case "csv":
	default:
		// throw error
		msg := fmt.Sprintf("Invalid response format requested: %s.\n", datatype)
		return nil, errors.New(msg)
	}
	// add to receiver
	c.datatype = &datatype
	return c, nil
}

// Required
func (c *Client) Function(function string) (*Client, error) {
	// validate input
	switch function {

	// Stock Time Series Data API Parameters
	case "TIME_SERIES_INTRADAY":
	case "TIME_SERIES_DAILY":
	case "TIME_SERIES_DAILY_ADJUSTED":
	case "TIME_SERIES_WEEKLY":
	case "TIME_SERIES_WEEKLY_ADJUSTED":
	case "TIME_SERIES_MONTHLY":
	case "TIME_SERIES_MONTHLY_ADJUSTED":
	case "BATCH_STOCK_QUOTES":

	// Error
	default:
		// throw error
		msg := fmt.Sprintf("Invalid API funtion requested: %s.\n", function)
		return nil, errors.New(msg)
	}
	// add to receiver
	c.function = &function
	return c, nil
}

// Optional
// include option if you would like to change default response size
func (c *Client) OutputSize(outputsize string) (*Client, error) {
	// validate input
	switch outputsize {
	// default
	case "compact":
	case "full":
	default:
		// throw error
		msg := fmt.Sprintf("Invalid outputsize parameter %s requested.\n", outputsize)
		return nil, errors.New(msg)
	}
	// add to receiver
	c.outputsize = &outputsize
	return c, nil
}

// Optional
//
func (c *Client) Symbols(symbols ...string) *Client {
	if c.Inputs == nil {
		c.Inputs = make(map[string][]string)
	}
	if len(symbols) == 1 {
		c.Inputs["symbol"] = symbols
	} else {
		c.Inputs["symbol"] = symbols
	}
	return c
}

func (c *Client) Do(method string) (*request.Request, error) {
	values := make(map[string][]string)

	// Required input
	if c.function == nil {
		return nil, errors.New("Please include API function for request using client.Function() method.\n")
	}
	values["function"] = []string{*c.function}

	// Required input
	if c.APIKey == nil {
		return nil, errors.New("Please include API key in request using client.Auth() method.\n")
	}
	values["apikey"] = []string{*c.APIKey}

	// Optional input
	if c.outputsize != nil {
		values["outputsize"] = []string{*c.outputsize}
	}

	// Optional input
	if c.datatype != nil {
		values["datatype"] = []string{*c.datatype}
	}

	// Optional input
	if c.interval != nil {
		values["interval"] = []string{*c.interval}
	}

	// Optional input
	if c.outputsize != nil {
		values["outputsize"] = []string{*c.outputsize}
	}

	// are inputs required for all functions?
	// // Required input for all API funcs
	if c.Inputs == nil {
		return nil, errors.New("Please add input values to your request using one of the input methods Symbols(symbol ...string).\n")
	}

	// add input values to query
	for key, vals := range c.Inputs {
		values[key] = vals
	}

	// build api enpoint
	url, err := endpoints.New(values)
	if err != nil {
		return nil, err
	}

	// send request
	c.Request = request.New(method, url.String(), nil)
	if c.Error != nil {
		return nil, c.Error
	}
	return c.Request, nil
}
