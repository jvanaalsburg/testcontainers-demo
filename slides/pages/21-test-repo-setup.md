---
layout: default
---

# Test Repo Setup

Single-use containers

````md magic-move
```go {*|4|4,12,17|6-10}
// src/repos/repo.ge
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
```
````
