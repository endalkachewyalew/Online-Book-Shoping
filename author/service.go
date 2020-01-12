package author

import "github.com/Endalk/Online-Book-Shoping/entity"

type BookService interface {
	Books() ([]entity.Book, error)
	Book(id int) (entity.Book, error)
	UpdateBook(book entity.Book) error
	DeleteBook(id int) error
	AddBook(book entity.Book) error
	BookByAuthorOwner(id int) ([]entity.Book, error)
}
