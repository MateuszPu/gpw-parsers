package stock

import "testing"

func assertThat(actual Stock, testing *testing.T) stockAssertion {
	return stockAssertion{actual, testing}
}

type stockAssertion struct {
	actual  Stock
	testing *testing.T
}

func (assertion stockAssertion) hasName(name string) stockAssertion {
	if assertion.actual.name != name {
		assertion.testing.Errorf("Wrong name for stock expected %s to be %s", assertion.actual.name, name)
	}
	return assertion
}

func (assertion stockAssertion) hasTicker(ticker string) stockAssertion {
	if assertion.actual.ticker != ticker {
		assertion.testing.Errorf("Wrong ticker for stock expected %s to be %s", assertion.actual.ticker, ticker)
	}
	return assertion
}

func (assertion stockAssertion) hasPrice(price string) stockAssertion {
	if assertion.actual.price != price {
		assertion.testing.Errorf("Wrong price for stock expected %s to be %s", assertion.actual.price, price)
	}
	return assertion
}
