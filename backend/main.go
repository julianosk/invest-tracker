package main

import (
	"fmt"
	"invest-tracker/config"
	"invest-tracker/controllers"
	"invest-tracker/models"
	"invest-tracker/routes"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Errorf("cannot load configuration: %w", err))
	}

	log.Println("Connecting to PostgreSQL with DSN:", cfg.PostgresDSN)

	db, err := gorm.Open(postgres.Open(cfg.PostgresDSN), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	if err := db.AutoMigrate(&models.Investment{}); err != nil {
		log.Fatal("failed to migrate database:", err)
	}
	log.Println("Connected to PostgreSQL!")

	investmentController := controllers.NewInvestmentsController(db)
	r := gin.Default()
	routes.SetupRoutes(r, investmentController)

	log.Println("Server is running on port", cfg.Port)
	log.Fatal(r.Run(":" + cfg.Port))
}
