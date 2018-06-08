// package api contains primary application interface.
package api

import (
	"github.com/twold/galphaV/client"
	"github.com/twold/galphaV/request"
)

type ServicerAPI interface {
	GetTimeSeriesDaily(symbols ...string) ([]*TimeSeriesDaily, error)

	GetTimeSeriesDailyAdjusted(symbols ...string) ([]*TimeSeriesDailyAdjusted, error)
}

type Config struct {
	Name string `json:"Name"`

	Sector string `json:"Sector"`

	Symbol string `json:"Symbol"`
}

type Service struct {
	*client.Client
	Error error
}

func New(key *string) *Service {
	return &Service{
		Client: client.New().Auth(key),
	}
}

func ReadConfig(path, filename string, filter *string) ([]string, error) {
	b, err := readFile(path, filename)
	if err != nil {
		return nil, err
	}
	var items []Config
	err = unmarshal(b, &items)
	if err != nil {
		return nil, err
	}

	symbols := make([]string, 0)
	for _, item := range items {
		if filter == nil {
			symbols = append(symbols, item.Symbol)
			continue
		}

		// if sector is not specified or if all sectors are requested
		if item.Sector == *filter {
			symbols = append(symbols, item.Symbol)
			continue
		}

		// match if sector is specified
		if item.Symbol == *filter {
			symbols = append(symbols, item.Symbol)
		}
	}
	return symbols, nil
}

func (c *Service) Full() *Service {
	c.Client, c.Error = c.OutputSize("full")
	return c
}

type MetaData struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	OutputSize    string `json:"4. Output Size"`
	TimeZone      string `json:"5. Time Zone"`
}

type TimeSeriesDaily struct {
	Date string `json:"date" name"date"`

	DayOfWeek string `json:"dayOfWeek" name"dayOfWeek"`

	Open string `json:"1. open"`

	High string `json:"2. high"`

	Low string `json:"3. low"`

	Close string `json:"4. close"`

	Volume string `json:"5. volume"`
}

func (c *Service) GetTimeSeriesDaily(symbols ...string) ([]*TimeSeriesDaily, error) {
	c.Client, c.Error = c.Function("TIME_SERIES_DAILY")
	if c.Error != nil {
		return nil, c.Error
	}
	// execute call to API
	resp, err := get(c.Symbols(symbols...))
	if err != nil {
		return nil, err
	}

	output := new(TimeSeriesDailyResponse)

	// read response body
	err = resp.Read(output)
	if err != nil {
		return nil, err
	}

	return output.Data, nil
}

type TimeSeriesDailyAdjusted struct {
	Date string `json:"date" name"date"`

	DayOfWeek string `json:"dayOfWeek" name"dayOfWeek"`

	Open string `json:"1. open"`

	High string `json:"2. high"`

	Low string `json:"3. low"`

	Close string `json:"4. close"`

	AdjustedClose string `json:"5. adjusted close"`

	Volume string `json:"6. volume"`

	DividendAmount string `json:"7. dividend amount"`

	SplitCoefficient string `json:"8. split coefficient"`
}

func (c *Service) GetTimeSeriesDailyAdjusted(symbols ...string) ([]*TimeSeriesDailyAdjusted, error) {

	// Add function to receiver
	c.Client, c.Error = c.Function("TIME_SERIES_DAILY_ADJUSTED")
	if c.Error != nil {
		return nil, c.Error
	}
	// execute call to API
	resp, err := get(c.Symbols(symbols...))
	if err != nil {
		return nil, err
	}

	output := new(TimeSeriesDailyAdjustedResponse)

	// read response body
	err = resp.Read(output)
	if err != nil {
		return nil, err
	}

	return output.Data, nil
}

type Response struct {
	r   *request.Request
	byt []byte
}

func (c *Response) Read(resp Unmarshaler) error {
	byt, err := read(c)
	if err != nil {
		return err
	}
	return resp.Unmarshal(byt)
}

type Unmarshaler interface {
	Date(input string) error
	Unmarshal(data []byte) error
}

type TimeSeriesDailyResponse struct {
	MetaData `json:"Meta Data"`

	Raw map[string]interface{} `json:"Time Series (Daily)"`

	Data []*TimeSeriesDaily

	obj TimeSeriesDaily
}

func (c *TimeSeriesDailyResponse) Date(input string) error {
	c.obj.Date = input
	// add day of week
	v, err := dayOfWeek(input)
	if err != nil {
		return err
	}
	c.obj.DayOfWeek = *v
	return nil
}

func (c *TimeSeriesDailyResponse) Unmarshal(data []byte) error {

	// unmarshal and process generic response
	err := unmarshal(data, c)
	if err != nil {
		return err
	}

	for key, val := range c.Raw {
		obj := c.obj
		// marshal response
		byt, err := marshal(val)
		if err != nil {
			return err
		}

		// marshal to new struct
		err = unmarshal(byt, &obj)
		if err != nil {
			return err
		}

		err = c.Date(key)
		if err != nil {
			return err
		}

		// append pointer to slice for response
		c.Data = append(c.Data, &obj)
	}
	return nil
}

type TimeSeriesDailyAdjustedResponse struct {
	MetaData `json:"Meta Data"`

	Raw map[string]interface{} `json:"Time Series (Daily)"`

	Data []*TimeSeriesDailyAdjusted

	obj TimeSeriesDailyAdjusted
}

func (c *TimeSeriesDailyAdjustedResponse) Date(input string) error {
	c.obj.Date = input
	// add day of week
	v, err := dayOfWeek(input)
	if err != nil {
		return err
	}
	c.obj.DayOfWeek = *v
	return nil
}

func (c *TimeSeriesDailyAdjustedResponse) Unmarshal(data []byte) error {

	// unmarshal and process generic response
	err := unmarshal(data, c)
	if err != nil {
		return err
	}

	for key, val := range c.Raw {
		obj := c.obj
		// marshal response
		byt, err := marshal(val)
		if err != nil {
			return err
		}

		// marshal to new struct
		err = unmarshal(byt, &obj)
		if err != nil {
			return err
		}

		err = c.Date(key)
		if err != nil {
			return err
		}

		// append pointer to slice for response
		c.Data = append(c.Data, &obj)
	}
	return nil
}

type Writer interface {
	S3Write(path, symbol string) error
	Write(path, symbol string) error
}

func (c *TimeSeriesDaily) S3Write(path, symbol string) error {
	return nil
}

func (c *TimeSeriesDaily) Write(path, symbol string) error {
	return writeFile(c.Date, path, symbol, *c)
}

func (c *TimeSeriesDailyAdjusted) S3Write(path, symbol string) error {
	return nil
}

func (c *TimeSeriesDailyAdjusted) Write(path, symbol string) error {
	return writeFile(c.Date, path, symbol, *c)
}
