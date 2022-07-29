package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	logger "github.com/kispeagle/go-binance/logs"
	"github.com/kispeagle/go-binance/stream"
	constant "github.com/kispeagle/trading-bot/constants"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	Log := logger.Log
	symbol := "ookiusdt"
	symbol = strings.ToLower(symbol)

	c := make(chan interface{})
	go callKlineService(symbol, c)
	for x := range c {

		if x == nil {
			return
		}
		switch v := x.(type) {

		case stream.MiniTickerList:
			Log.Debug(v.List)

		default:
			Log.Debug(v)
		}
	}

}

func Println(v interface{}) error {
	fmt.Println(v)
	return nil
}

type Timestamp time.Time

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	v, err := time.Parse(time.RubyDate, string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	*t = Timestamp(v)
	return nil
}

func callKlineService(symbol string, c chan interface{}) {
	interval := "1m"
	streamName := stream.CreateStreamName(constant.PATTERN_KLINE_SERVICE, symbol, interval)
	url := stream.CreateUrl(streamName)
	fmt.Println(url)

	s := stream.NewKlineService(symbol)
	go s.Call(url, c)
}

func callMiniTickerService(symbol string, c chan interface{}) {
	streamName := stream.CreateStreamName(constant.PATTERN_MINI_TICKER_SERVICE, symbol)
	url := stream.CreateUrl(streamName)
	s := stream.NewMiniTickerService(symbol)
	go s.Call(url, c)
}

func callAllMiniTickerService(symbol string, c chan interface{}) {
	streamName := stream.CreateStreamName(constant.PATTERN_ALL_MINI_TICKERS_SERVICE)
	logger.Log.Debug(streamName)
	url := stream.CreateUrl(streamName)
	s := stream.NewAllMiniTickersService(symbol)
	go s.Call(url, c)
}

func callAggTradeService(symbol string, c chan interface{}) {
	streamName := stream.CreateStreamName(constant.PATTERN_AGG_TRADE_SERVICE, symbol)
	url := stream.CreateUrl(streamName)
	logger.Log.Debug(url)
	s := stream.NewAggTradeService(symbol)
	s.Call(url, c)
}

func callTradeService(symbol string, c chan interface{}) {
	streamName := stream.CreateStreamName(constant.PATTERN_TRADE_SERVICE, symbol)
	url := stream.CreateUrl(streamName)
	logger.Log.Debug(url)
	s := stream.NewTradeService(symbol)
	s.Call(url, c)
}

func callAllTickersService(symbol string, c chan interface{}) {
	streamName := stream.CreateStreamName(constant.PATTERN_ALL_TICKERS_SERVICE)
	url := stream.CreateUrl(streamName)
	s := stream.NewAllTickersService(symbol)

	go s.Call(url, c)
}

func callRollingWindowStatsService(symbol string, c chan interface{}) {
	window := "1h"
	streamName := stream.CreateStreamName(constant.PATTERN_ROLLING_WINDOW_STATS_SERVICE, symbol, window)
	url := stream.CreateUrl(streamName)
	logger.Log.Debug(url)
	s := stream.NewRollingWindowStatsService(symbol)
	go s.Call(url, c)
}

func callAllRollingWindowStatsService(symbol string, c chan interface{}) {
	window := "1h"
	streamName := stream.CreateStreamName(constant.PATTERN_ALL_ROLLING_WINDOW_STATS_SERVICE, window)
	url := stream.CreateUrl(streamName)

	s := stream.NewAllRollingWindowStatsService(symbol)
	go s.Call(url, c)
}

func callBookTickerService(symbol string, c chan interface{}) {
	streamName := stream.CreateStreamName(constant.PATTERNBOOK_TICKER_SERVICE, symbol)
	url := stream.CreateUrl(streamName)
	logger.Log.Debug(url)
	s := stream.NewBookTickerService(symbol)
	go s.Call(url, c)
}

func callAllBookTickerService(symbol string, c chan interface{}) {
	streamName := stream.CreateStreamName(constant.PATTERN_ALL_BOOK_TICKERS_SERVICE)
	url := stream.CreateUrl(streamName)
	logger.Log.Debug(url)
	s := stream.NewAllBookTickerService(symbol)
	go s.Call(url, c)
}

func callPartialBookDepthService(symbol string, c chan interface{}) {
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
