package services

import ( 
	logger "github.com/kispeagle/go-binance/logs"
	"github.com/kispeagle/go-binance/stream"
	constant "github.com/kispeagle/trading-bot/constants"
	"fmt"
)


// var logger *zap.SugaredLogger = logger.Log
type Service interface {
	Call()
}

func CallKlineService(symbol string, c chan interface{}) {
	interval := "1m"
	streamName := stream.CreateStreamName(constant.PATTERN_KLINE_SERVICE, symbol, interval)
	url := stream.CreateUrl(streamName)
	fmt.Println(url)

	s := stream.NewKlineService(symbol)
	go s.Call(url, c)
}

func CallMiniTickerService(symbol string, c chan interface{}) {
	streamName := stream.CreateStreamName(constant.PATTERN_MINI_TICKER_SERVICE, symbol)
	url := stream.CreateUrl(streamName)
	s := stream.NewMiniTickerService(symbol)
	go s.Call(url, c)
}

func CallAllMiniTickerService(symbol string, c chan interface{}) {
	streamName := stream.CreateStreamName(constant.PATTERN_ALL_MINI_TICKERS_SERVICE)
	logger.Log.Debug(streamName)
	url := stream.CreateUrl(streamName)
	s := stream.NewAllMiniTickersService(symbol)
	go s.Call(url, c)
}

func CallAggTradeService(symbol string, c chan interface{}) {
	streamName := stream.CreateStreamName(constant.PATTERN_AGG_TRADE_SERVICE, symbol)
	url := stream.CreateUrl(streamName)
	logger.Log.Debug(url)
	s := stream.NewAggTradeService(symbol)
	s.Call(url, c)
}

func CallTradeService(symbol string, c chan interface{}) {
	streamName := stream.CreateStreamName(constant.PATTERN_TRADE_SERVICE, symbol)
	url := stream.CreateUrl(streamName)
	logger.Log.Debug(url)
	s := stream.NewTradeService(symbol)
	s.Call(url, c)
}

func CallAllTickersService(symbol string, c chan interface{}) {
	streamName := stream.CreateStreamName(constant.PATTERN_ALL_TICKERS_SERVICE)
	url := stream.CreateUrl(streamName)
	s := stream.NewAllTickersService(symbol)

	go s.Call(url, c)
}

func CallRollingWindowStatsService(symbol string, c chan interface{}) {
	window := "1h"
	streamName := stream.CreateStreamName(constant.PATTERN_ROLLING_WINDOW_STATS_SERVICE, symbol, window)
	url := stream.CreateUrl(streamName)
	logger.Log.Debug(url)
	s := stream.NewRollingWindowStatsService(symbol)
	go s.Call(url, c)
}

func CallAllRollingWindowStatsService(symbol string, c chan interface{}) {
	window := "1h"
	streamName := stream.CreateStreamName(constant.PATTERN_ALL_ROLLING_WINDOW_STATS_SERVICE, window)
	url := stream.CreateUrl(streamName)

	s := stream.NewAllRollingWindowStatsService(symbol)
	go s.Call(url, c)
}

func CallBookTickerService(symbol string, c chan interface{}) {
	streamName := stream.CreateStreamName(constant.PATTERNBOOK_TICKER_SERVICE, symbol)
	url := stream.CreateUrl(streamName)
	logger.Log.Debug(url)
	s := stream.NewBookTickerService(symbol)
	go s.Call(url, c)
}

func CallAllBookTickerService(symbol string, c chan interface{}) {
	streamName := stream.CreateStreamName(constant.PATTERN_ALL_BOOK_TICKERS_SERVICE)
	url := stream.CreateUrl(streamName)
	logger.Log.Debug(url)
	s := stream.NewAllBookTickerService(symbol)
	go s.Call(url, c)
}

func CallPartialBookDepthService(symbol string, c chan interface{}) {
	level := 5
	streamName := stream.CreateStreamName(constant.PATTERN_PARTIAL_BOOK_DEPTH_SERVICE, symbol, level)
	url := stream.CreateUrl(streamName)

	s := stream.NewPartialBookDepthService(symbol)
	go s.Call(url, c)
}

func CallDiffDepthService(symbol string, c chan interface{}) {
	streamName := stream.CreateStreamName(constant.PATTERN_DIFF_DEPTH_SERVICE, symbol)
	url := stream.CreateUrl(streamName)

	s := stream.NewDiffDepthService(symbol)
	go s.Call(url, c)
}