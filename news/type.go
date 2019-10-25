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

// Return url for fetch lasts results newses
func Results() news {
	return news{"http://biznes.pap.pl/pl/rss/6639"}
}

// Return url for fetch lasts challenges newses
func Challenge() news {
	return news{"http://biznes.pap.pl/pl/rss/6638"}
}

// Return url for fetch lasts recommendations newses
func Recommendations() news {
	return news{"http://biznes.pap.pl/pl/rss/6634"}
}

// Return url for fetch lasts ebi newses
func Ebi() news {
	return news{"http://biznes.pap.pl/pl/rss/6612"}
}

// Return url for fetch lasts espi newses
func Espi() news {
	return news{"http://biznes.pap.pl/pl/rss/6614"}
}

// Return url for fetch lasts info newses
func Info() news {
	return news{"http://biznes.pap.pl/pl/rss/6600"}
}
