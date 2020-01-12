package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/Endalk/Online-Book-Shoping/admin"
	"github.com/Endalk/Online-Book-Shoping/entity"
	"github.com/julienschmidt/httprouter"
)

// AdminAuthorHandler handles Admin handler admin requests
type AdminAuthorHandler struct {
	tmpl       *template.Template
	AuthorSrv admin.AuthorService
}

// NewAdminAuthorHandler initializes and returns new AdminAuthorHandler
func NewAdminAuthorHandler(T *template.Template, CS admin.AuthorService) *AdminAuthorHandler {
	return &AdminAuthorHandler{tmpl: T, AuthorSrv: CS}
}

// AdminAuthors handle requests on route /admin/Authors
func (ach *AdminAuthorHandler) AdminAuthors(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	Authors, err := ach.AuthorSrv.Authors()
	if err != nil {
		panic(err)
	}
	ach.tmpl.ExecuteTemplate(w, "admin.author.layout", Authors)
}

// AdminAuthorsNew hanlde requests on route /admin/Authors/new
func (ach *AdminAuthorHandler) AdminAuthorsNew(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	if r.Method == http.MethodPost {

		ctg := entity.Author{}
		ctg.Name = r.FormValue("name")
		ctg.Description = r.FormValue("description")
		ctg.Email = r.FormValue("email")
		ctg.PhoneNo = r.FormValue("phone")
		ctg.Password = r.FormValue("password")
		ctg.Address = r.FormValue("address")
		mf, fh, err := r.FormFile("catimg")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		ctg.ImagePath = fh.Filename

		writeFile(&mf, fh.Filename)

		err = ach.AuthorSrv.StoreAuthor(ctg)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/admin/author", http.StatusSeeOther)

	} else {

		ach.tmpl.ExecuteTemplate(w, "admin.author.new.layout", nil)

	}
}

// AdminAuthorsUpdate handle requests on /admin/Authors/update
func (ach *AdminAuthorHandler) AdminAuthorsUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		cat, err := ach.AuthorSrv.Author(id)
		if err != nil {
			panic(err)
		}

		ach.tmpl.ExecuteTemplate(w, "admin.author.update.layout", cat)

	} else if r.Method == http.MethodPost {

		ctg := entity.Author{}
		ctg.Name = r.FormValue("name")
		ctg.Description = r.FormValue("description")
		ctg.Email = r.FormValue("email")
		ctg.Password = r.FormValue("password")
		ctg.Address = r.FormValue("address")
		mf, fh, err := r.FormFile("catimg")
		if mf != nil {
			ctg.ImagePath = fh.Filename

			if err != nil {
				panic(err)
			}

			defer mf.Close()

			writeFile(&mf, ctg.ImagePath)

			fmt.Println(ctg.ImagePath)
		} else {
			ctg.ImagePath = r.FormValue("catimg")
		}

		err = ach.AuthorSrv.UpdateAuthor(ctg)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/author", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/author", http.StatusSeeOther)
	}

}

// AdminAuthorsDelete handle requests on route /admin/categories/delete
func (ach *AdminAuthorHandler) AdminAuthorsDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		err = ach.AuthorSrv.DeleteAuthor(id)

		if err != nil {
			panic(err)
		}

	}

	http.Redirect(w, r, "/admin/author", http.StatusSeeOther)
}
