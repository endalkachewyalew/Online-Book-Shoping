package user

import "github.com/Endalk/Online-Book-Shoping/entity"

//Service -
//This are definitions of the services we give
type Service interface {
	User(id string) (entity.User, error)
	UpdateUser(user entity.User) error
	DeleteUser(id string) error
	AddUser(user entity.User) error
	Users() ([]entity.User, error)
	//There are things that will be added on the flow
}
