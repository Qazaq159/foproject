package main

import (
	"fmt"
	"net/http"
)

type User struct {
	Username string
	Password string
	IsOnline bool
}

func main() {
	//http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	//	http.ServeFile(w, r, "static/main.html")
	//})
	//
	//http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
	//	http.ServeFile(w, r, "static/about.html")
	//})
	//
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	http.ServeFile(w, r, "static/user.html")
	//})

	var currentUser = User{}
	currentUser.IsOnline = false

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if !currentUser.IsOnline {
			fmt.Fprintf(writer, "User is not logged in, please log in.\n"+
				" Go /register if you do not have account"+
				"\nelse go /login")
			return
		}
		http.ServeFile(writer, request, "static/main.html")
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/register.html")
	})

	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		password := r.FormValue("password")
		currentUser.Username = username
		currentUser.Password = password

		fmt.Fprintf(w, "You are registered, please log in")
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/user.html")
	})

	http.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == currentUser.Username && password == currentUser.Password {
			currentUser.IsOnline = true
			fmt.Fprintf(w, "You are logged in, go main page /")
			return
		}

		fmt.Fprintf(w, "Incorrect password or username, try again")
	})

	fmt.Println("Server is listening ...")
	err := http.ListenAndServe(":8181", nil)

	if err != nil {
		return
	}
}
