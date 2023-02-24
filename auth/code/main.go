package main

import (
	"auth/code/handlers"
	"auth/code/repository"
	"auth/code/services"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func main() {
	repo := repository.NewRepository(&pgxpool.Pool{})
	serv := services.NewService(repo)
	hndlrs := handlers.NewHandler(serv)
	app := fiber.New()
	hndlrs.InitRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
