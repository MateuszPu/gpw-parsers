package stock

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

const baseStooqUrl = "https://stooq.pl/t/?i=513&v=0&l=%d"

// func to parse all stock details for todays date
// func should be call after warsaw stock exchange close
func ParseTodayStocksDetails() ([]StockDetails, error) {
	return parseTodayStocksDetails(baseStooqUrl)
}

func parseTodayStocksDetails(pathUrl string) ([]StockDetails, error) {
	var result []StockDetails
	page := 1
	for {
		parsedStockDetails := parseFrom(fmt.Sprintf(pathUrl, page))
		if len(parsedStockDetails) == 0 {
			break
		}
		result = append(result, parsedStockDetails...)
		page++
	}
	return result, nil
}

func parseFrom(url string) []StockDetails {
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
