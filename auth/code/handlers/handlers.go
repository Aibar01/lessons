package handlers

import (
	"auth/code/services"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	UserHandlerInterface
}

func NewHandler(serv *services.Service) *Handler {
	return &Handler{
		UserHandlerInterface: NewUserHandler(serv.UserServiceInterface),
	}
}

func (h *Handler) InitRoutes(app *fiber.App) {
	h.UserHandlerInterface.InitRoutes(app)
}
