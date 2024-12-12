package repos

import (
	"context"
	"log"
	"testing"

	"demo-api/models"
	"demo-api/testhelpers"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserRepoTestSuite struct {
	suite.Suite
	ctx       context.Context
	container *testhelpers.PostgresContainer
	repo      *Repo
}

func (suite *UserRepoTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	suite.container = testhelpers.CreatePostgresContainer(suite.ctx, testhelpers.AllInitScripts())

	repo, err := NewRepo(suite.container.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	suite.repo = repo
}

func (suite *UserRepoTestSuite) TearDownSuite() {
	if err := suite.container.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func TestUserRepoTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}

// Calls Repo.GetAllUsers, and verifies that the returned list of user records
// is correct.
func (suite *UserRepoTestSuite) TestUserRepoGetAllUsers() {
	t := suite.T()

	users, err := suite.repo.GetAllUsers()

	assert.NoError(t, err)
	assert.Equal(t, 3, len(users))

	expectedUsers := []models.User{
		{
			Id:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			FirstName: "Harry",
			LastName:  "Potter",
			Email:     "hpotter@hogwarts.edu",
		},
		{
			Id:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			FirstName: "Ron",
			LastName:  "Weasley",
			Email:     "rweasley@hogwarts.edu",
		},
		{
			Id:        uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			FirstName: "Hermione",
			LastName:  "Granger",
			Email:     "hgranger@hogwarts.edu",
		},
	}

	assert.Equal(t, expectedUsers, users)
}

// Calls Repo.GetUser with multiple, valid user IDs, and verifies the returned
// user records are correct.
func (suite *UserRepoTestSuite) TestUserRepoGetUser() {
	t := suite.T()

	testCases := map[string]models.User{
		"00000000-0000-0000-0000-000000000001": {
			Id:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			FirstName: "Harry",
			LastName:  "Potter",
			Email:     "hpotter@hogwarts.edu",
		},
		"00000000-0000-0000-0000-000000000002": {
			Id:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			FirstName: "Ron",
			LastName:  "Weasley",
			Email:     "rweasley@hogwarts.edu",
		},
		"00000000-0000-0000-0000-000000000003": {
			Id:        uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			FirstName: "Hermione",
			LastName:  "Granger",
			Email:     "hgranger@hogwarts.edu",
		},
	}

	for userId, expected := range testCases {
		user, err := suite.repo.GetUser(uuid.MustParse(userId))
		assert.NoError(t, err)
		assert.Equal(t, expected, user)
	}
}

// Calls Repo.CreateUser with valid user data, and verifies that the user
// record is created successfully.
func TestUserRepoCreateUser(t *testing.T) {
  t.Parallel()

	repo := TestRepo(t, []string{"000-init-schema.sql"})

	user := models.User{
		FirstName: "Draco",
		LastName:  "Malfoy",
		Email:     "dmalfoy@hogwarts.edu",
	}

	newUser, err := repo.CreateUser(user)
	assert.NoError(t, err)

	_, err = uuid.Parse(newUser.Id.String())
	assert.NoError(t, err)

	assert.Equal(t, "Draco", newUser.FirstName)
	assert.Equal(t, "Malfoy", newUser.LastName)
	assert.Equal(t, "dmalfoy@hogwarts.edu", newUser.Email)

	users, _ := repo.GetAllUsers()
	assert.Equal(t, 1, len(users))
}

// Calls Repo.UpdateUser with valid data, and verifies that the user record is
// updated successfully.
func TestUserRepoUpdateUser(t *testing.T) {
  t.Parallel()

	repo := TestRepo(t, testhelpers.AllInitScripts())

	id := uuid.MustParse("00000000-0000-0000-0000-000000000001")

	user, _ := repo.GetUser(id)

	user.FirstName = "Draco"
	user.LastName = "Malfoy"
	user.Email = "dmalfoy@hogwarts.edu"

	updatedUser, err := repo.UpdateUser(user)
	assert.NoError(t, err)

	assert.Equal(t, "Draco", updatedUser.FirstName)
	assert.Equal(t, "Malfoy", updatedUser.LastName)
	assert.Equal(t, "dmalfoy@hogwarts.edu", updatedUser.Email)

	user, _ = repo.GetUser(id)
	assert.Equal(t, "Draco", user.FirstName)
	assert.Equal(t, "Malfoy", user.LastName)
	assert.Equal(t, "dmalfoy@hogwarts.edu", user.Email)
}

// Calls Repo.DeleteUser with multiple, valid user IDs, and verifies that the
// user records are deleted successfully.
func TestUserRepoDeleteUser(t *testing.T) {
  t.Parallel()

	repo := TestRepo(t, testhelpers.AllInitScripts())

	users, _ := repo.GetAllUsers()
	assert.Equal(t, 3, len(users))

	ids := []string{
		"00000000-0000-0000-0000-000000000001",
		"00000000-0000-0000-0000-000000000002",
		"00000000-0000-0000-0000-000000000003",
	}

	remaining := len(users)

	for _, id := range ids {
		userId := uuid.MustParse(id)

		err := repo.DeleteUser(userId)
		assert.NoError(t, err)

		remaining--

		users, _ := repo.GetAllUsers()
		assert.Equal(t, remaining, len(users))

		_, err = repo.GetUser(userId)
		assert.Error(t, err)
	}
}
