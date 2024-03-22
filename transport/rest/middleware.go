package rest

import (
	"idk/user"
	"idk/util"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(userService user.UserService, jwtSecret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get the token from the request, split the token from the "Bearer" prefix
			headerValue := c.Request().Header.Get("Authorization")
			if headerValue == "" || len(headerValue) < len("Bearer ")+1 {
				return c.JSON(401, BaseResponse{
					Status:  401,
					Message: "Unauthorized",
					Data:    "No token provided",
				})
			}

			token := headerValue[len("Bearer "):]

			if token == "" {
				return c.JSON(401, BaseResponse{
					Status:  401,
					Message: "Unauthorized",
					Data:    "No token provided",
				})
			}

			// Validate the token
			claims, err := util.ValidateToken(token, jwtSecret)
			if err != nil {
				return c.JSON(401, BaseResponse{
					Status:  401,
					Message: "Unauthorized",
					Data:    err.Error(),
				})
			}

			// Get the user from the database
			u, err := userService.GetUserByID(c.Request().Context(), claims)
			if err != nil {
				return c.JSON(401, BaseResponse{
					Status:  401,
					Message: "Unauthorized",
					Data:    err.Error(),
				})
			}

			// Set the user in the context
			c.Set("user", u)

			return next(c)
		}
	}
}
