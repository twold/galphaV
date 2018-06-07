package main

import (
	"flag"
	"log"

	"github.com/twold/galphaV/api"
)

var (
	full bool

	apikey, configFile, path, sector, symbol string
)

func init() {
	// Add api key
	flag.StringVar(&apikey, "apikey", "$AVAPIKEY", "-apikey=xyzABCD1234567890 add api key for data pull")
	// Modified SP500 input file from
	// https://pkgstore.datahub.io/core/s-and-p-500-companies/constituents_json/data/64dd3e9582b936b0352fdd826ecd3c95/constituents_json.json
	flag.StringVar(&configFile, "configFile", "SP500.json", "-configFile=specify location of configuration file. Default is SP500.json")
	// Optional input param.
	flag.BoolVar(&full, "full", false, "-full=")
	// This is path to data files including input and output folders
	flag.StringVar(&path, "path", "", "-path=")
	// Optional input param.
	flag.StringVar(&sector, "sector", "", "-sector= input sector to retrieve data set")
	// Optional input param.
	flag.StringVar(&symbol, "symbol", "", "-symbol=FB input ticker symbol to retrieve data set")

}

// sample input where $QUANDLAPIKEY is your api key and $GOPATH/src/github.com/twold/go-quandl/data
// is where you have input file and is desired output location

// to pull all data
// go run main.go -apikey=$AVAPIKEY -path=$GOPATH/src/github.com/twold/galphaV/data

// to pull a sector
// go run main.go -apikey=$AVAPIKEY -path=$GOPATH/src/github.com/twold/galphaV/data -sector=Financials

// to pull a single equity
// go run main.go -apikey=$AVAPIKEY -path=$GOPATH/src/github.com/twold/galphaV/data -symbol=MSFT

// to pull all results for a single equity
// go run main.go -apikey=$AVAPIKEY -path=$GOPATH/src/github.com/twold/galphaV/data -symbol=GE -full=true

//
func setFilter(sector, symbol string) *string {
	if sector != "" {
		return &sector
	}

	if symbol != "" {
		return &symbol
	}
	return nil
}

func main() {

	// parse input flags
	flag.Parse()

	// set filter based on inputs
	filter := setFilter(sector, symbol)

	// read config and select symbols based on inputs
	symbols, err := api.ReadConfig(path, configFile, filter)
	if err != nil {
		log.Fatalln(err)
	}

	// create service using input format and data type
	svc := api.New(&apikey)

	// retrieve full data set
	if full == true {
		svc = svc.Full()
	}

	for _, symbol := range symbols {
		log.Printf("Performing call for symbol: %+v\n", symbol)

		resp, err := svc.GetTimeSeriesDailyAdjusted(symbol)
		if err != nil {
			log.Fatalln(err)
		}

		for _, item := range resp {
			log.Printf("Writing data file for %s on %s.\n", symbol, item.Date)
			item.Write(path, symbol)
		}
	}
}
