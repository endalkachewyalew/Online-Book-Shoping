package service

import (
	"github.com/Endalk/Online-Book-Shoping/admin/repository"
	"github.com/Endalk/Online-Book-Shoping/entity"
)

//AdminService -
type AdminService struct {
	repository repository.AdminRepositoryImpl
}

//NewAdminService -
func NewAdminService(repo repository.AdminRepositoryImpl) *AdminService {
	return &AdminService{repository: repo}
}

//Admins -
func (ar *AdminService) Admins() ([]entity.Admin, error) {
	admins, err := ar.repository.Admins()
	if err != nil {
		return admins, err
	}
	return admins, nil
}

//Admin -
func (ar *AdminService) Admin(username string) (entity.Admin, error) {
	admin, err := ar.repository.Admin(username)
	if err != nil {
		return admin, err
	}
	return admin, nil
}

//AddAdmin -
func (ar *AdminService) AddAdmin(admin entity.Admin) error {
	err := ar.repository.StoreAdmin(admin)
	if err != nil {
		return err
	}
	return nil
}

//UpdateAdmin -
func (ar *AdminService) UpdateAdmin(admin entity.Admin) error {
	err := ar.repository.UpdateAdmin(admin)
	if err != nil {
		return err
	}
	return nil
}

//DeleteAdmin -
func (ar *AdminService) DeleteAdmin(id string) error {
	err := ar.repository.DeleteAdmin(id)
	if err != nil {
		return err
	}
	return nil
}
