package stock

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"strings"
)

func ParseLast(url string) ([]Stock, error) {
	var stocks []Stock
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
		stock := Stock{ticker, name, price}
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

type Stock struct {
	ticker string
	name string
	price string
}
