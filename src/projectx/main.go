package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"projectx/resources"
)

func main() {
		
	r := mux.NewRouter()
	
	// redirect initial requests to the index
	redirectToIndex := http.RedirectHandler("/index.htm", 302)
	r.Handle("", redirectToIndex)
	r.Handle("/", redirectToIndex)
	
	// for any requests starting with index.htm, serve index.htm
	r.PathPrefix("/index.htm").HandlerFunc(serveIndex)
	
	// serve all static assets for the web app's JS, CSS etc
	fileServer := http.FileServer(http.Dir("static"))
	r.PathPrefix("/app/").Handler(fileServer)
	
	// add the REST services in under the /rest prefix
	s := r.PathPrefix("/rest").Subrouter()
	resources.Configure(s)
	
	// fire 'er up
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", r)
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.htm")
}
