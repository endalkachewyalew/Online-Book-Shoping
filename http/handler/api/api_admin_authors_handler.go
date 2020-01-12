package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Endalk/Online-Book-Shoping/admin"
	"github.com/Endalk/Online-Book-Shoping/entity"
	"github.com/julienschmidt/httprouter"
)

// AdminAuthorHandler handles Author related http requests
type AdminAuthorHandler struct {
	AuthorService admin.AuthorService
}

// NewAdminAuthorHandler returns new AdminAuthorHandler object
func NewAdminAuthorHandler(cmntService admin.AuthorService) *AdminAuthorHandler {
	return &AdminAuthorHandler{AuthorService: cmntService}
}

// GetAuthors handles GET /v1/admin/Authors request
func (ach *AdminAuthorHandler) GetAuthors(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	Authors, errs := ach.AuthorService.Authors()

	if errs != nil {
		fmt.Println(errs)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(Authors, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

	return

}

// GetSingleAuthor handles GET /v1/admin/Authors/:id request
func (ach *AdminAuthorHandler) GetSingleAuthor(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")

		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	Author, errs := ach.AuthorService.Author(id)

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(Author, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// PostAuthor handles POST /v1/admin/Authors request
func (ach *AdminAuthorHandler) PostAuthor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	Author := &entity.Author{}

	err := json.Unmarshal(body, Author)

	if err != nil {

		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	errs := ach.AuthorService.StoreAuthor(*Author)

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/v1/admin/Authors/%d", Author.AuthorID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}

// PutAuthor handles PUT /v1/admin/Authors/:id request
func (ach *AdminAuthorHandler) PutAuthor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	Author, errs := ach.AuthorService.Author(id)

	if errs != nil {

		fmt.Println(errs)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	json.Unmarshal(body, &Author)

	errs = ach.AuthorService.UpdateAuthor(Author)

	if errs != nil {

		fmt.Println(errs)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(Author, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// DeleteAuthor handles DELETE /v1/admin/Authors/:id request
func (ach *AdminAuthorHandler) DeleteAuthor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {

		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	errs := ach.AuthorService.DeleteAuthor(id)

	if errs != nil {

		fmt.Println(errs)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
