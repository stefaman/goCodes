package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path, r.Method)
	printCookie(r)
	if tag, ok := r.Header["Etag"]; ok {
		log.Println("Etag", tag)
	}
	header := w.Header()
	header.Add("Set-Cookie", "key=value; HttpOnly; Secure;SameSite=Strict;Max-Age=120;Path="+r.URL.Path)
	header.Add("ETag", `"key=value"`) //use Etag as cookie
	path := "index.html"
	http.ServeFile(w, r, path)

}

func printCookie(r *http.Request) {
	cks := r.Cookies()
	for _, c := range cks {
		log.Println(c.String())
	}
}
func hello(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path, r.Method)
	fmt.Println(w, r.URL.Path, r.Method)
}

//go run $GOROOT/src/crypto/tls/generate_cert.go -host="localhost:10.0.0.8" -start-date="Jan 1 00:00:00 2017" -ca
func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil))
}
