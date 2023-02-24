package handlers

import (
	"auth/code/services"
	"github.com/gofiber/fiber/v2"
)

type UserHandlerInterface interface {
	CreateUser(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	InitRoutes(router fiber.Router)
}

type UserHandler struct {
	serv services.UserServiceInterface
}

func NewUserHandler(serv services.UserServiceInterface) *UserHandler {
	return &UserHandler{serv: serv}
}

func (u *UserHandler) CreateUser(c *fiber.Ctx) error {
	return c.SendString("hello")
}

func (u *UserHandler) GetUser(c *fiber.Ctx) error {
	return c.SendString("hello")
}

func (u *UserHandler) InitRoutes(router fiber.Router) {
	router.Post("/users", u.CreateUser)
	router.Get("/users", u.GetUser)
}
