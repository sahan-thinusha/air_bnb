package controller

import (
	"air_bnb/pkg/entity"
	"air_bnb/pkg/env"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllAirBNB(index, limit int) ([]*entity.AirBnb, error) {
	var reviews []*entity.AirBnb
	reviews = []*entity.AirBnb{}
	ctx := context.Background()
	db := env.MongoDBConnection

	if index >= 0 && limit >= 0 {
		offset := index * limit
		opts := options.Find()
		opts = opts.SetLimit(int64(limit))
		opts = opts.SetSkip(int64(offset))
		cursor, err := db.Collection("listingsAndReviews").Find(context.Background(), bson.M{}, opts)
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)
		if err = cursor.All(context.Background(), &reviews); err != nil {
			return nil, err
		}

		return reviews, nil
	} else {
		cursor, err := db.Collection("listingsAndReviews").Find(context.Background(), bson.M{})
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)
		if err = cursor.All(context.Background(), &reviews); err != nil {
			return nil, err
		}

		return reviews, nil
	}
}
