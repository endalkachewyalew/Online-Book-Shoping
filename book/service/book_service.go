package service

import (
	"github.com/Endalk/Online-Book-Shoping/entity"
	"github.com/Endalk/Online-Book-Shoping/book"
)

//BookService -
type BookService struct {
	repository book.BookRepository
}

//NewBookService -
func NewBookService(repo book.BookService) *BookService {
	return &BookService{repository: repo}
}

//Books -
func (ms *BookService) Books() ([]entity.Book, error) {
	books, err := ms.repository.Books()
	if err != nil {
		return books, err
	}
	return books, nil
}

//Book -
func (ms *BookService) Book(id int) (entity.Book, error) {
	book, err := ms.repository.Book(id)
	if err != nil {
		return book, err
	}
	return book, nil
}

//UpdateBook -
func (ms *BookService) UpdateBook(book entity.Book) error {
	err := ms.repository.UpdateBook(book)
	if err != nil {
		return err
	}
	return nil
}

//DeleteBook -
func (ms *BookService) DeleteBook(id int) error {
	err := ms.repository.DeleteBook(id)
	if err != nil {
		return err
	}
	return nil
}

//AddBook -
func (ms *BookService) AddBook(book entity.Book) error {
	err := ms.repository.AddBook(book)
	if err != nil {
		return err
	}
	return nil
}

//BookByAuthorOwner -
func (ms *BookService) BookByAuthorOwner(id int) ([]entity.Book, error) {
	books, err := ms.repository.BookByAuthorOwner(id)
	if err != nil {
		return books, err
	}
	return books, nil
}
