package news

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func LastNews(newsType rssType) Rss {
	resp, err := http.Get(newsType.url())
	if err != nil {
		fmt.Println("test")
	}
	body, err := ioutil.ReadAll(resp.Body)

	var rss Rss
	xml.Unmarshal(body, &rss)
	return rss
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