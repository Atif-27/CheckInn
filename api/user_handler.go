package api

import (
	"errors"

	"github.com/Atif-27/hotel-reservation/database"
	"github.com/Atif-27/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	store *database.DbStore
}

func NewUserHandler(store *database.DbStore) *UserHandler {
	return &UserHandler{
		store: store,
	}
}
func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var (
		user_id = c.Params("id")
	)
	user, err := h.store.User.GetUserByID(c.Context(), user_id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(map[string]string{"error": "No user found"})
		}
		return err
	}
	return c.JSON(user)
}
func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.store.User.GetUsers(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(users)
}
func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var params types.CreateUserParams
	err := c.BodyParser(&params)
	if err != nil {
		return err
	}
	errors := params.Validate()
	if len(errors) != 0 {
		return c.JSON(map[string]any{
			"message": "Error while creating user",
			"errors":  errors,
		})
	}
	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}

	insertedUser, err := h.store.User.CreateUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(map[string]any{
		"message": "User has been created successfully",
		"data":    insertedUser,
	})
}
func (h *UserHandler) HandlePutUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var data types.UpdateUserParams
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	err = h.store.User.PutUser(c.Context(), id, data)
	if err != nil {
		return err
	}
	return c.JSON(map[string]string{
		"message": "User updated successfully",
	})
}
func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.store.User.DeleteUser(c.Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(map[string]string{
		"message": "User deleted successfully",
	})
}
