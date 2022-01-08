package pgsql_db

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type User struct {
	UserId uuid.UUID `json:"user_id" query:"user_id"`
	FirstName string `json:"first_name" query:"first_name" validate:"required"`
	LastName string `json:"last_name" query:"last_name" validate:"required"`
	Email string `json:"email" query:"email" validate:"required,email,unique"`
	PhoneNumber int `json:"phone_number" validate:"max=14"`
}

type UserDependency struct {
	DB *sqlx.DB
}

type UserValidator struct {
	Validator *validator.Validate
}

var ErrNotExists = errors.New("User id was not found")
var ErrConnFailed = errors.New("Connection failed!!!")
var ErrQuery = errors.New("Execute query error!!!")
