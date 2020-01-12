package repository

import (
	"database/sql"
	"errors"

	"github.com/Endalk/Online-Book-Shoping/entity"
)

// AuthorRepositoryImpl implements the menu.AuthorRepository interface
type AuthorRepositoryImpl struct {
	conn *sql.DB
}

// NewAuthorRepositoryImpl will create an object of PsqlAuthorRepository
func NewAuthorRepositoryImpl(Conn *sql.DB) *AuthorRepositoryImpl {
	return &AuthorRepositoryImpl{conn: Conn}
}

// Authors returns all cateogories from the database
func (cri *AuthorRepositoryImpl) Authors() ([]entity.Author, error) {

	rows, err := cri.conn.Query("SELECT * FROM authors")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	ctgs := []entity.Author{}

	for rows.Next() {
		Author := entity.Author{}
		err = rows.Scan(&Author.AuthorID, &Author.Name, &Author.Email, &Author.Address, &Author.PhoneNo, &Author.Description, &Author.Password, &Author.ImagePath, &Author.Rating, &Author.Account, &Author.Activated)
		if err != nil {
			return nil, err
		}
		ctgs = append(ctgs, Author)
	}

	return ctgs, nil
}

// Author returns a Author with a given id
func (cri *AuthorRepositoryImpl) Author(id int) (entity.Author, error) {

	row := cri.conn.QueryRow("SELECT * FROM authors WHERE id = $1", id)

	Author := entity.Author{}

	err := row.Scan(&Author.AuthorID, &Author.Name, &Author.Email, &Author.Address, &Author.PhoneNo, &Author.Description, &Author.Password, &Author.ImagePath, &Author.Rating, &Author.Account, &Author.Activated)
	if err != nil {
		return Author, err
	}

	return Author, nil
}

// UpdateAuthor updates a given object with a new data
func (cri *AuthorRepositoryImpl) UpdateAuthor(c entity.Author) error {

	_, err := cri.conn.Exec("UPDATE authors SET name=$1,description=$2, imagepath=$3,email=$4,phone=$5,address=$6,rating=$7,password=$8,activated=$9 WHERE id=$10", c.Name, c.Description, c.ImagePath, c.Email, c.PhoneNo, c.Address, c.Rating, c.Password, c.Activated, c.AuthorID)
	if err != nil {
		return errors.New("Update has failed")
	}

	return nil
}

// DeleteAuthor removes a Author from a database by its id
func (cri *AuthorRepositoryImpl) DeleteAuthor(id int) error {

	_, err := cri.conn.Exec("DELETE FROM authors WHERE id=$1", id)
	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}

// StoreAuthor stores new Author information to database
func (cri *AuthorRepositoryImpl) StoreAuthor(c entity.Author) error {

	_, err := cri.conn.Exec("INSERT INTO authors (name,email,phone,address,description,imagepath,password) values($1, $2, $3,$4,$5,$6,$7)", c.Name, c.Email, c.PhoneNo, c.Address, c.Description, c.ImagePath, c.Password)
	if err != nil {
		return errors.New("Insertion has failed")
	}

	return nil
}
