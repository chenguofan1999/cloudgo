package cloudgo

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

// NewServer returns a server
func newTestServer() *negroni.Negroni {
	router := mux.NewRouter()
	router.HandleFunc("/hello/{name}", helloHandler)
	router.HandleFunc("/GPA/{name}", gpaHandler)
	router.HandleFunc("/crawl/", crawl)
	router.HandleFunc("/calc/{mode}/", calc)
	router.PathPrefix("/").HandlerFunc(defaultHandler)

	n := negroni.Classic()
	n.UseHandler(router)
	return n
}

func newServer() *negroni.Negroni {
	router := mux.NewRouter()
	router.Handle("/static/", http.FileServer(http.Dir("")))

	n := negroni.Classic()
	n.UseHandler(router)
	return n
}
