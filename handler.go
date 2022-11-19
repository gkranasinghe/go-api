package main

import (
	"net/http"

	"github.com/gkranasinghe/go-api/user"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Users *user.UserService
}

var _ ServerInterface = &Handler{}

func NewHandler(u *user.UserService) *Handler {
	return &Handler{

		Users: u,
	}
}

type ServerInterface interface {
	FindAll(ctx echo.Context) error
}

var _ ServerInterface = &Handler{}

func (h *Handler) FindAll(ctx echo.Context) error {
	_, err := h.Users.FindAll(ctx)
	if err != nil {
		return err
	}
	// TODO: return ctx.JSON(http.StatusOK,NewUsersDto(list))
	return ctx.JSON(http.StatusOK, "HELLO")
}
