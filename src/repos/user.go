package repos

import (
	"demo-api/models"

	"github.com/google/uuid"
)

func (r Repo) GetAllUsers() ([]models.User, error) {
	users := []models.User{}

	err := r.conn.Select(&users, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r Repo) GetUser(id uuid.UUID) (models.User, error) {
	user := models.User{}

	err := r.conn.Get(&user, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return user, err
	}

	return user, nil
}
