package types

type User struct{
	ID string `bson:"_id" json:"id"`
	Firstname string `bson:"firstName" json:"firstName"`
	Lastname string `bson:"lastName" json:"lastName"` 
}