---
layout: default
---

# User Model

Custom Go type

````md magic-move
```go
// src/models/user.go
type User struct {
	Id        uuid.UUID
	FirstName string
	LastName  string
	Email     string
}
```

```go
// src/models/user.go
type User struct {
	Id        uuid.UUID `json:"id" db:"id"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Email     string    `json:"email" db:"email"`
}
```
````
