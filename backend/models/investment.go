package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AssetClass struct {
    Name        string `bson:"name" json:"name"`
    Description string `bson:"description" json:"description"`
}

type AssetSubClass struct {
    Name        string `bson:"name" json:"name"`
    Description string `bson:"description" json:"description"`
    AssetClass  string `bson:"asset_class" json:"asset_class"`
}

// Update Investment struct to include AssetSubClass
type Investment struct {
    ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Name          string             `bson:"name" json:"name"`
    Amount        float64            `bson:"amount" json:"amount"`
    AssetClass    string             `bson:"asset_class" json:"asset_class"`
    AssetSubClass string             `bson:"asset_sub_class" json:"asset_sub_class"`
    CreatedAt     time.Time          `bson:"created_at" json:"created"`
}