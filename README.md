# Warsaw stock exchange parsers

## News parser
Parser get news from rss channels. List of available channels are available in news package. For example news.Results()
### Usage:  
To get last news about results.
    
    import github.com/MateuszPu/gpw-parsers/news
    news.ParseLast(news.Results())


## Stock details parser
Parser get stock details from stooq.pl site. 
### Usage:
To get las stocks details (open, close, high, low) price, date, ticker, name
    
    import github.com/MateuszPu/gpw-parsers/stock`
    stock.parsTodayStocksDetails()
    
