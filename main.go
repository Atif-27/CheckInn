package main

import (
	"context"
	"log"
	"os"

	"github.com/Atif-27/hotel-reservation/api"
	"github.com/Atif-27/hotel-reservation/config"
	"github.com/Atif-27/hotel-reservation/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using system environment variables")
    }
	mongoURI := os.Getenv("MONGO_URI");if mongoURI == "" {
        log.Fatal("ENV ERROR: MONGO_URI not set")
    }
	port:= os.Getenv("PORT"); if port==""{
		log.Fatal("ENV ERROR: PORT is not set")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	var (
		//! Store Initialization
		userStore  = database.NewMongoUserStore(client, database.DBNAME)
		hotelStore = database.NewMongoHotelStore(client, database.DBNAME)
		roomStore  = database.NewMongoRoomStore(client, database.DBNAME, hotelStore)
		dbStore    = database.MakeDbStore(userStore, roomStore, hotelStore)

		//! Handler Initialization
		userHandler = *api.NewUserHandler(dbStore)
		hotelHander = *api.NewHotelHandler(dbStore)

		// ! APIS
		app   = fiber.New(config.ErrConfig)
		apiV1 = app.Group("/api/v1")
	)

	//* User Routes
	apiV1.Get("/users", userHandler.HandleGetUsers)
	apiV1.Get("/users/:id", userHandler.HandleGetUser)
	apiV1.Post("/users", userHandler.HandlePostUser)
	apiV1.Delete("/users/:id", userHandler.HandleDeleteUser)
	apiV1.Put("/users/:id", userHandler.HandlePutUser)

	//* Hotel Routes
	apiV1.Get("/hotels", hotelHander.HandleGetHotels)
	apiV1.Get("/hotels/:id", hotelHander.HandleGetHotelById)
	apiV1.Get("/hotels/:id/rooms", hotelHander.HandleGetRooms)

	app.Listen(port)
}
