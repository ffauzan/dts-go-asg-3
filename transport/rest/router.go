package rest

import (
	"idk/config"
	"idk/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type BaseResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewRouter(userService user.UserService, c *config.AppConfig) *echo.Echo {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "ip=${remote_ip}, method=${method}, uri=${uri}, status=${status}, latency=${latency}\n",
	}))

	e.Validator = NewValidator()

	userHandler := NewUserHandler(userService, c.JWTSecret)

	e.POST("/register", userHandler.Register)
	e.POST("/login", userHandler.Login)

	e.GET("/me", userHandler.Me, AuthMiddleware(userService, c.JWTSecret))

	return e
}
