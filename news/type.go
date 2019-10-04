package news

type rssType interface {
	Url() string
}

type news struct {
	rssUrl string
}

func (n news) Url() string {
	return n.rssUrl
}

func Results() news{
	return news{"http://biznes.pap.pl/pl/rss/6639"}
}

func Challenge() news {
	return news{"http://biznes.pap.pl/pl/rss/6638"}
}

func Recommendations() news {
	return news{"http://biznes.pap.pl/pl/rss/6634"}
}

func Ebi() news {
	return news{"http://biznes.pap.pl/pl/rss/6612"}
}

func Espi() news {
	return news{"http://biznes.pap.pl/pl/rss/6614"}
}

func Info() news {
	return news{"http://biznes.pap.pl/pl/rss/6600"}
}