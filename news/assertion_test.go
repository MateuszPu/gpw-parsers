package news

import "testing"

func assertThat(news News, testing *testing.T) newsAssertion {
	return newsAssertion{news, testing}
}

type newsAssertion struct {
	news    News
	testing *testing.T
}

func (assertion newsAssertion) hasTitle(title string) newsAssertion {
	if assertion.news.Title != title {
		assertion.testing.Errorf("Wrong title for news %s != %s", assertion.news.Title, title)
	}
	return assertion
}

func (assertion newsAssertion) hasLink(link string) newsAssertion{
	if assertion.news.Link != link {
		assertion.testing.Errorf("Wrong link for news %s != %s", assertion.news.Link, link)
	}
	return assertion
}

func (assertion newsAssertion) hasContent(content string) newsAssertion {
	if assertion.news.Content != content {
		assertion.testing.Errorf("Wrong content for news %s != %s", assertion.news.Content, content)
	}
	return assertion
}

func (assertion newsAssertion) hasDate(date string) newsAssertion {
	if assertion.news.Date != date {
		assertion.testing.Errorf("Wrong date for news %s != %s", assertion.news.Date, date)
	}
	return assertion
}