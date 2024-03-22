package sql

import (
	"context"
	"idk/store/sql/db"
	"idk/user"

	"github.com/jackc/pgx/v5"
)

type userRepository struct {
	store   *store
	Conn    *pgx.Conn
	Queries *db.Queries
}

func NewUserRepository(s *store) user.UserRepository {
	return &userRepository{
		store:   s,
		Conn:    s.Conn,
		Queries: s.Queries,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, u user.User) (user.User, error) {
	tx, err := r.Conn.Begin(ctx)
	if err != nil {
		return user.User{}, err
	}
	defer tx.Rollback(ctx)

	qtx := r.Queries.WithTx(tx)
	dbUser, err := qtx.CreateUser(ctx, db.CreateUserParams{
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
	})
	if err != nil {
		return user.User{}, err
	}

	dbProfile, err := qtx.CreateProfile(ctx, db.CreateProfileParams{
		UserID:    dbUser.ID,
		FirstName: u.Profile.FirstName,
		LastName:  u.Profile.LastName,
	})
	if err != nil {
		return user.User{}, err
	}

	u.ID = int(dbUser.ID)
	u.Profile.ID = int(dbProfile.ID)
	u.Profile.UserID = int(dbProfile.UserID)
	u.CreatedAt = dbUser.CreatedAt.Time
	u.UpdatedAt = dbUser.UpdatedAt.Time
	u.Profile.CreatedAt = dbProfile.CreatedAt.Time
	u.Profile.UpdatedAt = dbProfile.UpdatedAt.Time

	err = tx.Commit(ctx)
	if err != nil {
		return user.User{}, err
	}

	return u, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, userID int) (user.User, error) {
	dbUser, err := r.Queries.GetUserByID(ctx, int32(userID))
	if err != nil {
		return user.User{}, err
	}

	dbProfile, err := r.Queries.GetProfileByUserID(ctx, dbUser.ID)
	if err != nil {
		return user.User{}, err
	}

	u := user.User{
		ID:       int(dbUser.ID),
		Username: dbUser.Username,
		Password: dbUser.Password,
		Email:    dbUser.Email,
		Profile: user.Profile{
			ID:        int(dbProfile.ID),
			UserID:    int(dbProfile.UserID),
			FirstName: dbProfile.FirstName,
			LastName:  dbProfile.LastName,
			CreatedAt: dbProfile.CreatedAt.Time,
			UpdatedAt: dbProfile.UpdatedAt.Time,
		},
		CreatedAt: dbUser.CreatedAt.Time,
		UpdatedAt: dbUser.UpdatedAt.Time,
	}

	return u, nil
}

func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (user.User, error) {
	dbUser, err := r.Queries.GetUserByUsername(ctx, username)
	if err != nil {
		return user.User{}, err
	}

	dbProfile, err := r.Queries.GetProfileByUserID(ctx, dbUser.ID)
	if err != nil {
		return user.User{}, err
	}

	u := user.User{
		ID:       int(dbUser.ID),
		Username: dbUser.Username,
		Password: dbUser.Password,
		Email:    dbUser.Email,
		Profile: user.Profile{
			ID:        int(dbProfile.ID),
			UserID:    int(dbProfile.UserID),
			FirstName: dbProfile.FirstName,
			LastName:  dbProfile.LastName,
			CreatedAt: dbProfile.CreatedAt.Time,
			UpdatedAt: dbProfile.UpdatedAt.Time,
		},
		CreatedAt: dbUser.CreatedAt.Time,
		UpdatedAt: dbUser.UpdatedAt.Time,
	}

	return u, nil
}
