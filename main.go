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

func main() {
	portPtr := flag.String("port", ":9000", "The PORT to which API server listens")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(database.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	var(
		//! Store Initialization
		userStore=database.NewMongoUserStore(client, database.DBNAME)
		hotelStore=database.NewMongoHotelStore(client,database.DBNAME)
		roomStore=database.NewMongoRoomStore(client, database.DBNAME,hotelStore)
		dbStore= database.MakeDbStore(userStore,roomStore,hotelStore)
		
		//! Handler Initialization
		userHandler = *api.NewUserHandler(dbStore)
		hotelHander= *api.NewHotelHandler(dbStore)


		// ! APIS
		app = fiber.New(config.ErrConfig)
		apiV1 = app.Group("/api/v1")
	)


	//* User Routes
	apiV1.Get("/users", userHandler.HandleGetUsers)
	apiV1.Get("/users/:id", userHandler.HandleGetUser)
	apiV1.Post("/users", userHandler.HandlePostUser)
	apiV1.Delete("/users/:id", userHandler.HandleDeleteUser)
	apiV1.Put("/users/:id", userHandler.HandlePutUser)

	//* Hotel Routes
	apiV1.Get("/hotels",hotelHander.HandleGetHotels)
	apiV1.Get("/hotels/:id",hotelHander.HandleGetHotelById)
	apiV1.Get("/hotels/:id/rooms",hotelHander.HandleGetRooms)


	app.Listen(*portPtr)
}
