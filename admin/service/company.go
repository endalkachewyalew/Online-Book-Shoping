package service

import (
	"github.com/Endalk/Online-Book-Shoping/admin"
	"github.com/Endalk/Online-Book-Shoping/entity"
)

// AuthorServiceImpl implements menu.AuthorService interface
type AuthorServiceImpl struct {
	AuthorRepo admin.AuthorRepository
}

// NewAuthorServiceImpl will create new AuthorService object
func NewAuthorServiceImpl(CatRepo admin.AuthorRepository) *AuthorServiceImpl {
	return &AuthorServiceImpl{AuthorRepo: CatRepo}
}

// Authors ..() returns list of authors
func (cs *AuthorServiceImpl) Authors() ([]entity.Author, error) {

	authors, err := cs.AuthorRepo.Authors()

	if err != nil {
		return nil, err
	}

	return authors, nil
}

// StoreAuthor persists new Author information
func (cs *AuthorServiceImpl) StoreAuthor(Author entity.Author) error {

	err := cs.AuthorRepo.StoreAuthor(Author)

	if err != nil {
		return err
	}

	return nil
}

// Author returns a Author object with a given id
func (cs *AuthorServiceImpl) Author(id int) (entity.Author, error) {

	c, err := cs.AuthorRepo.Author(id)

	if err != nil {
		return c, err
	}

	return c, nil
}

// UpdateAuthor updates a cateogory with new data
func (cs *AuthorServiceImpl) UpdateAuthor(Author entity.Author) error {

	err := cs.AuthorRepo.UpdateAuthor(Author)

	if err != nil {
		return err
	}

	return nil
}

// DeleteAuthor delete a Author by its id
func (cs *AuthorServiceImpl) DeleteAuthor(id int) error {

	err := cs.AuthorRepo.DeleteAuthor(id)
	if err != nil {
		return err
	}
	return nil
}
