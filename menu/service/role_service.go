package service

import (
	"github.com/Endalk/Exercise/restaurant/entity"
	"github.com/Endalk/Exercise/restaurant/menu"
)

//RoleServiceImpl ..

//RoleServiceImpl -
type RoleServiceImpl struct {
	repository menu.RoleRepository
}

//NewRoleServiceImpl -
func NewRoleServiceImpl(repo menu.RoleRepository) *RoleServiceImpl {
	return &RoleServiceImpl{repository: repo}
}

//Roles -
func (rs *RoleServiceImpl) Roles() ([]entity.Role, error) {
	roles, err := rs.repository.Roles()
	if err != nil {
		return roles, nil
	}
	return roles, nil
}

//Role -
func (rs *RoleServiceImpl) Role(id int) (entity.Role, error) {
	role, err := rs.repository.Role(id)
	if err != nil {
		return role, err
	}
	return role, err
}

//UpdateRole -
func (rs *RoleServiceImpl) UpdateRole(role entity.Role) error {
	err := rs.repository.UpdateRole(role)
	if err != nil {
		return err
	}
	return nil
}

//StoreRole -
func (rs *RoleServiceImpl) StoreRole(role entity.Role) error {
	err := rs.repository.StoreRole(role)
	if err != nil {
		return err
	}
	return nil
}

//DeleteRole -
func (rs *RoleServiceImpl) DeleteRole(id int) error {
	err := rs.repository.DeleteRole(id)
	if err != nil {
		return err
	}
	return nil
}
