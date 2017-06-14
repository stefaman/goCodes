package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func sseHandler(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	header.Add("Content-Type", "text/event-stream")
	header.Add("Access-Control-Allow-Origin", "http://localhost:8080")
	header.Add("Vary", "Origin")
	header.Add("Access-Control-Allow-Credentials", "true")

	for count := 0; count < 10; count++ {
		// fmt.Fprintln(w, "event: ping")
		fmt.Fprintf(w, "id: %d\n", count)
		fmt.Fprintf(w, "data:{\"time\":\"%s\"}\n\n", time.Now())
		if c, err := r.Cookie("key"); err != nil {
			fmt.Fprintln(w, "data:", c.String())
		}

		if flusher, ok := w.(http.Flusher); ok {
			flusher.Flush()
		}
		time.Sleep(time.Second)
	}

	log.Println(r.URL.Path, r.Method)
}

func index(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	header.Add("Set-Cookie", "key=value; HttpOnly")
	path := "index.html"
	http.ServeFile(w, r, path)

	log.Println(r.URL.Path, r.Method)
}

func printCookie() {

}
func hello(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path, r.Method)
	fmt.Println(w, r.URL.Path, r.Method)
}
func main() {
	http.HandleFunc("/event", index)
	http.HandleFunc("/sse", sseHandler)
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
