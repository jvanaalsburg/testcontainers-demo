package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"demo-api/models"
	"demo-api/testhelpers"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type UserHandlerTestSuite struct {
	suite.Suite
	ctx       context.Context
	container *testhelpers.PostgresContainer
	handler   *Handler
}

func (suite *UserHandlerTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	suite.container = testhelpers.CreatePostgresContainer(suite.ctx, testhelpers.AllInitScripts())

	suite.handler = NewHandler(suite.container.ConnectionString)
}

func (suite *UserHandlerTestSuite) TeardownSuite() {
	if err := suite.container.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func TestUserHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}

// Makes a GET request to Handler.GetAllUsers, and verifies that the returned
// list of users is correct.
func (suite *UserHandlerTestSuite) TestUserHandlerGetAllUsers() {
	t := suite.T()
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users")

	require.NoError(t, suite.handler.GetAllUsers(c))
	assert.Equal(t, http.StatusOK, rec.Code)

	var response []models.User
	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Equal(t, 3, len(response))

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

	assert.Equal(t, expectedUsers, response)
}

// Makes multiple GET requests to Handler.GetUser, and verifies that the
// return users are correct.
func (suite *UserHandlerTestSuite) TestUserHandlerGetUser() {
	t := suite.T()
	e := echo.New()

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
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues(userId)

		require.NoError(t, suite.handler.GetUser(c))
		assert.Equal(t, http.StatusOK, rec.Code)

		var response models.User
		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, expected, response)
	}
}

// Makes a POST request to Handler.CreateUser with valid data, and verifies that
// the user is created successfully.
func TestUserHandlerCreateUser(t *testing.T) {
  t.Parallel()

	e := echo.New()
	handler := TestHandler(t, []string{"000-init-schema.sql"})

	payload := `{
    "first_name": "Draco",
    "last_name": "Malfoy",
    "email": "dmalfoy@hogwarts.edu"
  }`

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users")

	assert.NoError(t, handler.CreateUser(c))
	assert.Equal(t, http.StatusCreated, rec.Code)

	var response models.User
	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Equal(t, "Draco", response.FirstName)
	assert.Equal(t, "Malfoy", response.LastName)
	assert.Equal(t, "dmalfoy@hogwarts.edu", response.Email)
}

// Makes a PUT request to Handler.UpdateUser with valid data, and verifies that
// the user is updated successfully.
func TestUserHandlerUpdateUser(t *testing.T) {
  t.Parallel()

	e := echo.New()
	handler := TestHandler(t, testhelpers.AllInitScripts())

	payload := `{
    "first_name": "Draco",
    "last_name": "Malfoy",
    "email": "dmalfoy@hogwarts.edu"
  }`

	userId := "00000000-0000-0000-0000-000000000001"

	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues(userId)

	assert.NoError(t, handler.UpdateUser(c))
	assert.Equal(t, http.StatusOK, rec.Code)

	var response models.User
	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Equal(t, models.User{
		Id:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
		FirstName: "Draco",
		LastName:  "Malfoy",
		Email:     "dmalfoy@hogwarts.edu",
	}, response)
}

// Makes multiple DELETE requests to Handler.DeleteUser, and verifies that the
// users are successfully deleted.
func TestUserHandlerDeleteUser(t *testing.T) {
  t.Parallel()

	e := echo.New()
	handler := TestHandler(t, testhelpers.AllInitScripts())

	userIds := []string{
		"00000000-0000-0000-0000-000000000001",
		"00000000-0000-0000-0000-000000000002",
		"00000000-0000-0000-0000-000000000003",
	}

	remaining := 3

	for _, userId := range userIds {
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues(userId)

		assert.NoError(t, handler.DeleteUser(c))
		assert.Equal(t, http.StatusNoContent, rec.Code)

		remaining--

		req = httptest.NewRequest(http.MethodGet, "/", nil)
		rec = httptest.NewRecorder()
		c = e.NewContext(req, rec)
		c.SetPath("/users")

		assert.NoError(t, handler.GetAllUsers(c))
		assert.Equal(t, http.StatusOK, rec.Code)

		var response []models.User
		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, remaining, len(response))
	}
}
