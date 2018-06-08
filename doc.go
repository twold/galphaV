// Golang package for hitting the alphavantage APIs
//
// The primary interfaces for accessing API functions are found in the api
// package. Lower level operations including auth, input validation, and
// endpoint configurations can be found in the client, endpoints and request
// packages.
//
// Package API contains all supported operations. ServicerAPI interface methods
// are named and organizated using a combination of the HTTP method and
// the underlying endpoint. For example, "GET" requests to the TIME_SERIES_DAILY_ADJUSTED
// api https://www.alphavantage.co/documentation/#dailyadj endpoint will be
// created using the GetTimeSeriesDaily method once the *Service receiver
// has been created using the desired input parameters.
package galphaV
