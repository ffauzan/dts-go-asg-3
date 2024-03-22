package user

import (
	"context"
	"idk/util"
)

type service struct {
	repo UserRepository
}

func NewService(repo UserRepository) UserService {
	return &service{repo}
}

func (s *service) CreateUser(ctx context.Context, u User) (User, error) {
	hashedPassword, err := util.HashPassword(u.Password)
	if err != nil {
		return User{}, err
	}

	u.Password = hashedPassword

	user, err := s.repo.CreateUser(ctx, u)
	if err != nil {
		return User{}, err
	}

	user.Password = ""
	return user, nil
}

func (s *service) GetUserByID(ctx context.Context, userID int) (User, error) {
	return s.repo.GetUserByID(ctx, userID)
}

func (s *service) GetUserByUsername(ctx context.Context, username string) (User, error) {
	return s.repo.GetUserByUsername(ctx, username)
}
