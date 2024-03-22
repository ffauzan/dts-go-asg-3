package user

import (
	"context"
	"time"
)

type User struct {
	ID        int
	Username  string
	Password  string
	Email     string
	Profile   Profile
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Profile struct {
	ID        int
	UserID    int
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserService interface {
	CreateUser(ctx context.Context, user User) (User, error)
	GetUserByID(ctx context.Context, userID int) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, user User) (User, error)
	GetUserByID(ctx context.Context, userID int) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
}
