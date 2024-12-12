package repos

import (
	"context"
	"log"
	"testing"

	"demo-api/testhelpers"

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

func TestRepo(t *testing.T, dbInitScripts []string) *Repo {
	ctx := context.Background()
	container := testhelpers.CreatePostgresContainer(ctx, dbInitScripts)

	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})

	repo, err := NewRepo(container.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	return repo
}
