package pgsql_db

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          uuid.UUID `json:"id" query:"id"`
	FirstName   string    `json:"first_name" query:"first_name" validate:"required"`
	LastName    string    `json:"last_name" query:"last_name" validate:"required"`
	Email       string    `json:"email" query:"email" validate:"required,email,unique"`
	PhoneNumber int       `json:"phone_number" validate:"max=14"`
}

type UserDependency struct {
	DB *sqlx.DB
}

type UserValidator struct {
	Validator *validator.Validate
}

var (
	ErrNotExists  = errors.New("user id was not found")
	ErrConnFailed = errors.New("connection failed")
	ErrQuery      = errors.New("execute query error")
	ErrBeginTx    = errors.New("begin transaction error")
	ErrScan       = errors.New("scan error")
	ErrCommit     = errors.New("commit error")
)
