# galphav

Golang package for hitting the [alphavantage apis](https://www.alphavantage.co/documentation/)

## Prerequisites

You will need an [alphavantage api key](https://www.alphavantage.co/support/#support).

## Local Usage

1) Import package using the following command:

```
go get -u "github.com/twold/galphaV/api"
```

2) Build configuration file and store in the input folder:

```
wget -O $GOPATH/src/github.com/twold/galphaV/data/input/SP500.json https://pkgstore.datahub.io/core/s-and-p-500-companies/constituents_json/data/64dd3e9582b936b0352fdd826ecd3c95/constituents_json.json
```

The input file has an input struct with a slice of type [Config](https://godoc.org/github.com/twold/galphaV/api#Config) and is required at this time.  The 'Sector' and 'Symbol' are used as optional filter inputs. You may need to normalize the symbols as they should not contain any special characters like *BRK^B*.

Note: I used a basic S&P500 list.

Sample:

```
[
    { 
        "Name": "3M Company",
        "Sector": "Industrials",
        "Symbol": "MMM"
    },
    { 
        "Name": "Microsoft Corp.",
        "Sector": "Information Technology",
        "Symbol": "MSFT"
    }
]
```


3) Create output folder:

```
mkdir $GOPATH/src/github.com/twold/galphaV/data/input/
```

Note: current functionality requires input and output folders be in same "data" folder.


4) Configure environment variables and execute main:

```
// sample input where $AVAPIKEY is your api key and $GOPATH/src/github.com/twold/galphaV/data
// is where you have input file and is desired output location

// to pull all data
// go run main.go -apikey=$AVAPIKEY -path=$GOPATH/src/github.com/twold/galphaV/data

// to pull a sector
// go run main.go -apikey=$AVAPIKEY -path=$GOPATH/src/github.com/twold/galphaV/data -sector=Financials

// to pull a single equity
// go run main.go -apikey=$AVAPIKEY -path=$GOPATH/src/github.com/twold/galphaV/data -symbol=MSFT

// to pull all results for a single equity
// go run main.go -apikey=$AVAPIKEY -path=$GOPATH/src/github.com/twold/galphaV/data -symbol=GE -full=true

```

Your output will be saved in the data directory using the following folder structure with the symbol and date identifying the underlying timeseries data in the corresponding file.

structure:

```
data/output/{Symbol}/{YYYY-MM-DD}.json
```

example:

```
data/output/MSFT/2018-06-01.json
```

## Output

### TIME_SERIES_DAILY

raw api output sample:

```
{
    "Meta Data": {
        "1. Information": "Daily Prices (open, high, low, close) and Volumes",
        "2. Symbol": "MSFT",
        "3. Last Refreshed": "2018-06-01",
        "4. Output Size": "Compact",
        "5. Time Zone": "US/Eastern"
    },
    "Time Series (Daily)": {
        "2018-06-01": {
            "1. open": "99.2798",
            "2. high": "100.8600",
            "3. low": "99.1700",
            "4. close": "100.7900",
            "5. volume": "28655624"
        },
        "2018-05-31": {
            "1. open": "99.2900",
            "2. high": "99.9900",
            "3. low": "98.6100",
            "4. close": "98.8400",
            "5. volume": "34140891"
        },
        "2018-05-30": {
            "1. open": "98.3100",
            "2. high": "99.2500",
            "3. low": "97.9100",
            "4. close": "98.9500",
            "5. volume": "22158528"
        }
	}
}
```

Sample file above would contain the following:

```
{
    "1. open": "99.2798",
    "2. high": "100.8600",
    "3. low": "99.1700",
    "4. close": "100.7900",
    "5. volume": "28655624"
}
```
