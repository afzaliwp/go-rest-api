package services

import (
	"fmt"
	"github.com/afzaliwp/go-rest-api/db"
	"github.com/afzaliwp/go-rest-api/models"
)

func GetUsers() ([]models.User, error) {
	query, err := db.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	defer query.Close()

	var users []models.User

	for query.Next() {
		var user models.User
		err = query.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func InsertUser(user *models.User) (err error) {
	query := `INSERT INTO users(name, email, password) VALUES(?, ?, ?)`

	statement, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("Error while preparing statement: %v", err)
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("Error while executing insert statement: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("Error getting last insert ID: %v", err)
	}

	user.ID = id
	return nil
}

func GetUserById(userId int64) (user *models.User, err error) {
	query := `SELECT * FROM users WHERE id=?`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("Error while preparing get statement: %v", err)
	}

	defer statement.Close()

	row := statement.QueryRow(userId)
	user = &models.User{}
	err = row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func DeleteUserById(userId int64) error {
	_, err := GetUserById(userId)
	if err != nil {
		return err
	}

	query := `DELETE FROM users WHERE id=?`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("Error while preparing delete statement: %v", err)
	}

	defer statement.Close()

	_, err = statement.Exec(userId)
	if err != nil {
		return fmt.Errorf("Error while executing delete statement: %v", err)
	}

	return nil
}

func UpdateUserById(user *models.User) error {
	query := `UPDATE users
			SET name=?, email=?, password=?
			WHERE id=?`

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return fmt.Errorf("Error while preparing update statement: %v", err)
	}

	defer statement.Close()

	_, err = statement.Exec(
		user.Name,
		user.Email,
		user.Password,
		user.ID,
	)
	if err != nil {
		return fmt.Errorf("Error while executing update statement: %v", err)
	}

	return nil
}
