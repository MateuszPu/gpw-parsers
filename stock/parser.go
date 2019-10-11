package stock

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
	"sync"
)

const BASE_STOOQ_URL = "https://stooq.pl/t/?i=513&v=0&l=%d"

func ParseTodayStocksDetails() ([]StockDetails, error) {
	return parseTodayStocksDetails(BASE_STOOQ_URL)
}

func parseTodayStocksDetails(pathUrl string) ([]StockDetails, error) {
	var wg sync.WaitGroup
	var result []StockDetails
	page := 1
	for {
		wg.Add(1)
		onExit := func() { wg.Done() }
		parsedStockDetails := parseFrom(fmt.Sprintf(pathUrl, page), onExit)
		if len(parsedStockDetails) == 0 {
			break
		}
		result = append(result, parsedStockDetails...)
		page++
	}
	wg.Wait()
	return result, nil
}

func parseFrom(url string, onExit func()) []StockDetails {
	defer onExit()
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		panic(err)
	}
	stocksDetails := parseStockDetailsFrom(doc)
	return stocksDetails
}

func parseStockDetailsFrom(doc *goquery.Document) []StockDetails {
	var stocks []StockDetails
	doc.Find("tr[id^='r_']").Each(func(i int, s *goquery.Selection) {
		ticker := strings.TrimSpace(s.Children().Eq(0).Text())
		name := strings.TrimSpace(s.Children().Eq(1).Text())
		price := strings.TrimSpace(s.Children().Eq(2).Text())
		stock := StockDetails{ticker, name, price}
		stocks = append(stocks, stock)
	})
	return stocks
}

type StockDetails struct {
	ticker string
	name   string
	price  string
}
