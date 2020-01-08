package repository

import (
	"database/sql"
	"errors"

	"github.com/Endalk/Exercise/onlineShoping/entity"
)

//UserRepositoryImpl ..
type UserRepositoryImpl struct {
	conn *sql.DB
}

//NewUserRepositoryImpl ..
func NewUserRepositoryImpl(con *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{conn: con}
}

//Users ..
func (uri UserRepositoryImpl) Users() ([]entity.User, error) {
	rows, err := uri.conn.Query("select * from users")
	if err != nil {
		return nil, errors.New("users failed")
	}
	usrs := []entity.User{}
	for rows.Next() {
		user := entity.User{}
		err = rows.Scan(&user.ID, &user.UUID, &user.FullName, &user.Email, &user.Phone, &user.Password)
		if err != nil {
			return nil, err
		}
		usrs = append(usrs, user)
	}
	return usrs, nil
}

//User ..
func (uri UserRepositoryImpl) User(id int) (entity.User, error) {
	rows := uri.conn.QueryRow("select * from users where id=$1", id)
	usrs := entity.User{}
	err := rows.Scan(&usrs.ID, &usrs.UUID, &usrs.FullName, &usrs.Email, &usrs.Phone, &usrs.Password)
	if err != nil {
		return usrs, errors.New("user by id failed")
	}
	return usrs, nil
}

//UpdateUser ..
func (uri UserRepositoryImpl) UpdateUser(user entity.User) error {
	_, err := uri.conn.Exec("update users set uuid=$1,full_name=$2,email=$3,phone=$4,password=$5 where id=$6", user.UUID, user.FullName, user.Email, user.Phone, user.Password, user.ID)
	if err != nil {
		return errors.New("update failed")
	}
	return nil
}

//DeleteUser ..
func (uri UserRepositoryImpl) DeleteUser(id int) error {
	_, err := uri.conn.Exec("delete from users where id = $1", id)
	if err != nil {
		return errors.New("delete failed")
	}
	return nil
}

//StoreUser ..
func (uri UserRepositoryImpl) StoreUser(user entity.User) error {
	_, err := uri.conn.Exec("insert into users (uuid,full_name,email,phone,password) values ($1,$2,$3,$4,$5)", user.UUID, user.FullName, user.Email, user.Phone, user.Password)
	if err != nil {
		return errors.New("insertion failed")
	}

	return nil
}
