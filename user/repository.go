package user

import "github.com/Endalk/Online-Book-Shoping/entity"

//Repository -
type Repository interface {
	User(id string) (entity.User, error)
	UpdateUser(user entity.User) error
	DeleteUser(id string) error
	AddUser(user entity.User) error
	Users() ([]entity.User, error)
	//There are things that will be added on the flow
}
