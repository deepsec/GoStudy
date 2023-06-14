package main

import (
	"fmt"
	"net/http"
)

var PORT = ":1443"

func Default(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is an example HTTPS server.\n")
}

func main() {
	http.HandleFunc("/", Default)
	fmt.Println("Listening on port number", PORT)

	err := http.ListenAndServeTLS(PORT, "../../certs/server.crt", "../../certs/server.key", nil)
	if err != nil {
		fmt.Println("LisntenAndServeTLS: ", err)
		return
	}
}
