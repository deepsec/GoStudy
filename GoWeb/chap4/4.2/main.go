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
		for k, v := range r.Form {
			fmt.Printf("%s ==> %v\n", k, v)
			fmt.Fprintf(w, "%s ==> %v\n", k, v)
		}
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
