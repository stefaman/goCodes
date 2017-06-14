package main

import (
	"fmt"
	"net/http"
	"time"
)

func sseHandler(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	header.Add("Content-Type", "text/event-stream")
	for count := 0;;count++ {
		fmt.Fprintln(w, "event: ping")
		fmt.Fprintf(w, "id: %d\n", count)
		fmt.Fprintf(w, "data:{\"time\":\"%s\"}\n\n", time.Now())
		if flusher, ok := w.(http.Flusher); ok {
			flusher.Flush()
		}
		time.Sleep(time.Second)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	path := "index.html"
	http.ServeFile(w, r, path)
}

func main() {
	http.HandleFunc("/event", index)
	http.HandleFunc("/sse", sseHandler)
	http.ListenAndServe(":8080", nil)
}
