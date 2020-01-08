package menu

import "github.com/Endalk/Exercise/onlineShoping/entity"

// CategoryService specifies food menu category services
type CategoryService interface {
	Categories() ([]entity.Category, error)
	Category(id int) (entity.Category, error)
	UpdateCategory(category entity.Category) error
	DeleteCategory(id int) error
	StoreCategory(category entity.Category) error
}

//RoleService specifies application user role related
type RoleService interface {
	Roles() ([]entity.Role, error)
	Role(id int) (entity.Role, error)
	UpdateRole(role entity.Role) error
	DeleteRole(id int) error
	StoreRole(role entity.Role) error
}

//UserService ..
type UserService interface {
	Users() ([]entity.User, error)
	User(id int) (entity.User, error)
	UpdateUser(user entity.User) error
	DeleteUser(id int) error
	StoreUser(user entity.User) error
}
