package stock

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"strings"
)

const BASE_STOOQ_URL = "https://stooq.pl/t/?i=513&v=0&l=%d"

func ParseTodayStocksDetails() ([]StockDetails, error){
	return parseTodayStocksDetails(BASE_STOOQ_URL)
}

func parseTodayStocksDetails(pathUrl string) ([]StockDetails, error){
	var result []StockDetails
	page := 1
	url := fmt.Sprintf(pathUrl, page)
	stc, err := parseFrom(url)
	result = append(result, stc...)
	if err != nil {
		return nil, err
	}
	for len(stc) > 0 {
		page++
		url = fmt.Sprintf(pathUrl, page)
		stc, _ = parseFrom(url)
		result = append(result, stc...)
	}
	return result, nil
}

func parseFrom(url string) ([]StockDetails, error) {
	var stocks []StockDetails
	content, err := open(url)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(content)
	if err != nil {
		return nil, err
	}

	doc.Find("tr[id^='r_']").Each(func(i int, s *goquery.Selection) {
		ticker := strings.TrimSpace(s.Children().Eq(0).Text())
		name := strings.TrimSpace(s.Children().Eq(1).Text())
		price := strings.TrimSpace(s.Children().Eq(2).Text())
		stock := StockDetails{ticker, name, price}
		stocks = append(stocks, stock)
	})
	return stocks, nil
}


func open(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, err
}

type StockDetails struct {
	ticker string
	name string
	price string
}
