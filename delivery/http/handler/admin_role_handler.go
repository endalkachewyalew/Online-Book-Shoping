package handler

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/Endalk/Exercise/onlineShoping/entity"
	"github.com/Endalk/Exercise/onlineShoping/menu"
)

//OwnerRoleHandler handles roles related requests
type OwnerRoleHandler struct {
	tmpl    *template.Template
	roleSrv menu.RoleService
}

//NewOwnerRoleHandler ..
func NewOwnerRoleHandler(T *template.Template, RS menu.RoleService) *OwnerRoleHandler {
	return &OwnerRoleHandler{tmpl: T, roleSrv: RS}
}

//OwnerRoles ..
func (arh OwnerRoleHandler) OwnerRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := arh.roleSrv.Roles()
	if err != nil {
		panic(err)
	}
	arh.tmpl.ExecuteTemplate(w, "Owner.roles.layout", roles)
}

//OwnerRoleNew handles routes at /Owner/roles/new
func (arh OwnerRoleHandler) OwnerRoleNew(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		rol := entity.Role{}
		rol.Name = r.FormValue("name")
		err := arh.roleSrv.StoreRole(rol)
		if err != nil {
			http.Redirect(w, r, "/Owner/roles", http.StatusSeeOther)
		}
		http.Redirect(w, r, "/Owner/roles", http.StatusSeeOther)
	} else {
		arh.tmpl.ExecuteTemplate(w, "Owner.roles.new.layout", nil)

	}
}

//OwnerRoleUpdate ..
func (arh OwnerRoleHandler) OwnerRoleUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			panic(err)
		}
		rol, err := arh.roleSrv.Role(id)
		if err != nil {
			panic(err)
		}
		arh.tmpl.ExecuteTemplate(w, "Owner.roles.update.layout", rol)
	} else if r.Method == http.MethodPost {
		rawID := r.FormValue("id")
		id, err := strconv.Atoi(rawID)
		if err != nil {
			panic(err)
		}

		rol := entity.Role{}
		rol.ID = id
		rol.Name = r.FormValue("name")
		err = arh.roleSrv.UpdateRole(rol)
		if err != nil {
			panic(err)
		}
		http.Redirect(w, r, "/Owner/roles", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/Owner/roles", http.StatusSeeOther)
	}

}

//OwnerRoleDelete ..
func (arh OwnerRoleHandler) OwnerRoleDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idraw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idraw)
		if err != nil {
			panic(err)
		}
		err = arh.roleSrv.DeleteRole(id)
		if err != nil {
			panic(err)
		}

	}
	http.Redirect(w, r, "/Owner/roles", http.StatusSeeOther)
}
