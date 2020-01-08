package handler

import (
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Endalk/Exercise/onlineShoping/entity"
	"github.com/Endalk/Exercise/onlineShoping/menu"
)

// OwnerCategoryHandler handles category handler Owner requests
type OwnerCategoryHandler struct {
	tmpl        *template.Template
	categorySrv menu.CategoryService
}

// NewOwnerCategoryHandler initializes and returns new OwnerCateogryHandler
func NewOwnerCategoryHandler(T *template.Template, CS menu.CategoryService) *OwnerCategoryHandler {
	return &OwnerCategoryHandler{tmpl: T, categorySrv: CS}
}

// OwnerCategories handle requests on route /Owner/categories
func (ach *OwnerCategoryHandler) OwnerCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := ach.categorySrv.Categories()
	if err != nil {
		panic(err)
	}
	ach.tmpl.ExecuteTemplate(w, "Owner.categ.layout", categories)
}

// OwnerCategoriesNew hanlde requests on route /Owner/categories/new
func (ach *OwnerCategoryHandler) OwnerCategoriesNew(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		ctg := entity.Category{}
		ctg.Name = r.FormValue("name")
		ctg.Description = r.FormValue("description")

		mf, fh, err := r.FormFile("catimg")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		ctg.Image = fh.Filename

		writeFile(&mf, fh.Filename)

		err = ach.categorySrv.StoreCategory(ctg)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/Owner/categories", http.StatusSeeOther)

	} else {

		ach.tmpl.ExecuteTemplate(w, "Owner.categ.new.layout", nil)

	}
}

// OwnerCategoriesUpdate handle requests on /Owner/categories/update
func (ach *OwnerCategoryHandler) OwnerCategoriesUpdate(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		cat, err := ach.categorySrv.Category(id)

		if err != nil {
			panic(err)
		}

		ach.tmpl.ExecuteTemplate(w, "Owner.categ.update.layout", cat)

	} else if r.Method == http.MethodPost {

		ctg := entity.Category{}
		ctg.ID, _ = strconv.Atoi(r.FormValue("id"))
		ctg.Name = r.FormValue("name")
		ctg.Description = r.FormValue("description")
		mf, fh, err := r.FormFile("catimg")
		if mf != nil {
			ctg.Image = fh.Filename

			if err != nil {
				panic(err)
			}

			defer mf.Close()

			writeFile(&mf, ctg.Image)

			fmt.Println(ctg.Image)
		} else {
			ctg.Image = r.FormValue("image")
		}

		err = ach.categorySrv.UpdateCategory(ctg)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/Owner/categories", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/Owner/categories", http.StatusSeeOther)
	}

}

// OwnerCategoriesDelete handle requests on route /Owner/categories/delete
func (ach *OwnerCategoryHandler) OwnerCategoriesDelete(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		err = ach.categorySrv.DeleteCategory(id)

		if err != nil {
			panic(err)
		}

	}

	http.Redirect(w, r, "/Owner/categories", http.StatusSeeOther)
}

func writeFile(mf *multipart.File, fname string) {

	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	path := filepath.Join(wd, "/../../ui", "assets", "img", fname)
	image, err := os.Create(path)

	if err != nil {
		panic(err)
	}
	defer image.Close()
	io.Copy(image, *mf)
}
