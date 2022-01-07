package products

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

func (d *Dependency) FindId(ctx context.Context, postId int) (Product, error) {
	conn, err := d.DB.Conn(ctx)
	if err != nil {
		return Product{}, errors.New("Connection failed!!!")
	}
	defer conn.Close()

	//check data exists or nor
	row, err := d.DB.QueryContext(ctx,"SELECT EXISTS ( SELECT id FROM products WHERE id=?)",postId)
	if err != nil {
		return Product{},err
	}
	defer row.Close()
	var exist bool
	if row.Next(){
		errScan := row.Scan(&exist)
		if errScan != nil {
			return Product{},err
		}
	}
	if !exist {
		return Product{}, errors.New("ID Was Not Found")
	}

	//if data is exists
	row, err = d.DB.QueryContext(ctx, "SELECT id, name, stock, price FROM products WHERE id=?", postId)
	if err != nil {
		return Product{}, err
	}
	defer row.Close()
	product := Product{}
	if row.Next() {
		err = row.Scan(&product.ID, &product.Name, &product.Stock, &product.Price)
		if err != nil {
			return Product{}, err
		}
	}
	return product, nil
}

func (d *Dependency) Update2(ctx context.Context, postId int, product Product) (*Product, error) {
	conn, err := d.DB.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	tx, err := d.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}
	//check ID
	checkId, err := tx.QueryContext(ctx,"SELECT EXISTS ( SELECT id FROM products WHERE id=?)",postId)
	if err != nil {
		return nil, err
	}
	defer checkId.Close()
	var exist bool
	if checkId.Next(){
		errScan := checkId.Scan(&exist)
		if errScan != nil {
			return nil, err
		}
	}
	if !exist {
		return nil, errors.New("ID Was Not Found")
	}

	//if check id was successfully and ID was found
	_, execErr := tx.ExecContext(ctx, "UPDATE products SET name=?, stock=?, price=? WHERE id=?", &product.Name, &product.Stock, &product.Price, postId)
	if execErr != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, err
		}
		return nil, errors.New("Execute error")
	}
	//row, errRow := res.RowsAffected()
	//if errRow != nil {
	//	errRoll := tx.Rollback()
	//	if errRoll != nil {
	//		return nil, errRoll
	//	}
	//	return nil, errRow
	//}
	//if row == 0 {
	//	errRoll := tx.Rollback()
	//	if errRoll != nil {
	//		return nil, errRoll
	//	}
	//	return nil, nil
	//}

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

	rows, err := tx.QueryContext(ctx,"SELECT EXISTS ( SELECT id FROM products WHERE id=?)",postId)
	if err != nil {
		return nil, errors.New("err query")
	}
	var exists bool
	if rows.Next() {
		err = rows.Scan(&exists)
		if err!=nil {
			return nil, err
		}
	}
	if !exists {
		return nil, errors.New("not found")
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

func (d *Dependency) Delete2(ctx context.Context, postId int) error {
	conn, err := d.DB.Conn(ctx)
	if err != nil {
		return errors.New("Connection failed!!!")
	}
	defer conn.Close()
	tx, err := d.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	row, err := tx.QueryContext(ctx,"SELECT EXISTS ( SELECT id FROM products WHERE id=?)",postId)
	if err != nil {
		return err
	}
	defer row.Close()
	var exist bool
	if row.Next(){
		errScan := row.Scan(&exist)
		if errScan != nil {
			return err
		}
	}
	if !exist {
		return errors.New("ID Was Not Found")
	}

	//if check id was successfully and ID was found
	_, err = tx.ExecContext(ctx,"DELETE FROM products WHERE id =? ",postId)
	if err != nil {
		errRoll := tx.Rollback()
		if errRoll != nil {
			return errRoll
		}
		return errors.New("exec error")
	}
	errCommit := tx.Commit()
	if errCommit != nil {
		err := tx.Rollback()
		if err != nil {
			return errors.New("roll error")
		}
		return errors.New("Commit error")
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

func (d *Dependency) Delete(ctx context.Context, postId int) error {
	conn, err := d.DB.Conn(ctx)
	if err != nil {
		return errors.New("Connection failed!!!")
	}
	defer conn.Close()
	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	row, err := tx.QueryContext(ctx,"SELECT EXISTS ( SELECT * FROM products WHERE id=?)",postId)
	if err != nil {
		return err
	}
	defer row.Close()
	var exist bool
	if row.Next(){
		errScan := row.Scan(&exist)
		if errScan != nil {
			return err
		}
	}
	if !exist {
		return errors.New("ID Was Not Found")
	}

	//if check id was successfully and ID was found
	_, err = tx.ExecContext(ctx,"DELETE FROM products WHERE id =? ",postId)
	if err != nil {
		errRoll := tx.Rollback()
		if errRoll != nil {
			fmt.Println(errRoll)
			return errors.New("rollback error")
		}
		return errors.New("exec error")
	}
	errCommit := tx.Commit()
	if errCommit != nil {
		err := tx.Rollback()
		if err != nil {
			return errors.New("roll error")
		}
		return errors.New("Commit error")
	}
	return nil
}
