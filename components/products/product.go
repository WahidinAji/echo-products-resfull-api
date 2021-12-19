package products

import "database/sql"

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Stock int     `json:"stock"`
	Price float32 `json:"price"`
}

/** Inject Instance DB into Dependency*/
type Dependency struct {
	DB *sql.DB
}