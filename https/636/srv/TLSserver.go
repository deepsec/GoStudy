package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

var PORT = ":1443"

type handler struct{}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!\n"))
}

func main() {
	caCert, err := ioutil.ReadFile("../../certs/client.crt")
	if err != nil {
		fmt.Println(err)
		return
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	cfg := &tls.Config{
		ClientCAs:  caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}

	srv := &http.Server{
		Addr:      PORT,
		Handler:   handler{},
		TLSConfig: cfg,
	}
	fmt.Println("Listening on port number", PORT)
	fmt.Println(srv.ListenAndServeTLS("../../certs/server.crt", "../../certs/server.key"))
}
