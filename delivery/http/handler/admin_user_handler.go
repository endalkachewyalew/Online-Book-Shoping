package handler

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/Endalk/Exercise/onlineShoping/entity"
	"github.com/Endalk/Exercise/onlineShoping/menu"
)

//OwnerUserHandler ..
type OwnerUserHandler struct {
	temp        *template.Template
	userService menu.UserService
}

//NewOwnerUserHandler ..
func NewOwnerUserHandler(t *template.Template, uServ menu.UserService) *OwnerUserHandler {
	return &OwnerUserHandler{temp: t, userService: uServ}
}

//OwnerUsers ..
func (auh OwnerUserHandler) OwnerUsers(w http.ResponseWriter, r *http.Request) {
	usr, err := auh.userService.Users()
	if err != nil {
		panic(err)
	}
	auh.temp.ExecuteTemplate(w, "Owner.users.layout", usr)
}

//OwnerUsersNew ..
func (auh OwnerUserHandler) OwnerUsersNew(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		//uuid, _ := exec.Command("uuidgen").Output()
		//UUI := string(uuid)
		usr := entity.User{}
		usr.UUID = r.FormValue("uuid")
		usr.FullName = r.FormValue("fullname")
		usr.Email = r.FormValue("email")
		usr.Phone = r.FormValue("phone")
		usr.Password = r.FormValue("password")
		err := auh.userService.StoreUser(usr)
		if err != nil {

			auh.temp.ExecuteTemplate(w, "Owner.message.layout", nil)
		}

		http.Redirect(w, r, "/Owner/users", http.StatusSeeOther)
	} else {

		auh.temp.ExecuteTemplate(w, "Owner.users.new.layout", nil)
	}
}

//OwnerUsersUpdate ..
func (auh OwnerUserHandler) OwnerUsersUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idraw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idraw)
		if err != nil {
			panic(err)
		}
		usr, err := auh.userService.User(id)
		if err != nil {
			panic(err)
		}

		auh.temp.ExecuteTemplate(w, "Owner.users.update.layout", usr)
	} else if r.Method == http.MethodPost {
		rawID := r.FormValue("id")
		id, err := strconv.Atoi(rawID)
		if err != nil {
			http.Redirect(w, r, "/Owner/users", http.StatusSeeOther)
		}
		usr := entity.User{}
		usr.ID = id
		usr.UUID = r.FormValue("uuid")
		usr.FullName = r.FormValue("fullname")
		usr.Email = r.FormValue("email")
		usr.Phone = r.FormValue("phone")
		usr.Password = r.FormValue("password")
		err = auh.userService.UpdateUser(usr)
		if err != nil {

			http.Redirect(w, r, "/Owner/users", http.StatusSeeOther)
		}
		http.Redirect(w, r, "/Owner/users", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/Owner/users", http.StatusSeeOther)
	}
}

//OwnerUserDelete ..
func (auh OwnerUserHandler) OwnerUserDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idraw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idraw)
		if err != nil {
			panic(err)
		}
		err = auh.userService.DeleteUser(id)
		if err != nil {
			panic(err)
		}

	}
	http.Redirect(w, r, "/Owner/users", http.StatusSeeOther)
}
