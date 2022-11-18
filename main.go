package main

import (
	"github.com/labstack/echo/v4"

	"github.com/gkranasinghe/go-api/handlers"
)

func main() {
	e := echo.New()
	e.POST("/users", handlers.AddUser)
	// e.GET("/users/:id", handlers.getUser)
	e.Logger.Fatal(e.Start(":1323"))
}

// func ListAllUsers(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello, There!")
// }

// // e.GET("/users/:id", getUser)
//
//	func getUser(c echo.Context) error {
//		// User ID from path `users/:id`
//		id := c.Param("id")
//		return c.String(http.StatusOK, "id:"+id)
//	}
