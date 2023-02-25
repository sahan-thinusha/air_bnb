package controller

import (
	"air_bnb/pkg/entity"
	"air_bnb/pkg/env"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/mongo"
)

func GetSingleAirBNB() (*entity.AirBnb, error) {
	review := entity.AirBnb{}
	db := env.MongoDBConnection
	coll := db.Collection("listingsAndReviews").FindOne(context.Background(), bson.M{})
	coll.Decode(&review)
	return &review, nil
}
