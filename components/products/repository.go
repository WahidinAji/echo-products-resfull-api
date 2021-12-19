package products

import (
	"context"
	"database/sql"
)

func (d *Dependency) FindAll(ctx context.Context) ([]Product, error)  {
	conn, err := d.DB.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	tx, err := d.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}
	result, execErr := tx.QueryContext(ctx, "SELECT id, name, stock, price FROM products")
	if execErr != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, err
		}
		return []Product{}, execErr
	}
	var products []Product
	for result.Next(){
		var product Product
		errResult := result.Scan(&product.ID,&product.Name,&product.Stock,&product.Price)
		if errResult != nil {
			err := tx.Rollback()
			if err != nil {
				return nil, err
			}
			return nil, errResult
		}
		products = append(products, product)
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return products, nil
}
