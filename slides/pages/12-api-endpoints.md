---
layout: default
---

# Demo API

Basic CRUD endpoints for user records

```go {*|2,5,15|7|8-13|*}
// src/server.go
import "github.com/labstack/echo/v4"

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
