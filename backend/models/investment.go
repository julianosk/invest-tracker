package models

import (
	"time"
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
	ID            uint      `gorm:"primaryKey" json:"id"`
	Name          string    `json:"name"`
	Amount        float64   `json:"amount"`
	AssetClass    string    `json:"asset_class"`
	AssetSubClass string    `json:"asset_sub_class"`
	CreatedAt     time.Time `json:"created"`
}
