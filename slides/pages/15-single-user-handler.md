---
layout: default
---

# Get Single User

Request handler

````md magic-move
```go {8}
// src/server.go
func main() {
	e := echo.New()

	handler := initHandler()

	e.GET("/users", handler.GetAllUsers)
	e.GET("/users/:id", handler.GetUser)
	e.POST("/users", handler.CreateUser)
	e.PUT("/users/:id", handler.UpdateUser)
	e.DELETE("/users/:id", handler.DeleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}

```

```go {*}
// src/handlers/users.go
func (h Handler) GetUser(c echo.Context) error {
}
```

```go {*|3,8,13,18|8}
// src/handlers/users.go
func (h Handler) GetUser(c echo.Context) error {
	repo, err := repos.NewRepo(h.connStr)
	if err != nil {
		return err
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return err
	}

	user, err := repo.GetUser(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
```
````
