package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/Endalk/Exercise/onlineShoping/delivery/http/handler"
	"github.com/Endalk/Exercise/onlineShoping/menu/repository"
	"github.com/Endalk/Exercise/onlineShoping/menu/service"
)

func main() {

	dbconn, err := sql.Open("postgres", "postgres://postgres:password1234@localhost/onlineshoping?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	if err := dbconn.Ping(); err != nil {
		panic(err)
	}

	tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))
	//user
	usrRepo := repository.NewUserRepositoryImpl(dbconn)
	usrServ := service.NewUserServiceImpl(usrRepo)
	usrHandl := handler.NewOwnerUserHandler(tmpl, usrServ)
	//category
	categoryRepo := repository.NewCategoryRepositoryImpl(dbconn)
	categoryServ := service.NewCategoryServiceImpl(categoryRepo)
	OwnerCatgHandler := handler.NewOwnerCategoryHandler(tmpl, categoryServ)
	//role
	roleRepo := repository.NewRoleRepositoryImpl(dbconn)
	roleSrv := service.NewRoleServiceImpl(roleRepo)
	roleHandler := handler.NewOwnerRoleHandler(tmpl, roleSrv)
	//menu
	menuHandler := handler.NewMenuHandler(tmpl, categoryServ)

	fs := http.FileServer(http.Dir("../../ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", menuHandler.Index)
	http.HandleFunc("/about", menuHandler.About)
	http.HandleFunc("/contact", menuHandler.Contact)
	http.HandleFunc("/menu", menuHandler.Menu)
	//http.HandleFunc("/home", menuHandler.home)

	http.HandleFunc("/Owner", menuHandler.Owner)

	http.HandleFunc("/Owner/categories", OwnerCatgHandler.OwnerCategories)
	http.HandleFunc("/Owner/categories/new", OwnerCatgHandler.OwnerCategoriesNew)
	http.HandleFunc("/Owner/categories/update", OwnerCatgHandler.OwnerCategoriesUpdate)
	http.HandleFunc("/Owner/categories/delete", OwnerCatgHandler.OwnerCategoriesDelete)

	http.HandleFunc("/Owner/users", usrHandl.OwnerUsers)
	http.HandleFunc("/Owner/users/new", usrHandl.OwnerUsersNew)
	http.HandleFunc("/Owner/users/update", usrHandl.OwnerUsersUpdate)
	http.HandleFunc("/Owner/users/delete", usrHandl.OwnerUserDelete)

	http.HandleFunc("/Owner/roles", roleHandler.OwnerRoles)
	http.HandleFunc("/Owner/roles/new", roleHandler.OwnerRoleNew)
	http.HandleFunc("/Owner/roles/update", roleHandler.OwnerRoleUpdate)
	http.HandleFunc("/Owner/roles/delete", roleHandler.OwnerRoleDelete)

	http.ListenAndServe(":8181", nil)
}
