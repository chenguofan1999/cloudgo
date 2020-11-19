package cloudgo

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	//	"github.com/unrolled/render"
)

// NewServer returns a server
func NewServer() *negroni.Negroni {
	router := mux.NewRouter()
	router.HandleFunc("/", defaultHandler)
	router.HandleFunc("/hello/{name}", helloHandler)

	n := negroni.Classic()
	n.UseHandler(router)
	return n
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Some descriptions here")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintln(w, "Hello, ", vars["name"])
}
