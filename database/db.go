package database

const (
	DBURI     = "mongodb://localhost:27017"
	DBNAME    = "hotel-reservation"
	UserColl  = "users"
	HotelColl = "hotels"
	RoomColl  = "rooms"
)

type DbStore struct {
	User  UserStore
	Room  RoomStore
	Hotel HotelStore
}

func MakeDbStore(user UserStore, room RoomStore, hotel HotelStore) *DbStore {
	return &DbStore{
		User:  user,
		Hotel: hotel,
		Room:  room,
	}
}
