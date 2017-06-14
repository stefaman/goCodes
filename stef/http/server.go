package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/file/", FileHandler)
	http.HandleFunc("/trailer", TrailerHandler)
	http.Handle("/content/", http.TimeoutHandler(http.HandlerFunc(ContentHandler), 100*time.Second, "time out..."))
	http.ListenAndServe(":8080", nil)

}

func FileHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	path := strings.TrimPrefix(r.URL.Path, "/file")
	http.ServeFile(w, r, path)
}

func ContentHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/content")
	file, err := os.Open(path)
	fileInfo, err := os.Stat(path)
	if err != nil {
		panic(http.ErrAbortHandler)
	}

	mtime := fileInfo.ModTime()
	http.ServeContent(w, r, path, mtime, file)
}

func TrailerHandler(w http.ResponseWriter, req *http.Request) {
	// Before any call to WriteHeader or Write, declare
	// the trailers you will set during the HTTP
	// response. These three headers are actually sent in
	// the trailer.
	// w.Header().Set("Trailer", "AtEnd1, AtEnd2")
	// w.Header().Add("Trailer", "AtEnd3")

	w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)

	w.Header().Set("AtEnd1", "value 1")
	io.WriteString(w, "This HTTP response has both headers before this text and trailers at the end.\n")
	if flusher, ok := w.(http.Flusher); ok {
		log.Println("flush")
		flusher.Flush()
	}
	time.Sleep(5 * time.Second)
	io.WriteString(w, "AtEnd1: valu1\n")
	w.Header().Set("AtEnd2", "value 2")
	w.Header().Set("AtEnd3", "value 3") // These will appear as trailers.

	w.Header().Set(http.TrailerPrefix+"kk", "dd")
}
