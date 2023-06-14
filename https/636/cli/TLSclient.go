package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
		return
	}

	URL := os.Args[1]

	caCert, err := ioutil.ReadFile("../../certs/server.crt")
	if err != nil {
		fmt.Println(err)
		return
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	cert, err := tls.LoadX509KeyPair("../../certs/client.crt", "../../certs/client.key")
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:            caCertPool,
				Certificates:       []tls.Certificate{cert},
				InsecureSkipVerify: true,
			},
		},
	}
	response, err := client.Get(URL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	s := strings.TrimSpace(string(content))
	fmt.Printf("%v\n", response.Status)
	fmt.Println(s)
}
