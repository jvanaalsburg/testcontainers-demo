package repos

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repo struct {
	conn *sqlx.DB
}

func NewRepo(connStr string) (*Repo, error) {
	conn, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &Repo{conn: conn}, nil
}
