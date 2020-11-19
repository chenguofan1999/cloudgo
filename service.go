package cloudgo

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	//	"github.com/unrolled/render"
)

// NewServer returns a server
func NewServer() *negroni.Negroni {
	router := mux.NewRouter()
	router.HandleFunc("/hello/{name}", helloHandler)
	router.HandleFunc("/GPA/{name}", gpaHandler)
	router.PathPrefix("/").HandlerFunc(defaultHandler)
	n := negroni.Classic()
	n.UseHandler(router)
	return n
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Whatever")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintln(w, "Hello, ", vars["name"])
}

func gpaHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	n := 2 + 2*rand.Float32()
	fmt.Fprintln(w, "GPA of ", vars["name"], " = ", n)
}
