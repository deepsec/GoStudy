package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("Form:", r.Form)
	fmt.Println("Scheme:", r.URL.Scheme)
	fmt.Println("Path:", r.URL.Path)
	fmt.Println(r.Form["url_long"])
	fmt.Println(r.Form["username"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("value: ", strings.Join(v, " "))
	}

	fmt.Fprintf(w, "Hello 贺新民!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.tmpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()

		fmt.Println("Form:", r.Form)
		fmt.Println("Scheme:", r.URL.Scheme)
		fmt.Println("Path:", r.URL.Path)
		fmt.Println(r.Form["url_long"])
		fmt.Println(r.Form["username"])
		fmt.Fprintf(w, "method: ", r.Method)
		fmt.Fprintf(w, "username: %s\n", r.Form["username"])
		fmt.Fprintf(w, "password: %s\n", r.Form["password"])
		fmt.Fprintf(w, "files: %s\n", r.Form["files"])
		fmt.Fprintf(w, "files: Content\n", r.Body)
	}
}

func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
