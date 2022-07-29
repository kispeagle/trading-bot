package dao

import (
	"context"
	"time"

	db "github.com/kispeagle/trading-bot/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var pairSymbolCollection = db.ConnectCollection(db.Client, "pair_symbol")

func UpdateSymbolList(symbolList []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pairSymbolData := primitive.D{{"list_symbol", symbolList}}
	pairSymbolData = append(pairSymbolData, bson.E{"name", "Pair Symbols"})
	pairSymbolData = append(pairSymbolData, bson.E{"count", len(symbolList)})

	upsert := true
	opts := options.UpdateOptions{
		Upsert: &upsert,
	}
	_, err := pairSymbolCollection.UpdateOne(ctx, bson.M{"name": "Pair Symbols"}, bson.D{{"$set", pairSymbolData}}, &opts)
	if err != nil {
		return err
	}
	return nil
}

func GetLatestSymbolList() interface{} {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var data primitive.D

	pairSymbolCollection.FindOne(ctx, bson.M{}).Decode(&data)
	return data[3].Value
}
