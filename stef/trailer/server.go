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

//TrailerHandler add trailer headers.
//output see "out.txt", curl -i --raw , explorer's dev tool cannot display trailers
func TrailerHandler(w http.ResponseWriter, req *http.Request) {
	// Before any call to WriteHeader or Write, declare
	// the trailers you will set during the HTTP
	// response. These three headers are actually sent in
	// the trailer.
	w.Header().Set("Trailer", "AtEnd1, AtEnd2")
	w.Header().Add("Trailer", "Expires")
	w.Header().Set(http.TrailerPrefix+"atfirst", "dd")

	w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Header().Set("AtEnd1", "value 1")

	io.WriteString(w, "This HTTP response has both headers before this text and trailers at the end.\n")
	if flusher, ok := w.(http.Flusher); ok {
		flusher.Flush()
	}
	time.Sleep(1 * time.Second)
	io.WriteString(w, "just another string\n")
	w.Header().Set("AtEnd2", "value 2")
	w.Header().Set("Expires", time.Now().Add(120*time.Second).UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")) // These will appear as trailers.

	w.Header().Set(http.TrailerPrefix+"atlast", "dd")
	w.Header().Set("lost", "error")//not work
}
