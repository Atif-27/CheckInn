package api

import (
	"context"

	"github.com/Atif-27/hotel-reservation/database"
	"github.com/Atif-27/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)


type UserHandler struct {
	userStore database.UserStore
}
func NewUserHandler(userStore database.UserStore) *UserHandler{
	return &UserHandler{
		userStore: userStore,
	}
}
func(h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var(
		user_id= c.Params("id")
		ctx= context.Background()
	)
	user,err:= h.userStore.GetUserByID(ctx,user_id)
	if err!=nil{
		return err
	}
	return c.JSON(user)
}
func(h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	u:= types.User{
		ID: "123",
		Firstname: "Atif",
		Lastname: "Ali",
	}
	return c.JSON(u)
}


