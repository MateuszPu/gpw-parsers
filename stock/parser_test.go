package stock

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParseLast(testing *testing.T) {
	//given
	server := mockRssServer("stock_response.html")
	defer server.Close()

	//when
	stocks, err := parseTodayStocks(server.URL)

	//then
	if err != nil || len(stocks) == 0 {
		testing.Errorf("Ups something was wrong with parsing")
	}
	assertThat(stocks[0], testing).hasTicker("06N").hasName("06MAGNA").hasPrice("0.2000")
	assertThat(stocks[1], testing).hasTicker("08N").hasName("08OCTAVA").hasPrice("0.860")
	assertThat(stocks[2], testing).hasTicker("11B").hasName("11BIT").hasPrice("384.0")
	assertThat(stocks[99], testing).hasTicker("DCR").hasName("DECORA").hasPrice("17.50")
}

func BenchmarkParseLast(b *testing.B) {
	server := mockRssServer("stock_response.html")
	defer server.Close()
	for i:=0; i < b.N; i++ {
		_, _ = parseTodayStocks(server.URL)
	}
}

func mockRssServer(path string) (*httptest.Server) {
	bytes, _ := ioutil.ReadFile(path)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write(bytes)
	}))
	return server
}
