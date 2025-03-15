package api

import (
	"github.com/Atif-27/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u:= types.User{
		ID: "123",
		Firstname: "Atif",
		Lastname: "Ali",
	}
	return c.JSON(u)
}


func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("sdasd")
}