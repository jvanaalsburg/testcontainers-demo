package handlers

import (
	"context"
	"testing"

	"demo-api/testhelpers"
)

type Handler struct {
	connStr string
}

func NewHandler(connStr string) *Handler {
	return &Handler{connStr: connStr}
}

func TestHandler(t *testing.T, dbInitScripts []string) *Handler {
	ctx := context.Background()
	container := testhelpers.CreatePostgresContainer(ctx, dbInitScripts)

	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})

	return NewHandler(container.ConnectionString)
}
