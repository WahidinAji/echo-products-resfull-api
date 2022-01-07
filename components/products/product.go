package products

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type ProductDomain struct {
	ID    int
	Name  string
	Stock int
	Price float32
}

type Product struct {
	ID    int     `json:"id" query:"id"`
	Name  string  `json:"name" validate:"required"`
	Stock int     `json:"stock" validate:"required"`
	Price float32 `json:"price" validate:"required"`
}

type CustomValidator struct {
	Validator *validator.Validate
}

type MsgDel struct {
	Code   int
	Status string
}

/** Inject Instance DB into Dependency*/

type Dependency struct {
	DB *sql.DB
}
