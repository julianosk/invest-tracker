package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "investment-management-app/backend/config"
    "investment-management-app/backend/routes"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    cfg := config.LoadConfig()

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

    r := mux.NewRouter()
    routes.SetupRoutes(r)

    http.Handle("/", r)
    log.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}