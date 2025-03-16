package database

import (
	"context"

	"github.com/Atif-27/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomStore interface {
	InsertRoom(context.Context, *types.Room) (*types.Room, error)
	GetRooms(ctx context.Context, filter bson.M) ([]*types.Room, error)
}

type MongoDbRoomStore struct {
	client     *mongo.Client
	roomColl   *mongo.Collection
	hotelStore HotelStore
}

func NewMongoRoomStore(client *mongo.Client, dbname string, hotelStore HotelStore) *MongoDbRoomStore {
	return &MongoDbRoomStore{
		client:     client,
		roomColl:   client.Database(dbname).Collection(RoomColl),
		hotelStore: hotelStore,
	}
}

func (m *MongoDbRoomStore) InsertRoom(ctx context.Context, room *types.Room) (*types.Room, error) {
	res, err := m.roomColl.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}
	room.ID = res.InsertedID.(primitive.ObjectID)

	// update the hotel with this room
	filter := bson.M{"_id": room.HotelID}
	update := bson.M{"$push": bson.M{"rooms": room.ID}}
	if err := m.hotelStore.UpdateHotel(ctx, filter, update); err != nil {
		return nil, err
	}
	return room, nil
}

func (m *MongoDbRoomStore) GetRooms(ctx context.Context, filter bson.M) ([]*types.Room, error) {
	resp, err := m.roomColl.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var rooms []*types.Room
	if err := resp.All(ctx, &rooms); err != nil {
		return nil, err
	}
	return rooms, nil
}
