package news

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"
)

func ParseLast(newsType rssType) (*Rss, error) {
	url := newsType.Url()
	body, err := open(url)
	defer body.Close()

	return parse(read(body, err))
}

func open(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, err
}

func read(body io.ReadCloser, err error) (*[]byte, error) {
	content, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return &content, err
}

func parse(content *[]byte, err error) (*Rss, error) {
	rss := &Rss{}
	err = xml.Unmarshal(*content, &rss)
	return rss, err
}

type Rss struct {
	Type  string `xml:"channel>title"`
	Items []News `xml:"channel>item"`
}

type News struct {
	Title   string `xml:"title"`
	Link    string `xml:"link"`
	Content string `xml:"encoded"`
	Date    string `xml:"pubDate"`
}
