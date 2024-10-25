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

func (r Repo) CreateUser(user models.User) (models.User, error) {
	query := `
    INSERT INTO users (first_name, last_name, email)
      VALUES ($1, $2, $3)
      RETURNING id`

	var userId uuid.UUID

	err := r.conn.QueryRow(query, user.FirstName, user.LastName, user.Email).Scan(&userId)
	if err != nil {
		return user, err
	}

	user.Id = userId
	return user, nil
}

func (r Repo) DeleteUser(id uuid.UUID) error {
	_, err := r.conn.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return err
	}

	return nil
}
