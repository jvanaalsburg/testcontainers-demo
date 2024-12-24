---
layout: default
---

# Get All Users

Request handler

````md magic-move
```go {7}
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
func (h Handler) GetAllUsers(c echo.Context) error {
}
```

```go {*|3,8,13}
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
````
