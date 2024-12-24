---
layout: default
---

# User Repo

Retrieving records from database

````md magic-move
```go {2,8}
// src/handlers/users.go
func (h Handler) GetAllUsers(c echo.Context) error {
	repo, err := repos.NewRepo(h.connStr)
	if err != nil {
		return err
	}

	users, err := repo.GetAllUsers()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}
```

```go {*}
// src/repos/users.go
func (r Repo) GetAllUsers() ([]models.User, error) {
}
```

```go {*|3|5-8|10}
// src/repos/users.go
func (r Repo) GetAllUsers() ([]models.User, error) {
	users := []models.User{}

	err := r.conn.Select(&users, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	return users, nil
}
```
````
