package stock

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

//func TestParseReal(testing *testing.T) {
//	//given
//	//when
//	stocks, err := ParseTodayStocksDetails()
//
//	//then
//	if err != nil || len(stocks) == 0 {
//		testing.Errorf("Ups something was wrong with parsing")
//	}
//}

func TestParseLast(testing *testing.T) {
	//given
	server := mockRssServer("stock_response_%d.html")
	defer server.Close()

	//when
	url := server.URL + "/%d"
	stocks, err := parseTodayStocksDetails(url)

	//then
	if err != nil || len(stocks) == 0 {
		testing.Errorf("Ups something was wrong with parsing")
	}
	assertThat(stocks[0], testing).hasTicker("06N").hasName("06MAGNA").hasPrice("0.2000")
	assertThat(stocks[1], testing).hasTicker("08N").hasName("08OCTAVA").hasPrice("0.860")
	assertThat(stocks[2], testing).hasTicker("11B").hasName("11BIT").hasPrice("384.0")
	assertThat(stocks[3], testing).hasTicker("DCR").hasName("DECORA").hasPrice("17.50")
}

func BenchmarkParseLast(b *testing.B) {
	server := mockRssServer("stock_response_1.html")
	defer server.Close()
	for i := 0; i < b.N; i++ {
		parseFrom(server.URL + "/1", func(){})
	}
}

func mockRssServer(path string) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/1", func(w http.ResponseWriter, r *http.Request) {
		bytes, _ := ioutil.ReadFile(fmt.Sprintf(path, 1))
		_, _ = w.Write(bytes)
	})
	mux.HandleFunc("/2", func(w http.ResponseWriter, r *http.Request) {
		bytes, _ := ioutil.ReadFile(fmt.Sprintf(path, 2))
		_, _ = w.Write(bytes)
	})
	return httptest.NewServer(mux)
}
