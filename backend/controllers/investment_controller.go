package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"investment-management-app/backend/models"
	"investment-management-app/backend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"time"
)

var investmentCollection *mongo.Collection = config.GetCollection(config.DB, "investments")

func CreateInvestment(c *gin.Context) {
	var investment models.Investment
	if err := c.BindJSON(&investment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	investment.ID = primitive.NewObjectID()
	investment.CreatedAt = time.Now()

	_, err := investmentCollection.InsertOne(context.Background(), investment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, investment)
}

func GetInvestments(c *gin.Context) {
	var investments []models.Investment
	cursor, err := investmentCollection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var investment models.Investment
		if err := cursor.Decode(&investment); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		investments = append(investments, investment)
	}
	c.JSON(http.StatusOK, investments)
}

func UpdateInvestment(c *gin.Context) {
	id := c.Param("id")
	var investment models.Investment
	if err := c.BindJSON(&investment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": investment}

	_, err := investmentCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, investment)
}