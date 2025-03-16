package database

import (
	"context"

	"github.com/Atif-27/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelStore interface {
	InsertHotel(context.Context, *types.Hotel) (*types.Hotel, error)
	UpdateHotel(ctx context.Context, filter map[string]any, update map[string]any) error
	GetHotelById(context.Context, string) (*types.Hotel, error)
	GetHotels(ctx context.Context, filter map[string]any,
	//  paginaton *Pagination
	) ([]*types.Hotel, error)
}

type MongoHotelStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

// constructor
func NewMongoHotelStore(client *mongo.Client, dbname string) *MongoHotelStore {
	return &MongoHotelStore{
		client: client,
		coll:   client.Database(dbname).Collection(HotelColl),
	}
}
func (m *MongoHotelStore) InsertHotel(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error) {
	res, err := m.coll.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}
	hotel.ID = res.InsertedID.(primitive.ObjectID)
	return hotel, nil
}

func (m *MongoHotelStore) UpdateHotel(ctx context.Context, filter map[string]any, update map[string]any) error {
	_, err := m.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (s *MongoHotelStore) GetHotels(ctx context.Context, filter map[string]any,

// pag *Pagination
) ([]*types.Hotel, error) {
	// var (
	// 	skip = (pag.Page - 1) * pag.Limit
	// )
	// opts := &options.FindOptions{
	// 	Limit: &pag.Limit,
	// 	Skip:  &skip,
	// }

	resp, err := s.coll.Find(ctx, filter)//  opts

	if err != nil {
		return nil, err
	}
	var hotels []*types.Hotel
	if err := resp.All(ctx, &hotels); err != nil {
		return []*types.Hotel{}, nil
	}
	return hotels, nil
}

func (m *MongoHotelStore) GetHotelById(ctx context.Context, id string) (*types.Hotel, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// return nil, NewResourceError(err.Error())
		return nil, err
	}
	var hotel types.Hotel
	if err := m.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&hotel); err != nil {
		return nil, err
	}
	return &hotel, nil
}
