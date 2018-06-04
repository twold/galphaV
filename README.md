# galphav

Golang package for hitting the alphavantage API

## Usage

WIP

```
go run main.go -apikey=$AVAPIKEY -symbol=MSFT
```


## Output


### TIME_SERIES_DAILY

raw output sample

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
