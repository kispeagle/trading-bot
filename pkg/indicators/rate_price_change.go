package indicator

import (
	"errors"
	"strconv"
	"time"

	logger "github.com/kispeagle/go-binance/logs"
)

func RateOfChangePrice(interval string, cur float64) (float64, error) {
	interval_, err := strconv.Atoi(interval[:len(interval)-1])
	if err != nil {
		logger.Log.Error(err)
		return -1, err
	}
	timestamp := time.Now()
	timestamp.Second()
	var prev float64
	switch interval {
	case "1m":
		prev = getMinClosePrice(time.Now().Add(-time.Duration(interval_) * time.Minute))
	case "5m":
		prev = getMinClosePrice(time.Now().Add(-time.Duration(interval_) * time.Minute))

	case "15m":
		prev = getMinClosePrice(time.Now().Add(-time.Duration(interval_) * time.Minute))
	case "30m":
		prev = getMinClosePrice(time.Now().Add(-time.Duration(interval_) * time.Minute))
	case "1h":
		prev = getMinClosePrice(time.Now().Add(-time.Duration(interval_) * time.Hour))
	case "4h":
		prev = getMinClosePrice(time.Now().Add(-time.Duration(interval_) * time.Hour))
	case "6h":
		prev = getMinClosePrice(time.Now().Add(-time.Duration(interval_) * time.Hour))
	case "12h":
		prev = getMinClosePrice(time.Now().Add(-time.Duration(interval_) * time.Hour))
	case "1d":
		interval_ = 24
		prev = getMinClosePrice(time.Now().Add(-time.Duration(interval_) * time.Hour))
	case "1w":
		interval_ = 24 * 7
		prev = getMinClosePrice(time.Now().Add(-time.Duration(interval_) * time.Hour))
	case "1M":
		interval_ = 24 * 30
		prev = getMinClosePrice(time.Now().Add(-time.Duration(interval_) * time.Hour))
	default:
		return -1, errors.New("Wrong interval")
	}

	prev = getMinClosePrice(timestamp)
	r := (cur/prev - 1) * 100
	return r, nil
}

func getMinClosePrice(timestamp time.Time) float64 {

	return 0
}
