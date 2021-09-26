package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	resp, err := http.Get("https://golang.org/")

	if err != nil {
		fmt.Println(err)
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	print(string(bytes))

}
