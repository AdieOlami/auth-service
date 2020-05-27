package model

import uuid "github.com/satori/go.uuid"

type ResponseData struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data    User   `json:"data"`
	Error   bool   `json:"error"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
}

// type User struct {
// 	ID        uuid.UUID `json:"id"`
// 	FirstName string    `json:"firstName"`
// 	LastName  string    `json:"lastName"`
// 	Email     string    `json:"email"`
// }

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
