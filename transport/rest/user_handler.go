package rest

import (
	"idk/user"
	"idk/util"
	"log"

	"github.com/labstack/echo/v4"
)

type registerRequest struct {
	Username  string `json:"username" validate:"required,min=4,max=32"`
	Password  string `json:"password" validate:"required,min=8,max=64"`
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

type loginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type userHandler struct {
	userService user.UserService
	jwtSecret   string
}

func NewUserHandler(userService user.UserService, jwtSecret string) *userHandler {
	return &userHandler{
		userService: userService,
		jwtSecret:   jwtSecret,
	}
}

func (h *userHandler) Register(c echo.Context) error {
	req := new(registerRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, BaseResponse{
			Status:  400,
			Message: "Invalid request",
			Data:    err.Error(),
		})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(400, BaseResponse{
			Status:  400,
			Message: "Invalid request",
			Data:    err.Error(),
		})
	}

	u := user.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Profile: user.Profile{
			FirstName: req.FirstName,
			LastName:  req.LastName,
		},
	}

	createdUser, err := h.userService.CreateUser(c.Request().Context(), u)
	if err != nil {
		return c.JSON(500, BaseResponse{
			Status:  500,
			Message: err.Error(),
		})
	}

	return c.JSON(201, BaseResponse{
		Status:  201,
		Message: "User created",
		Data:    createdUser,
	})
}

func (h *userHandler) Login(c echo.Context) error {
	req := new(loginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, BaseResponse{
			Status:  400,
			Message: "Invalid request",
			Data:    err.Error(),
		})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(400, BaseResponse{
			Status:  400,
			Message: "Invalid request",
			Data:    err.Error(),
		})
	}

	u, err := h.userService.GetUserByUsername(c.Request().Context(), req.Username)
	if err != nil {
		return c.JSON(500, BaseResponse{
			Status:  500,
			Message: err.Error(),
		})
	}

	log.Println(u)

	// Check password
	if err := util.VerifyPassword(req.Password, u.Password); err != nil {
		return c.JSON(400, BaseResponse{
			Status:  400,
			Message: "Invalid password",
		})
	}

	// Generate JWT token
	token, err := util.GenerateToken(u.ID, 24, h.jwtSecret)
	if err != nil {
		return c.JSON(500, BaseResponse{
			Status:  500,
			Message: err.Error(),
		})
	}

	returnBody := map[string]interface{}{
		"token": token,
	}

	return c.JSON(200, BaseResponse{
		Status:  200,
		Message: "Login success",
		Data:    returnBody,
	})
}

func (h *userHandler) Me(c echo.Context) error {
	// Get the user from the context
	u := c.Get("user").(user.User)
	u.Password = ""

	return c.JSON(200, BaseResponse{
		Status:  200,
		Message: "Success",
		Data:    u,
	})
}
