package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/twold/galphaV/api"
)

var (
	apikey, symbol string
)

func init() {
	// Add api key
	flag.StringVar(&apikey, "apikey", "$AVAPIKEY", "-apikey=xyzABCD1234567890 add api key for data pull")
	// Optional input param.  This overwrites input file with multiple symbols if given.
	flag.StringVar(&symbol, "symbol", "", "-symbol=FB input ticker symbol to retrieve data set")
}

// go run main.go -apikey=$AVAPIKEY -symbol=MSFT

func main() {
	flag.Parse()

	// create service using input format and data type
	svc := api.New(&apikey)

	f, err := svc.Daily()
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := f.Get(symbol)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v\n", string(resp))

}
