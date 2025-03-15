package config

import "github.com/gofiber/fiber/v2"

var ErrConfig =fiber.Config{
    // Override default error handler
    ErrorHandler: func(ctx *fiber.Ctx, err error) error {
       return ctx.JSON(map[string]string{
        "error": err.Error(),
	   })
    },
}