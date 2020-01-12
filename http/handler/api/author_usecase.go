package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/Endalk/Online-Book-Shoping/admin/service"
)

//AuthorUseCaseHandler -
type AuthorUseCaseHandler struct {
	comp service.AuthorServiceImpl
}

//NewAuthorUseCaseHander -
func NewAuthorUseCaseHander(cp service.AuthorServiceImpl) *AuthorUseCaseHandler {
	return &AuthorUseCaseHandler{comp: cp}
}

//Login -
func (cuh *AuthorUseCaseHandler) Login(w http.ResponseWriter, r *http.Request) {
	//do the authentication here let's assume passed
	if r.Method == http.MethodPost {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["user"] = r.FormValue("username")
		claims["exp"] = time.Now().Add(time.Minute * 100).Unix()
		tokenString, err := token.SignedString([]byte("please-please"))
		if err != nil {
			panic(err)
		}
		cookie := http.Cookie{Name: "author-token", Value: tokenString}
		r.AddCookie(&cookie)
		// fmt.Println("You are now logged in ")
		json.NewEncoder(w).Encode(cookie)
	}
}

func (cuh *AuthorUseCaseHandler) Secret(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Congrats you passed the test")
}
