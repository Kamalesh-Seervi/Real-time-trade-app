package core

import "strings"

var tickerSet map[string]struct{}

func GetAllTickers() []string {
	tickerList := []string{}
	for key := range tickerSet {
		tickerList = append(tickerList, key)
	}
	return tickerList
}

func IsTickerAllowed(ticker string) bool {
	_, ok := tickerSet[strings.ToLower(ticker)]
	return ok
}

func LoadTikers(tickers []string) {
	if tickerSet == nil {
		tickerSet = make(map[string]struct{})
	}
	for _, t := range tickers {
		tickerSet[strings.ToLower(strings.Trim(strings.Trim(t, "\\"), "\""))] = struct{}{}
	}
}
