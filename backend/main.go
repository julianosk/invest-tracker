package main

import (
	"context"
	"fmt"
	"investment-management-app/config"
	"investment-management-app/controllers"
	"investment-management-app/routes"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Errorf("cannot load configuration", err))
	}

	clientOptions := options.Client().ApplyURI(cfg.MongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
	mongoCollection := client.Database(cfg.MongoCollection).Collection(cfg.MongoCollection)
	investmentController := controllers.NewInvestmentsController(mongoCollection)

	r := gin.Default()
	routes.SetupRoutes(r, investmentController)

	log.Println("Server is running on port 8080")
	log.Fatal(r.Run(":8080"))
}
