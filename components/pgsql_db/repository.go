package pgsql_db

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

type Repository interface {
	FindAll(ctx context.Context) ([]User, error)
	FindId(ctx context.Context, userId uuid.UUID) (*User, error)
	Update(ctx context.Context, userId uuid.UUID, user User) (*User, error)
	Delete(ctx context.Context, userId uuid.UUID) error
	Save(ctx context.Context, user User) (*User, error)
}

func (d *UserDependency) FindAll(ctx context.Context) ([]User, error) {
	conn, err := d.DB.Conn(ctx)
	if err != nil {
		return nil, fmt.Errorf(ErrConnFailed.Error()+" : ", err)
	}
	defer conn.Close()

	query := "SELECT id, first_name, last_name, email, phone_number FROM users ORDER BY created_at DESC"
	result, err := conn.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf(ErrQuery.Error()+" : ", err)
	}

	var users []User
	for result.Next() {
		var user User
		errRes := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.PhoneNumber)
		if errRes != nil {
			return nil, fmt.Errorf(ErrScan.Error()+" : ", err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (d *UserDependency) FindId(ctx context.Context, userId uuid.UUID) (*User, error) {
	conn, err := d.DB.Conn(ctx)
	if err != nil {
		return nil, fmt.Errorf(ErrConnFailed.Error()+" : ", err)
	}
	defer conn.Close()

	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf(ErrBeginTx.Error()+" : ", err)
	}
	defer tx.Rollback()

	var exists bool
	query := "SELECT EXISTS ( SELECT id FROM users WHERE id=$1)"
	err = tx.QueryRowContext(ctx, query, userId.String()).Scan(&exists)
	if err != nil {
		return nil, fmt.Errorf(ErrQuery.Error()+" !exists : ", err)
	}
	if !exists {
		return nil, fmt.Errorf(ErrNotExists.Error()+" : ", err)
	}

	query = "SELECT id, first_name, last_name, email, phone_number FROM users WHERE id=$1"
	row, err := tx.QueryContext(ctx, query, userId.String())
	if err != nil {
		return nil, fmt.Errorf(ErrQuery.Error()+" : ", err)
	}

	var user User
	if row.Next() {
		err = row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.PhoneNumber)
		if err != nil {
			return nil, fmt.Errorf(ErrScan.Error()+" : ", err)
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf(ErrCommit.Error()+" : ", err)
	}
	return &user, nil
}

func (d *UserDependency) Update(ctx context.Context, userId uuid.UUID, user User) (*User, error) {
	conn, err := d.DB.Conn(ctx)
	if err != nil {
		return nil, fmt.Errorf(ErrConnFailed.Error()+" : ", err)
	}
	defer conn.Close()

	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf(ErrBeginTx.Error()+" : ", err)
	}
	defer tx.Rollback()

	var exists bool
	query := "SELECT EXISTS ( SELECT id FROM users WHERE id=$1)"
	err = tx.QueryRowContext(ctx, query, userId.String()).Scan(&exists)
	if err != nil {
		return nil, fmt.Errorf(ErrQuery.Error()+" !exists : ", err)
	}

	if !exists {
		return nil, fmt.Errorf(ErrNotExists.Error()+" : ", err)
	}

	query = "UPDATE users SET first_name=$1,last_name=$2,email=$3,phone_number=$4 WHERE id=$5"
	_, execErr := tx.ExecContext(ctx, query, &user.FirstName, &user.LastName, &user.Email, &user.PhoneNumber, userId.String())
	if execErr != nil {
		return nil, fmt.Errorf(ErrQuery.Error()+" Exec : ", err)
	}
	user.ID = userId
	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf(ErrCommit.Error()+" : ", err)
	}
	return &user, nil
}

func (d *UserDependency) Delete(ctx context.Context, userId uuid.UUID) error {
	conn, err := d.DB.Conn(ctx)
	if err != nil {
		return fmt.Errorf(ErrConnFailed.Error()+" : ", err)
	}
	defer conn.Close()

	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf(ErrBeginTx.Error()+" : ", err)
	}
	defer tx.Rollback()

	var exists bool
	query := "SELECT EXISTS ( SELECT id FROM users WHERE id=$1)"
	err = tx.QueryRowContext(ctx, query, userId.String()).Scan(&exists)
	if err != nil {
		return fmt.Errorf(ErrQuery.Error()+" !exists : ", err)
	}

	if !exists {
		return fmt.Errorf(ErrNotExists.Error()+" : ", err)
	}

	query = "DELETE FROM users WHERE id=$1"
	_, execErr := tx.ExecContext(ctx, query, userId.String())
	if execErr != nil {
		return fmt.Errorf(ErrQuery.Error()+" Exec : ", err)
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf(ErrCommit.Error()+" : ", err)
	}
	return nil
}

func (d *UserDependency) Save(ctx context.Context, user User) (*User, error) {
	conn, err := d.DB.Conn(ctx)
	if err != nil {
		return nil, fmt.Errorf(ErrConnFailed.Error()+" : ", err)
	}
	defer conn.Close()

	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf(ErrBeginTx.Error()+" : ", err)
	}
	defer tx.Rollback()

	var exists bool
	query := "SELECT 1 FROM users WHERE email = $1"
	err = tx.QueryRowContext(ctx, query, user.Email).Scan(&exists)
	if err == nil {
		//if exists {
		//	return nil, fmt.Errorf(ErrExists.Error()+" exists 2 : ", err)
		//}
		return nil, fmt.Errorf(ErrExists.Error()+" "+ErrQuery.Error()+" exists 1 : ", err)
	}

	query = `INSERT INTO users (id,first_name,last_name,email,phone_number) VALUES ($1,$2,$3,$4,$5)`
	_, err = tx.ExecContext(ctx, query, user.ID, user.FirstName, user.LastName, user.Email, user.PhoneNumber)
	if err != nil {
		return nil, fmt.Errorf(ErrQuery.Error()+" Exec : ", err)
	}
	errCommit := tx.Commit()
	if errCommit != nil {
		return nil, fmt.Errorf(ErrCommit.Error()+" Exec : ", err)
	}
	return &user, nil
}
