package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var (
	defaultContentType = flag.String("content-type", "application/json", "Default content type to use if not specified in the url")
	webServerPort      = flag.String("port", ":8080", "Port to listen on")
)

func init() {
	flag.Parse()
}

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)
	r.HandleFunc("/api/{content-type}/", contentTypeProxy).Methods("GET")
	r.HandleFunc("/api/", contentTypeProxy).Methods("GET")
	fmt.Println("Listening on port", webServerPort)
	log.Fatal(http.ListenAndServe(*webServerPort, r))
}

func contentTypeProxy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queryString := r.URL.Query()
	urlFromQuery := queryString.Get("url")
	if urlFromQuery == "" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "url parameter is required")
		return
	}

	contentType := vars["content-type"]
	if contentType == "" {
		// use default content type from flag
		contentType = *defaultContentType
	}

	contentType = strings.Replace(contentType, "-", "/", -1)
	fetchResponse, err := http.Get(urlFromQuery)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "error fetching url")
		return
	}
	defer fetchResponse.Body.Close()
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
	io.Copy(w, fetchResponse.Body)
	// Rewriting content-type header for <url> from <old-content-type> to <new-content-type>
	log.Println("Rewriting content-type header for", urlFromQuery, "from", fetchResponse.Header.Get("Content-Type"), "to", contentType)
}
