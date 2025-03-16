package api

import (
	"errors"
	"fmt"

	"github.com/Atif-27/hotel-reservation/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelHandler struct {
	store *database.DbStore
}

func NewHotelHandler(store  *database.DbStore) *HotelHandler {
	return &HotelHandler{
		store: store,
	}
}
type ResourceResponse struct {
	Results int `json:"results"`
	Data    any `json:"data"`
	Page    int `json:"page"`
}
type HotelQueryParams struct {
	Rating int
	// database.Pagination
}
func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	var params HotelQueryParams
	if err := c.QueryParser(&params); err != nil {
		// return ErrBadRequest()
	}
	var filter bson.M
	if params.Rating!= 0 {
		filter = bson.M{
			"rating": params.Rating,
		}
	}
	
	hotels, err := h.store.Hotel.GetHotels(c.Context(), filter,
	//  &params.Pagination
	)
	if err != nil {
		return err
	}
	resp := ResourceResponse{
		Results: len(hotels),
		Data:    hotels,
		// Page:    int(params.Page),
	}
	return c.JSON(resp)
}


func (h *HotelHandler) HandleGetHotelById(c *fiber.Ctx) error {
	id := c.Params("id")
	if len(id) == 0 {
		fmt.Printf("supplied empty id %s", id)
		// return ErrInvalidId()
	}

	hotel, err := h.store.Hotel.GetHotelById(c.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(map[string]string{"error": "not found!"})
		}
		// if len(err.(db.DBError).Err) != 0 {
			// return ErrInvalidId()
		// }
		return err
	}
	return c.JSON(hotel)
}



func (h *HotelHandler) HandleGetRooms(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{
		"hotelId": oid,
	}
	rooms, err := h.store.Room.GetRooms(c.Context(), filter)
	if err != nil {
		return err
	}
	return c.JSON(rooms)
}