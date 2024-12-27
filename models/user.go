package models

type User struct {
	ID       int64  `json:"id"`
	Name     string `binding:"required" ,json:"name"`
	Email    string `binding:"required" ,json:"email"`
	Password string `binding:"required" ,json:"password"`
}

func NewUser(name, email, password string) User {
	return User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}
