package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	logger "github.com/kispeagle/go-binance/logs"
	"github.com/kispeagle/go-binance/stream"
	"github.com/kispeagle/trading-bot/services"

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
	go services.CallKlineService(symbol, c)
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


