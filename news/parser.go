package news

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	logger "github.com/Sirupsen/logrus"
)

func ParseLast(newsType rssType) (parsedNews *Rss, err error) {
	rss := &Rss{}
	url := newsType.Url()
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		logger.Fatalf("Could not open following url %s : %s", url, err)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Fatalf("Error during reading content of url: %s", err)
		return nil, err
	}

	err = xml.Unmarshal(body, &rss)
	if err != nil {
		logger.Fatalf("Error during parsing: %s", err)
		return nil, err
	}
	return rss, err
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