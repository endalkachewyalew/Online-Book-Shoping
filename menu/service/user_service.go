package service

import (
	"github.com/Endalk/Exercise/restaurant/entity"
	"github.com/Endalk/Exercise/restaurant/menu"
)

//UserServiceImpl ..
type UserServiceImpl struct {
	userRepo menu.UserRepository
}

//NewUserServiceImpl ..
func NewUserServiceImpl(urepo menu.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepo: urepo}
}

//Users ..
func (usi UserServiceImpl) Users() ([]entity.User, error) {
	c, err := usi.userRepo.Users()
	if err != nil {
		return nil, err
	}
	return c, nil
}

//User ..
func (usi UserServiceImpl) User(id int) (entity.User, error) {
	c, err := usi.userRepo.User(id)
	if err != nil {
		return c, err
	}
	return c, nil
}

//UpdateUser ..
func (usi UserServiceImpl) UpdateUser(user entity.User) error {
	err := usi.userRepo.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}

//DeleteUser ..
func (usi UserServiceImpl) DeleteUser(id int) error {
	err := usi.userRepo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

//StoreUser ..
func (usi UserServiceImpl) StoreUser(user entity.User) error {
	err := usi.userRepo.StoreUser(user)
	if err != nil {
		return err
	}
	return nil
}
