package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Atif-27/hotel-reservation/database"
	"github.com/Atif-27/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	userStore  database.UserStore
	hotelStore database.HotelStore
	roomStore  database.RoomStore
	ctx        = context.Background()
)

func seedHotelAndRoom(name string, location string,rating int) {
	hotel := types.Hotel{
		Name:     name,
		Location: location,
		Rating: rating,
		Rooms:    []primitive.ObjectID{},}
	room := types.Room{
		Type:      types.SINGLE,
		BasePrice: 500.50,
	}
	fmt.Println(room)
	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	room.HotelID = insertedHotel.ID
	_, err = roomStore.InsertRoom(ctx, &room)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(database.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	hotelStore = database.NewMongoHotelStore(client, database.DBNAME)
	roomStore = database.NewMongoRoomStore(client, database.DBNAME, hotelStore)
}

func main() {
	seedHotelAndRoom("Taj Mumbai", "Mumbai",5)
	seedHotelAndRoom("London Halls", "London",2)
	seedHotelAndRoom("London a2", "London",3)
	seedHotelAndRoom("London a2", "London",4)
}
