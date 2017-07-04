package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"net/http"
	"goblet"
)

// cookie handling

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func getUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}

func setSession(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

// login handler

func loginHandler(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	email := request.FormValue("email")
	pass := request.FormValue("password")
	redirectTarget := "/"
	err := goblet.AddUser(name, email, 1)
	if err != nil {
		switch err.(type) {
		case *goblet.UsernameBusyError:
			redirectTarget = "/usernamebusy"
		case *goblet.EmailBusyError:
			redirectTarget = "/emailbusy"
		default:
			panic(err)
		}
	} else {
		setSession(name, response)
		redirectTarget = "/internal"
	}
	fmt.Printf("User '%s' has logged in with password: '%s'\n", name, pass)
	http.Redirect(response, request, redirectTarget, 302)
}

// logout handler

func logoutHandler(response http.ResponseWriter, request *http.Request) {
	clearSession(response)
	http.Redirect(response, request, "/", 302)
}

// index page

const indexPage = `
<h1>Create new user</h1>
<form method="post" action="/login">
    <label for="name">User name</label>
    <input type="text" id="name" name="name">
    <label for="email">E-mail</label>
    <input type="text" id="email" name="email">
    <label for="password">Password</label>
    <input type="password" id="password" name="password">
    <button type="submit">Login</button>
</form>
`

func indexPageHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, indexPage)
}

const userNameTakenPage = `
<h1>That username is taken!</h1>
`

func userNameTakenPageHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, userNameTakenPage)
}
const emailTakenPage = `
<h1>That email has already been signed up!</h1>
`

func emailTakenPageHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, emailTakenPage)
}

// internal page

const internalPage = `
<h1>Internal</h1>
<hr>
<small>User: %s</small>
<form method="post" action="/logout">
    <button type="submit">Logout</button>
</form>
`

func internalPageHandler(response http.ResponseWriter, request *http.Request) {
	userName := getUserName(request)
	if userName != "" {
		fmt.Fprintf(response, internalPage, userName)
	} else {
		http.Redirect(response, request, "/", 302)
	}
}

// server main method

var router = mux.NewRouter()

func main() {
	goblet.Connect()

	router.HandleFunc("/", indexPageHandler)
	router.HandleFunc("/internal", internalPageHandler)

	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/logout", logoutHandler).Methods("POST")

	router.HandleFunc("/usernamebusy", userNameTakenPageHandler)
	router.HandleFunc("/emailbusy", emailTakenPageHandler)

	http.Handle("/", router)
	http.ListenAndServe(":9090", nil)
}
