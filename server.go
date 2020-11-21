package cloudgo

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

// NewTestServer returns a test server
func NewTestServer() *negroni.Negroni {
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

// NewServer returns a usable server
func NewServer() *negroni.Negroni {
	router := mux.NewRouter()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	router.HandleFunc("/js", jsonHandler)

	// 不要用"/login/", 否则 post 时不会触发该 handler
	router.HandleFunc("/login", loginHandler)

	//router.PathPrefix("/").HandlerFunc(mainPageHandler)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("")))

	n := negroni.Classic()
	n.UseHandler(router)
	return n
}
