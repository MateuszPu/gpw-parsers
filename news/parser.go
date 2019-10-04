package news

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

func ParseLast(newsType rssType) (*Rss, error) {
	rss := &Rss{}
	url := newsType.Url()

	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(body, &rss)

	return rss, err
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