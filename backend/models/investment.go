package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Investment struct {
    ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Name       string             `bson:"name" json:"name"`
    Amount     float64            `bson:"amount" json:"amount"`
    AssetClass string             `bson:"asset_class" json:"asset_class"`
}