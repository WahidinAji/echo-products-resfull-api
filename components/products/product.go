package products

import "database/sql"

type ProductDomain struct {
	ID    int
	Name  string
	Stock int
	Price float32
}

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Stock int     `json:"stock"`
	Price float32 `json:"price"`
}

type MsgDel struct {
	Code   int
	Status string
}

/** Inject Instance DB into Dependency*/

type Dependency struct {
	DB *sql.DB
}
