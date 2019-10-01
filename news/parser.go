package news

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ParseLast(newsType rssType) Rss {
	resp, err := http.Get(newsType.Url())
	if err != nil {
		fmt.Println("test")
	}
	body, err := ioutil.ReadAll(resp.Body)

	var rss Rss
	xml.Unmarshal(body, &rss)
	return rss
}

type rssType interface {
	Url() string
}

type Rss struct {
	Type string `xml:"channel>title"`
	Items []News `xml:"channel>item"`
}

type News struct {
	Title string `xml:"title"`
	Link string `xml:"link"`
	Content string `xml:"encoded"`
	Date string `xml:"pubDate"`
}