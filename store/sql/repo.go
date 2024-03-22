package sql

import (
	"context"
	"idk/store/sql/db"

	"github.com/jackc/pgx/v5"
)

type store struct {
	Conn    *pgx.Conn
	Queries *db.Queries
}

func New(dbUrl string) (*store, error) {
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		return nil, err
	}

	q := db.New(conn)
	return &store{conn, q}, nil
}

func (s *store) Close() {
	s.Conn.Close(context.Background())
}
