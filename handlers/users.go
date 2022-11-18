package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func AddUser(c echo.Context) error {
	user := User{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		log.Printf("Failed processing addUser request: %s\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("this is your dog: %#v", user)
	return c.String(http.StatusOK, "we got your user!")
}
