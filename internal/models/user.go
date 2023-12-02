package models

type User struct {
	Id       int       `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Name     string    `json:"name"`
	Phone    string    `json:"phone"`
	Projects []Project `json:"projects"`
}
