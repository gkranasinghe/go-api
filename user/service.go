package user

import (
	"github.com/labstack/echo/v4"
)

type UserService struct {
	config     *Config
	repository *UserRepository
}

func (service *UserService) FindAll(ctx echo.Context) ([]User, error) {
	if service.config.Enabled {
		return service.repository.All()
	}

	return []User{}, nil
}

func NewUserService(config *Config, repository *UserRepository) *UserService {
	return &UserService{config: config, repository: repository}
}

type Config struct {
	Enabled      bool
	DatabasePath string
	Port         string
}

func NewConfig() *Config {
	return &Config{
		Enabled:      true,
		DatabasePath: "./example.db",
		Port:         "8000",
	}
}
