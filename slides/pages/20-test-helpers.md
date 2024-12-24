---
layout: default
---

# Test Helpers

Functions for initializing _Testcontainers_

````md magic-move
```go {*|2-5,7,21-27|8-17|19|*}
// src/testhelpers/container.go
type PostgresContainer struct {
	*postgres.PostgresContainer
	ConnectionString string
}

func CreatePostgresContainer(ctx context.Context) *PostgresContainer {
	pgContainer, _ := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		postgres.WithDatabase("test"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2),
		),
	)

	connStr, _ := pgContainer.ConnectionString(ctx, "sslmode=disable")

	return &PostgresContainer{
		PostgresContainer: pgContainer,
		ConnectionString:  connStr,
	}
}
```

```go {11,14|4-9}
// src/testhelpers/container.go
type PostgresContainer struct { ... }

func AllInitScripts() []string {
	return []string{
		"000-init-schema.sql",
		"001-add-users.sql",
	}
}

func CreatePostgresContainer(ctx context.Context, initScripts []string) *PostgresContainer {
	pgContainer, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithInitScripts(initScripts...),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		postgres.WithDatabase("test"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2),
		),
	)

	// ...
}
```
````
