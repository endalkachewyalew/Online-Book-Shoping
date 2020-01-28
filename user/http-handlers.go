package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func handleSaveUser(w http.ResponseWriter, r *http.Request) {
	var id = 0
	var err error

	r.ParseForm()
	params := r.PostForm
	idStr := params.Get("id")

	if len(idStr) > 0 {
		id, err = strconv.Atoi(idStr)
		if err != nil {
			renderErrorPage(w, err)
			return
		}
	}

	fullname := params.Get("fullname")
	gender := params.Get("gender")
	country := params.Get("country")
	city := params.Get("city")
	email := params.Get("email")
	password := params.Get("password")

	pnumber := params.Get("pnumber")
	var pages int
	if len(pnumber) > 0 {
		pages, err = strconv.Atoi(pnumber)
		if err != nil {
			renderErrorPage(w, err)
			return
		}
	}

	dob := params.Get("dob")
	var publicationDate time.Time

	if len(dob) > 0 {
		publicationDate, err = time.Parse("2006-01-02", dob)
		if err != nil {
			renderErrorPage(w, err)
			return
		}
	}

	if id == 0 {
		_, err = insertUser(fullname, gender, publicationDate, country, city, pages, email, password)
	} else {
		_, err = updateUser(id, fullname, gender, publicationDate, country, city, pages, email, password)
	}

	if err != nil {
		renderErrorPage(w, err)
		return
	}

	http.Redirect(w, r, "/", 302)
}

func renderSuccessPage(w http.ResponseWriter,  r *http.Request) {

	users, err := allUsers()
	buf, err := ioutil.ReadFile("www/signup.html")
	if err != nil {
		renderErrorPage(w, err)
		return
	}

	var page = IndexPageUser{AllUsers: users}
	indexPage := string(buf)
	t := template.Must(template.New("indexPage").Parse(indexPage))
	t.Execute(w, page)


}
func handleLogin(w http.ResponseWriter, r *http.Request){

	var err error

	r.ParseForm()
	params := r.PostForm


	email := params.Get("email")
	//password := params.Get("password")

	
		n, err := getuser(email)
		if err != nil {

			renderErrorPage(w, err)
			return
		} 
		buf, err := ioutil.ReadFile("www/login.html")
	if err != nil {
		renderErrorPage(w, err)
		return
	}

		var page = UserPage{TargetUser: n}
	UserPage := string(buf)
	t := template.Must(template.New("userPage").Parse(UserPage))
	err = t.Execute(w, page)
	if err != nil {
		renderErrorPage(w, err)
		return
	}

	

	http.Redirect(w, r, "/", 302)
}
func handleListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := allUsers()
	if err != nil {
		renderErrorPage(w, err)
		return
	}

	buf, err := ioutil.ReadFile("www/index.html")
	if err != nil {
		renderErrorPage(w, err)
		return
	}

	var page = IndexPageUser{AllUsers: users}
	indexPage := string(buf)
	t := template.Must(template.New("indexPage").Parse(indexPage))
	t.Execute(w, page)
}

func handleViewUserLogin(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	email := params.Get("email")

	var currentUser = User{}
	currentUser.DOB = time.Now()

	if len(email) > 0 {
		_, err := strconv.Atoi(email)
		if err != nil {
			renderErrorPage(w, err)
			return
		}

		currentUser, err = getuser(email)
		if err != nil {
			renderErrorPage(w, err)
			return
		}
	}

	buf, err := ioutil.ReadFile("www/login.html")
	if err != nil {
		renderErrorPage(w, err)
		return
	}

	var page = UserPage{TargetUser: currentUser}
	UserPage := string(buf)
	t := template.Must(template.New("userPage").Parse(UserPage))
	err = t.Execute(w, page)
	if err != nil {
		renderErrorPage(w, err)
		return
	}
}

func handleViewUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	email := params.Get("email")

	var currentUser = User{}
	currentUser.DOB = time.Now()

	if len(email) > 0 {
		_, err := strconv.Atoi(email)
		if err != nil {
			renderErrorPage(w, err)
			return
		}

		currentUser, err = getuser(email)
		if err != nil {
			renderErrorPage(w, err)
			return
		}
	}

	buf, err := ioutil.ReadFile("www/signup.html")
	if err != nil {
		renderErrorPage(w, err)
		return
	}

	var page = UserPage{TargetUser: currentUser}
	UserPage := string(buf)
	t := template.Must(template.New("userPage").Parse(UserPage))
	err = t.Execute(w, page)
	if err != nil {
		renderErrorPage(w, err)
		return
	}
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	email := params.Get("email")

	if len(email) > 0 {
		_, err := strconv.Atoi(email)
		if err != nil {
			renderErrorPage(w, err)
			return
		}

		n, err := removeUser(email)
		if err != nil {
			renderErrorPage(w, err)
			return
		}

		fmt.Printf("Rows removed: %v\n", n)
	}
	http.Redirect(w, r, "/", 302)
}
func renderErrorPage(w http.ResponseWriter, errorMsg error) {
	buf, err := ioutil.ReadFile("www/error.html")
	if err != nil {
		log.Printf("%v\n", err)
		fmt.Fprintf(w, "%v\n", err)
		return
	}

	var page = ErrorPage{ErrorMsg: errorMsg.Error()}
	errorPage := string(buf)
	t := template.Must(template.New("errorPage").Parse(errorPage))
	t.Execute(w, page)
}
