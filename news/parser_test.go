package news

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParseLastWithR(testing *testing.T) {
	//given
	sources := [...]rssType{Results(), Challenge(), Recommendations(), Ebi(), Espi(), Info()}

	for _, value := range sources {
		parsedNews, err := ParseLast(value)
		if err != nil || parsedNews == nil {
			testing.Errorf("Ups something was wrong %s", err)
		}
	}
}

func TestParseLast(testing *testing.T) {
	//given
	server, rssSource := mockRssServer("rss_response")
	defer server.Close()

	//when
	parsedNews, err := ParseLast(rssSource)

	//then
	if err != nil || parsedNews == nil {
		testing.Errorf("Ups something was wrong %s", err)
	}
	assertThat(parsedNews.Items[0], testing).hasTitle("Title first").hasLink("Link first").hasContent("content first").hasDate("Thu, 03 Oct 2019 13:56:45 GMT")
}

func mockRssServer(path string) (*httptest.Server, rssSource) {
	bytes, _ := ioutil.ReadFile(path)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write(bytes)
	}))
	return server, rssSource{server.URL}
}

type rssSource struct {
	url string
}

func (m rssSource) Url() string {
	return m.url
}

