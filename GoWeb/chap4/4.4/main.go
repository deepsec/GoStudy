package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type MYTOKEN struct {
	pre string
	cur string
}

var myToken MYTOKEN

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		now := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(now, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("login.tmpl")
		t.Execute(w, token)

		myToken.pre = myToken.cur
		myToken.cur = token
		fmt.Printf("pre: %s,  cur: %s\n", myToken.pre, myToken.cur)
	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		fmt.Println("client token:", token)
		fmt.Printf("pre: %s,  cur: %s\n", myToken.pre, myToken.cur)
		if token != "" {
			if token != myToken.pre {
				fmt.Fprintf(w, "Invalid token, error\n")
				return
			} else {
				for k, v := range r.Form {
					fmt.Printf("%s ==> %v\n", k, v)
					fmt.Fprintf(w, "%s ==> %v\n", k, v)
				}
			}
		} else {
			fmt.Fprintf(w, "No token, error\n")
			return
		}

	}
}

func main() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
