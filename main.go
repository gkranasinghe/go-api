package main

import (
	"github.com/gkranasinghe/go-api/user"
)

func main() {
	config := user.NewConfig()

	db, err := user.ConnectDatabase(config)

	if err != nil {
		panic(err)
	}

	userRepository := user.NewUserRepository(db)

	userService := user.NewUserService(config, userRepository)

	handler := NewHandler(userService)

	server, err := NewServer(handler)
	if err != nil {
		panic(err)
	}
	server.Logger.Fatal(server.Start(":1323"))

}
