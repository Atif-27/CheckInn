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

const (
	dburi  = "mongodb://localhost:27017"
	DBNAME = "hotel-reservation"
)

func main() {
	portPtr := flag.String("port", ":9000", "The PORT to which API server listens")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	//Handlers Initialization
	userHandler := *api.NewUserHandler(database.NewMongoUserStore(client, DBNAME))
	app := fiber.New(config.ErrConfig)
	apiV1 := app.Group("/api/v1")

	apiV1.Get("/users", userHandler.HandleGetUsers)
	apiV1.Get("/users/:id", userHandler.HandleGetUser)
	apiV1.Post("/users", userHandler.HandlePostUser)
	apiV1.Delete("/users/:id", userHandler.HandleDeleteUser)
	apiV1.Put("/users/:id", userHandler.HandlePutUser)
	app.Listen(*portPtr)
}
