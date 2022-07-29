package dao

import (
	"context"
	"log"
	"time"

	"github.com/kispeagle/go-binance/stream"
	db "github.com/kispeagle/trading-bot/database"

	"go.mongodb.org/mongo-driver/bson"
)

var klineCollection = db.ConnectCollection(db.Client, "kline")

type KlineDao struct {
}

func InsertKline(k interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := klineCollection.InsertOne(ctx, k)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAll() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := klineCollection.DeleteMany(ctx, bson.M{})
	if err != nil {
		return err
	}

	return nil
}

func GetInterval(symbol string, start, end time.Time) []stream.Kline {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{"symbol", symbol}, {"time_event", bson.D{{"$gt", start}, {"$lt", end}}}}
	result, err := klineCollection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	var klines []stream.Kline
	err = result.All(ctx, &klines)
	return klines
}

func FindOne()
