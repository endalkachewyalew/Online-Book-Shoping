package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/Endalk/Online-Book-Shoping/author"
	"github.com/Endalk/Online-Book-Shoping/entity"
)

//AuthorBookHandler -
type AuthorBookHandler struct {
	books author.BookService
}

//NewAuthorBookHandler -
func NewAuthorBookHandler(mat author.BookService) *AuthorBookHandler {
	return &AuthorBookHandler{books: mat}
}

//Books - GET /
func (ch *AuthorBookHandler) Books(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	books, err := ch.books.Books()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//DeleteBook -
func (ch *AuthorBookHandler) DeleteBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("book_id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	err = ch.books.DeleteBook(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	http.Redirect(w, r, "/v1/authors/books", http.StatusSeeOther)
}

//StoreBook -
func (ch *AuthorBookHandler) StoreBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	book := &entity.Book{}

	err := json.Unmarshal(body, book)

	if err != nil {

		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	fmt.Println(book)
	err = ch.books.AddBook(*book)

	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	p := fmt.Sprintf("/v1/authors/books")
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}

//Book -
func (ch *AuthorBookHandler) Book(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("book_id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	book, err := ch.books.Book(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

//UpdateBook -
func (ch *AuthorBookHandler) UpdateBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	book, errs := ch.books.Book(id)

	if errs != nil {

		fmt.Println(errs)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	json.Unmarshal(body, &book)

	errs = ch.books.UpdateBook(book)

	if errs != nil {

		fmt.Println(errs)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(book, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
