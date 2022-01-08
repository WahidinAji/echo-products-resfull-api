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
}

func (d *UserDependency) FindAll(ctx context.Context) ([]User, error)  {
	conn, err := d.DB.Conn(ctx)
	if err !=nil {
		return nil, ErrConnFailed
	}
	defer conn.Close()
	query := "SELECT user_id, first_name, last_name, email, phone_number FROM users"
	result, err := conn.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf(ErrQuery.Error(),err)
	}
	var users []User
	for result.Next(){
		var user User
		errRes := result.Scan(&user.UserId,&user.FirstName,&user.LastName,&user.Email,&user.PhoneNumber)
		if errRes != nil {
			return nil, fmt.Errorf("Scan error : ",err)
		}
		users = append(users, user)
	}
	return users, nil
}
