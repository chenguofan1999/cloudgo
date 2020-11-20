package cloudgo

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"gopl.io/ch5/links"
)

// NewServer returns a server
func NewServer() *negroni.Negroni {
	router := mux.NewRouter()
	router.HandleFunc("/hello/{name}", helloHandler)
	router.HandleFunc("/GPA/{name}", gpaHandler)
	router.HandleFunc("/crawl/", crawl)
	router.PathPrefix("/").HandlerFunc(defaultHandler)

	n := negroni.Classic()
	n.UseHandler(router)
	return n
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Fprintln(w, "1. /hello/[your name]")
	fmt.Fprintln(w, "2. /GPA/[your name]")
	fmt.Fprintln(w, "3. /crawl/?url=[your URL]")

	for k, v := range r.Form {
		fmt.Println(k, " : ", strings.Join(v, ""))
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintln(w, "Hello, ", vars["name"])
}

func crawl(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	s, _ := links.Extract(r.Form["url"][0])
	fmt.Fprintln(w, "Found links:\n", strings.Join(s, "\n"))
}

func gpaHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	n := 2 + 2*rand.Float32()
	fmt.Fprintln(w, "GPA of ", vars["name"], " = ", n)
}
