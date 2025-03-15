package main

import (
	"context"
	"flag"
	"log"

	"github.com/Atif-27/hotel-reservation/api"
	"github.com/Atif-27/hotel-reservation/config"
	"github.com/Atif-27/hotel-reservation/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"
const dbName = "hotel"
const userCol="users"

func main() {
	portPtr := flag.String("port", ":9000", "The PORT to which API server listens")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}
	//Handlers Initialization
	userHandler := api.NewUserHandler(database.NewMongoUserStore(client))
	app := fiber.New(config.ErrConfig)
	apiV1 := app.Group("/api/v1")

	apiV1.Get("/users", userHandler.HandleGetUsers)
	apiV1.Get("/users/:id", userHandler.HandleGetUser)
	app.Listen(*portPtr)
}
