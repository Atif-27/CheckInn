package types

type User struct{
	ID string `bson:"_id,omitempty" json:"id"`
	Firstname string `bson:"firstName" json:"firstName"`
	Lastname string `bson:"lastName" json:"lastName"` 
}