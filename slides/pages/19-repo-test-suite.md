---
layout: default
---

# Repo Test Suite

Share containers across multiple tests

````md magic-move
```go {*|2-7,9,21|5-6,11,13,18|22}
// src/repos/user_test.go
type RepoTestSuite struct {
	suite.Suite
	ctx       context.Context
	container *testhelpers.PostgresContainer
	repo      *Repo
}

func (suite *RepoTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	suite.container = testhelpers.CreatePostgresContainer(suite.ctx, testhelpers.AllInitScripts())

	repo, err := NewRepo(suite.container.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	suite.repo = repo
}

func (suite *RepoTestSuite) TearDownSuite() {
	if err := suite.container.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}
```
````
