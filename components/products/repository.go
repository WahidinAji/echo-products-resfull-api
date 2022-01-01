package products

import (
	"context"
	"database/sql"
)

func (d *Dependency) FindAll(ctx context.Context) ([]Product, error) {
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
	for result.Next() {
		var product Product
		errResult := result.Scan(&product.ID, &product.Name, &product.Stock, &product.Price)
		if errResult != nil {
			err := tx.Rollback()
			if err != nil {
				return nil, err
			}
			return nil, errResult
		}
		products = append(products, product)
	}
	errCommit := tx.Commit()
	if errCommit != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, err
		}
		return nil, errCommit
	}
	return products, nil
}

func (d *Dependency) FindId(ctx context.Context, postId int) (*Product, error) {
	conn, err := d.DB.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	tx, err := d.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}

	row, execErr := tx.QueryContext(ctx, "SELECT id, name, stock, price FROM products WHERE id=?", postId)
	if execErr != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, err
		}
		return nil, execErr
	}
	defer row.Close()

	product := &Product{}
	if row.Next() {
		err := row.Scan(&product.ID, &product.Name, &product.Stock, &product.Price)
		if err != nil {
			return nil, err
		}
		errCommit := tx.Commit()
		if errCommit != nil {
			err := tx.Rollback()
			if err != nil {
				return nil, err
			}
			return nil, errCommit
		}
		return product, nil
	} else {
		err := tx.Rollback()
		if err != nil {
			return nil, err
		}
		return nil, err
	}
}

func (d *Dependency) Update(ctx context.Context, postId int, product Product) (*Product, error) {
	conn, err := d.DB.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	tx, err := d.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}
	_, execErr := tx.ExecContext(ctx, "UPDATE products SET name=?, stock=?, price=? WHERE id=?", &product.Name, &product.Stock, &product.Price, postId)
	if execErr != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, err
		}
		return nil, execErr
	}

	product.ID = postId
	errCommit := tx.Commit()
	if errCommit != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, err
		}
		return nil, errCommit
	}
	return &product, nil
}
