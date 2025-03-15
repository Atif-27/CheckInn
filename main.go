package main

import (
	"flag"

	"github.com/Atif-27/hotel-reservation/api"
	"github.com/gofiber/fiber/v2"
)

func main(){
	portPtr:= flag.String("port",":9000","The PORT to which API server listens")
	flag.Parse()
	app:= fiber.New()
	apiV1:= app.Group("/api/v1")
	apiV1.Get("/users", api.HandleGetUsers)
	apiV1.Get("/users/:id", api.HandleGetUser)
	app.Listen(*portPtr)
}

