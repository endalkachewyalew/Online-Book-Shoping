package admin

import "github.com/Endalk/Online-Book-Shoping/entity"

//AuthorService ..
type AuthorService interface {
	Authors() ([]entity.Author, error)
	Author(id int) (entity.Author, error)
	UpdateAuthor(Author entity.Author) error
	DeleteAuthor(id int) error
	StoreAuthor(Author entity.Author) error
}

//AdminService ..
type AdminService interface {
	Admins() ([]entity.Admin, error)
	Admin(uname string) (entity.Admin, error)
	UpdateAdmin(Admin entity.Admin) error
	DeleteAdmin(uname string) error
	StoreAdmin(Admin entity.Admin) error
}

//UserService ..
type UserService interface {
	Users() ([]entity.User, error)
	User(uname string) (entity.User, error)
	UpdateUser(User entity.User) error
	DeleteUser(uname string) error
	StoreUser(User entity.User) error
}

//CommentService ..
type CommentService interface {
	Comments() ([]entity.Comment, error)
	Comment(id int) (entity.Comment, error)
	UpdateComment(Comment entity.Comment) error
	DeleteComment(id int) error
	StoreComment(User entity.Comment) error
}
