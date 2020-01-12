package repository

import (
	"database/sql"
	"errors"

	"github.com/Endalk/Online-Book-Shoping/entity"
)

//BookRepository -
type BookRepository struct {
	conn *sql.DB
}

//NewBookRepository -
func NewBookRepository(Conn *sql.DB) *BookRepository {
	return &BookRepository{conn: Conn}
}

//Books -
func (mr *BookRepository) Books() ([]entity.Book, error) {
	books := make([]entity.Book, 0)
	query := "select * from books"
	data, err := mr.conn.Query(query)
	if err != nil {
		return books, errors.New("No user is found")
	}
	for data.Next() {
		var book entity.Book
		data.Scan(&book.ID, &book.Name, &book.Type, &book.Owner, &book.PricePerDay, &book.OnDiscount, &book.Discount, &book.OnSale, &book.ImagePath) //all the datas that will be added in the category
		books = append(books, book)
	}
	if err := data.Err(); err != nil {
		return books, errors.New("Some error is occured")
	}
	return books, nil
}

//Book -
func (mr *BookRepository) Book(id int) (entity.Book, error) {
	book := entity.Book{}
	query := "select * from books where id=$1"
	err := mr.conn.QueryRow(query, id).Scan(&book.ID, &book.Name, &book.Type, &book.Owner, &book.PricePerDay, &book.ImagePath)
	if err != nil {
		return book, err
	}
	return book, nil
}

//UpdateBook -
func (mr *BookRepository) UpdateBook(book entity.Book) error {
	query := "update books set name=$1,type=$2,owner=$3,priceperday=$4,ondiscount=$5,discount=$6,onsale=$7,imagepath=$8 where id=$9"
	_, err := mr.conn.Exec(query, book.Name, book.Type, book.Owner, book.PricePerDay, book.OnDiscount, book.Discount, book.OnSale, book.ImagePath, book.ID)
	if err != nil {
		return err
	}
	return nil
}

//DeleteBook -
func (mr *BookRepository) DeleteBook(id int) error {
	query := "delete from books where id=$1"
	_, err := mr.conn.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

//AddBook -
func (mr *BookRepository) AddBook(book entity.Book) error {
	query := "insert into books(name, type, owner,priceperday, ondiscount, discount, onsale, imagepath) values($1,$2,$3,$4,$5,$6,$7,$8)"

	_, err := mr.conn.Exec(query, book.Name, book.Type, book.Owner, book.PricePerDay, book.OnDiscount, book.Discount, book.OnSale, book.ImagePath)
	if err != nil {
		return err
	}
	return nil
}

//BookByAuthorOwner -
func (mr *BookRepository) BookByAuthorOwner(id int) ([]entity.Book, error) {
	books := make([]entity.Book, 0)
	query := "select * from books where owner=$1"
	data, err := mr.conn.Query(query, id)
	if err != nil {
		return books, errors.New("No user is found")
	}
	for data.Next() {
		var book entity.Book
		data.Scan(&book.ID, &book.Name, &book.Type, &book.Owner, &book.PricePerDay, &book.OnDiscount, &book.Discount, &book.OnSale, &book.ImagePath) //all the datas that will be added in the category
		books = append(books, book)
	}
	if err := data.Err(); err != nil {
		return books, errors.New("Some error is occured")
	}
	return books, nil
}
