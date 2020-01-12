package admin

import "github.com/Endalk/Online-Book-Shoping/entity"

//AuthorRepository ..
type AuthorRepository interface {
	Authors() ([]entity.Author, error)
	Author(id int) (entity.Author, error)
	UpdateAuthor(Author entity.Author) error
	DeleteAuthor(id int) error
	StoreAuthor(Author entity.Author) error
}

//AdminRepository ..
type AdminRepo interface {
	Admins() ([]entity.Admin, error)
	Admin(uname string) (entity.Admin, error)
	UpdateAdmin(Admin entity.Admin) error
	DeleteAdmin(uname string) error
	StoreAdmin(Admin entity.Admin) error
}

//UserRepository ..
type UserRepository interface {
	Users() ([]entity.User, error)
	User(uname string) (entity.User, error)
	UpdateUser(User entity.User) error
	DeleteUser(uname string) error
	StoreUser(User entity.User) error
}

//CommentRepository ..
type CommentRepository interface {
	Comments() ([]entity.Comment, error)
	Comment(id int) (entity.Comment, error)
	UpdateComment(Comment entity.Comment) error
	DeleteComment(id int) error
	StoreComment(User entity.Comment) error
}
