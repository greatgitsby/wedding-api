package api

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Context struct {
	DBPool *pgxpool.Pool
}

func GetDBConn(db_url string) (*pgxpool.Pool, error) {
	return pgxpool.Connect(context.Background(), db_url)
}
