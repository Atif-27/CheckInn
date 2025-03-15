package database

import (
	"context"

	"github.com/Atif-27/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const UserColl="users"
type UserStore interface{
	GetUserByID(context.Context,string) (*types.User,error)
}

type MongoUserStore struct{
	client *mongo.Client
	coll 	*mongo.Collection
}

//constructor
func NewMongoUserStore(client *mongo.Client) *MongoUserStore{
	return &MongoUserStore{
		client: client,
		coll: client.Database(DBNAME).Collection(UserColl),
	}
}

func(m *MongoUserStore) GetUserByID(ctx context.Context,id string) (*types.User, error){
	oid, err:= primitive.ObjectIDFromHex(id); if err!=nil{
		return nil, err
	}
	var user types.User
	err=m.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); if err!=nil{
		return nil, err
	}
	return &user,nil
}