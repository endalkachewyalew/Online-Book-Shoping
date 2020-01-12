package main

import (
	"database/sql"
	//"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/Endalk/Online-Book-Shoping/admin/repository"
	"github.com/Endalk/Online-Book-Shoping/admin/service"
	comprep "github.com/Endalk/Online-Book-Shoping/author/repository"
	compser "github.com/Endalk/Online-Book-Shoping/author/service"
	handlers "github.com/Endalk/Online-Book-Shoping/http/handler"
	"github.com/Endalk/Online-Book-Shoping/http/handler/api"
	_ "github.com/lib/pq"
)

var templ = template.Must(template.ParseGlob("../ui/templates/*"))

func index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	templ.ExecuteTemplate(w, "index.layout", nil)
}
func login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	templ.ExecuteTemplate(w, "login.layout", nil)
}
func loginAs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	templ.ExecuteTemplate(w, "loginAsAuthor.layout", nil)
}
func admin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	templ.ExecuteTemplate(w, "admin.layout", nil)
}
func userr(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	templ.ExecuteTemplate(w, "user.layout", nil)
}
func company(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	templ.ExecuteTemplate(w, "company.layout", nil)
}

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "1234"
	dbname   = "constructiondb"
)

func main() {

	dbconn, err := sql.Open("postgres", "postgres://postgres:password1234@localhost/onlinebookshopingdb?sslmode=disable")


	

	if err != nil {
		panic(err)
	} //this i
	defer dbconn.Close()

	if err := dbconn.Ping(); err != nil {
		panic(err)
	}

	//admin
	AdminRepo := repository.NewAdminRepositoryImpl(dbconn)
	AdminServ := service.NewAdminServiceImpl(AdminRepo)
	adminAdminsHandler := handlers.NewAdminAdminHandler(templ, AdminServ)
	apiAdminAdminsHandler := api.NewAdminAdminsHandler(AdminServ)

	//company
	AuthorRepo := repository.NewAuthorRepositoryImpl(dbconn)
	AuthorServ := service.NewAuthorServiceImpl(AuthorRepo)
	adminAuthorsHandler := handlers.NewAdminAuthorHandler(templ, AuthorServ)
	apiAdminAuthorsHandler := api.NewAdminAuthorHandler(AuthorServ)
	router := httprouter.New()
	//User
	UserRepo := repository.NewUserRepositoryImpl(dbconn)
	UserServ := service.NewUserServiceImpl(UserRepo)
	adminUsersHandler := handlers.NewAdminUserHandler(templ, UserServ)
	apiAdminUsersHandler := api.NewAdminUserHandler(UserServ)

	bookRepo := comprep.NewBookRepository(dbconn)
	ser := compser.NewBookService(bookRepo)
	hand := api.NewAuthorBookHandler(ser)

	bookHandle := handlers.NewAuthorBookHandler(templ, ser)
	// serv := api.NewAuthorBookHandler(bookSer)
	// ap := api.NewAuthorUseCaseHander(*AuthorServ)
	CommentRepo := repository.NewCommentRepositoryImpl(dbconn)
	CommentServ := service.NewCommentServiceImpl(CommentRepo)
	adminCommentsHandler := handlers.NewAdminCommentHandler(templ, CommentServ)

	fs := http.FileServer(http.Dir("../ui/assets"))
	router.ServeFiles("/assets/*filepath", http.Dir("../ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	router.GET("/", index)
	router.GET("/login", login)
	router.GET("/signinAuthor", loginAs)
	router.GET("/admin", admin)
	router.GET("/user", userr)
	router.GET("/company", company)

	//handle admin
	// router := httprouter.New()

	router.GET("/admin/admins", adminAdminsHandler.AdminAdmins)
	router.POST("/admin/admins/new", adminAdminsHandler.AdminAdminsNew)
	router.POST("/admin/admins/update", adminAdminsHandler.AdminAdminsUpdate)
	router.GET("/admin/admins/update", adminAdminsHandler.AdminAdminsUpdate)
	router.GET("/admin/admins/delete", adminAdminsHandler.AdminAdminsDelete)
	//handle company
	router.GET("/admin/company", adminAuthorsHandler.AdminAuthors)
	router.POST("/admin/company/new", adminAuthorsHandler.AdminAuthorsNew)
	router.POST("/admin/company/update", adminAuthorsHandler.AdminAuthorsNew)
	router.GET("/admin/company/update", adminAuthorsHandler.AdminAuthorsUpdate)
	router.GET("/admin/company/delete", adminAuthorsHandler.AdminAuthorsDelete)
	//handle user
	router.GET("/admin/user", adminUsersHandler.AdminUsers)
	router.POST("/admin/user/new", adminUsersHandler.AdminUsersNew)
	router.GET("/admin/user/new", adminUsersHandler.AdminUsersNew)
	router.PUT("/admin/user/update", adminUsersHandler.AdminUsersUpdate)
	router.GET("/admin/users/delete", adminUsersHandler.AdminUsersDelete)
	//handle user
	http.HandleFunc("/admin/comment", adminCommentsHandler.AdminComments)
	http.HandleFunc("/admin/comment/new", adminCommentsHandler.AdminCommentsNew)
	http.HandleFunc("/admin/comment/update", adminCommentsHandler.AdminCommentsUpdate)
	http.HandleFunc("/admin/comment/delete", adminCommentsHandler.AdminCommentsDelete)

	http.HandleFunc("/company/book", bookHandle.AuthorBooks)
	http.HandleFunc("/company/book/new", bookHandle.AuthorBooksNew)
	http.HandleFunc("/company/book/update", bookHandle.AuthorBooksUpdate)
	http.HandleFunc("/company/book/delete", bookHandle.AuthorBooksDelete)
	// http.HandleFunc("/v1/authors/login", ap.Login)
	// http.HandleFunc("/v1/authors/secret", middleware.IsAuthorized(ap.Secret))

	router.GET("/v1/authors/books", hand.Books)
	router.GET("/v1/authors/books/:book_id", hand.Book)
	router.PUT("/v1/authors/books/:id", hand.UpdateBook)
	router.DELETE("/v1/authors/books/delete/:book_id", hand.DeleteBook)
	router.POST("/v1/authors/books", hand.StoreBook)
	//handle company api
	router.GET("/v1/admin/company/:id", apiAdminAuthorsHandler.GetSingleAuthor)
	router.GET("/v1/admin/company", apiAdminAuthorsHandler.GetAuthors)
	router.PUT("/v1/admin/company/:id", apiAdminAuthorsHandler.PutAuthor)
	router.POST("/v1/admin/company", apiAdminAuthorsHandler.PostAuthor)
	router.DELETE("/v1/admin/company/:id", apiAdminAuthorsHandler.DeleteAuthor)
	//handle user api
	router.GET("/v1/admin/user/:username", apiAdminUsersHandler.GetSingleUser)
	router.GET("/v1/admin/user", apiAdminUsersHandler.GetUsers)
	router.PUT("/v1/admin/user/:username", apiAdminUsersHandler.PutUser)
	router.POST("/v1/admin/user", apiAdminUsersHandler.PostUser)
	router.DELETE("/v1/admin/user/:username", apiAdminUsersHandler.DeleteUser)
	//handle Admin api
	router.GET("/v1/admin/admins/:username", apiAdminAdminsHandler.GetSingleAdmin)
	router.GET("/v1/admin/admins", apiAdminAdminsHandler.GetAdmins)
	router.PUT("/v1/admin/admins/:username", apiAdminAdminsHandler.PutAdmin)
	router.POST("/v1/admin/admins", apiAdminAdminsHandler.PostAdmin)
	router.DELETE("/v1/admin/admins/:username", apiAdminAdminsHandler.DeleteAdmin)
	http.ListenAndServe(":8080", nil)
}
