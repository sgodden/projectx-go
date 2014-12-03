package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"projectx/resources"
)

func main() {
		
	r := mux.NewRouter()
	
	redirectToIndex := http.RedirectHandler("/index.htm", 302)
	
	r.Handle("", redirectToIndex)
	r.Handle("/", redirectToIndex)
	
	r.PathPrefix("/index.htm").HandlerFunc(serveIndex)
	
	fileServer := http.FileServer(http.Dir("static"))
	r.PathPrefix("/app/").Handler(fileServer)
	
	s := r.PathPrefix("/rest").Subrouter()
	resources.Configure(s)
	
	fmt.Println("Listening on port 8080...")

	http.ListenAndServe(":8080", r)
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.htm")
}
