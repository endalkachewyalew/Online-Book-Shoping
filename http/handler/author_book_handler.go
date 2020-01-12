package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/Endalk/Online-Book-Shoping/author"
		"github.com/Endalk/Online-Book-Shoping/entity"
)

// AuthorBookHandler handles Author handler Author requests
type AuthorBookHandler struct {
	tmpl        *template.Template
	BookSrv author.BookService
}

// NewAuthorBookHandler initializes and returns new AuthorBookHandler
func NewAuthorBookHandler(T *template.Template, CS author.BookService) *AuthorBookHandler {
	return &AuthorBookHandler{tmpl: T, BookSrv: CS}
}

// AuthorBooks handle requests on route /Author/Books
func (ach *AuthorBookHandler) AuthorBooks(w http.ResponseWriter, r *http.Request) {
	Books, err := ach.BookSrv.Books()
	if err != nil {
		panic(err)
		fmt.Println(err)
	}
	ach.tmpl.ExecuteTemplate(w, "author.book.layout", Books)
}

// AuthorBooksNew hanlde requests on route /Author/Books/new
func (ach *AuthorBookHandler) AuthorBooksNew(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		ctg := entity.Book{}
		ctg.Name = r.FormValue("name")
		// own, _ := strconv.Atoi(r.FormValue("owner"))
		// ctg.Owner = own
		ppd, _ := strconv.ParseFloat(r.FormValue("priceperday"), 10)
		ctg.PricePerDay = ppd
		ondiscount, _ := strconv.ParseBool(r.FormValue("ondiscount"))
		ctg.OnDiscount = ondiscount
		discout, _ := strconv.ParseFloat(r.FormValue("discount"), 10)
		ctg.Discount = float32(discout)
		onsal, _ := strconv.ParseBool(r.FormValue("onsale"))
		ctg.OnSale = onsal
		mf, fh, err := r.FormFile("catimg")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		ctg.ImagePath = fh.Filename

		writeFile(&mf, fh.Filename)

		err = ach.BookSrv.AddBook(ctg)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/author/book", http.StatusSeeOther)

	} else {

		ach.tmpl.ExecuteTemplate(w, "author.book.new.layout", nil)

	}
}

// AuthorBooksUpdate handle requests on /Author/Books/update
func (ach *AuthorBookHandler) AuthorBooksUpdate(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		cat, err := ach.BookSrv.Book(id)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		ach.tmpl.ExecuteTemplate(w, "author.book.update.layout", cat)

	} else if r.Method == http.MethodPost {

		ctg := entity.Book{}
		ctg.Name = r.FormValue("name")
		own, _ := strconv.Atoi(r.FormValue("owner"))
		ctg.Owner = own
		ppd, _ := strconv.ParseFloat(r.FormValue("priceperday"), 10)
		ctg.PricePerDay = ppd
		ondiscount, _ := strconv.ParseBool(r.FormValue("ondiscount"))
		ctg.OnDiscount = ondiscount
		discout, _ := strconv.ParseFloat(r.FormValue("discount"), 10)
		ctg.Discount = float32(discout)
		onsal, _ := strconv.ParseBool(r.FormValue("onsale"))
		ctg.OnSale = onsal
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

		err = ach.BookSrv.UpdateBook(ctg)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/book", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/book", http.StatusSeeOther)
	}

}

// AuthorBooksDelete handle requests on route /Author/categories/delete
func (ach *AuthorBookHandler) AuthorBooksDelete(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		err = ach.BookSrv.DeleteBook(id)

		if err != nil {
			panic(err)
		}

	}

	http.Redirect(w, r, "/author/book", http.StatusSeeOther)
}
