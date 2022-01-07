package products

import (
	"context"
	"database/sql"
	"errors"
)

func (d *Dependency) FindAll(ctx context.Context) ([]Product, error) {
	conn, err := d.DB.Conn(ctx)
	if err != nil {
		return nil, errors.New("Connection failed!!!")
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
		return nil, errors.New("Connection failed!!!")
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
	res, execErr := tx.ExecContext(ctx, "UPDATE products SET name=?, stock=?, price=? WHERE id=?", &product.Name, &product.Stock, &product.Price, postId)
	if execErr != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, err
		}
		return nil, execErr
	}
	row, errRow := res.RowsAffected()
	if errRow != nil {
		errRoll := tx.Rollback()
		if errRoll != nil {
			return nil, errRoll
		}
		return nil, errRow
	}
	if row == 0 {
		errRoll := tx.Rollback()
		if errRoll != nil {
			return nil, errRoll
		}
		return nil, nil
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

func (d *Dependency) Delete(ctx context.Context, postId int) error {
	conn, err := d.DB.Conn(ctx)
	if err != nil {
		return errors.New("Connection failed!!!")
	}
	defer conn.Close()
	tx, err := d.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx,"DELETE FROM products WHERE id =? ",postId)
	if err != nil {
		errRoll := tx.Rollback()
		if errRoll != nil {
			return errRoll
		}
		return err
	}
	errCommit := tx.Commit()
	if errCommit != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return errCommit
	}
	return nil
}

func (d *Dependency)Save(ctx context.Context, product Product) (*Product, error)  {
	conn, err := d.DB.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	tx, err := d.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}
	query := `INSERT INTO products (name,stock,price) VALUES (?,?,?)`
	exec, err := tx.ExecContext(ctx,query, product.Name,product.Stock,product.Price)
	if err != nil {
		errRoll := tx.Rollback()
		if errRoll != nil {
			return nil, errRoll
		}
		return nil, err
	}
	id, err := exec.LastInsertId()
	if err != nil {
		errRoll := tx.Rollback()
		if errRoll != nil {
			return nil, errRoll
		}
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
	product.ID = int(id)
	return &product, nil
}
