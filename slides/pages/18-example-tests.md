---
layout: two-cols-header
---

# Example Tests

User repo tests

::left::

**As part of a test suite**

````md magic-move
```go {*|4-5|7}
func (suite *RepoTestSuite) TestRepoGetAllUsers() {
	t := suite.T()

	users, err := suite.repo.GetAllUsers()
	assert.NoError(t, err)

	assert.Equal(t, 3, len(users))
}
```
````

::right::

**Individual test function**

````md magic-move
```go {*|4-11|13-15|*}
func TestRepoCreateUser(t *testing.T) {
	repo := TestRepo(t, []string{"000-init-schema.sql"})

	user := models.User{
		FirstName: "Draco",
		LastName:  "Malfoy",
		Email:     "dmalfoy@hogwarts.edu",
	}

	newUser, err := repo.CreateUser(user)
	assert.NoError(t, err)

	assert.Equal(t, "Draco", newUser.FirstName)
	assert.Equal(t, "Malfoy", newUser.LastName)
	assert.Equal(t, "dmalfoy@hogwarts.edu", newUser.Email)
}
```
````
