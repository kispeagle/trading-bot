package constant

import (
	"time"
)

const (
	BINANCE_WEBSOCKET_URL                    = "stream.binance.com:9443"
	PING_PERIOD                              = 5 * time.Minute
	PATTERN_KLINE_SERVICE                    = "%s@kline_%s"
	PATTERN_AGG_TRADE_SERVICE                = "%s@aggTrade"
	PATTERN_ALL_BOOK_TICKERS_SERVICE         = "!bookTicker"
	PATTERN_ALL_MINI_TICKERS_SERVICE         = "!miniTicker@arr"
	PATTERN_ALL_ROLLING_WINDOW_STATS_SERVICE = "!ticker_%s@arr"
	PATTERN_ALL_TICKERS_SERVICE              = "!ticker@arr"
	PATTERN_DIFF_DEPTH_SERVICE               = "%s@depth"
	PATTERNBOOK_TICKER_SERVICE               = "%s@bookTicker"
	PATTERN_MINI_TICKER_SERVICE              = "%s@miniTicker"
	PATTERN_ROLLING_WINDOW_STATS_SERVICE     = "%s@ticker_%s"
	PATTERN_TICKER_SERVICE                   = "%s@ticker"
	PATTERN_PARTIAL_BOOK_DEPTH_SERVICE       = "%s@depth%d"
	PATTERN_TRADE_SERVICE                    = "%s@trade"
)
