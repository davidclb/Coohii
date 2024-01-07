package sqlc

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// SQLStore provides all functions to execute SQL queries and transactions
type Repo struct {
	connPool *pgxpool.Pool
	*Queries
}

// NewStore creates a new store
func NewRepo(connPool *pgxpool.Pool) *Repo {
	return &Repo{
		connPool: connPool,
		Queries:  New(connPool),
	}
}