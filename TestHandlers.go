package cloudgo

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"gopl.io/ch5/links"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Fprintln(w, testPageHTML)

	for k, v := range r.Form {
		fmt.Println(k, " : ", strings.Join(v, ""))
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintln(w, "Hello,", vars["name"])
}

func gpaHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	n := 2 + 3*rand.Float32()
	fmt.Fprintf(w, "GPA of %s is %.2f\n", vars["name"], n)
}

func crawl(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	s, _ := links.Extract(r.Form["url"][0])
	fmt.Fprintln(w, "Found links:\n", strings.Join(s, "\n"))
}

func calc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	r.ParseForm()
	a, _ := strconv.Atoi(r.Form["a"][0])
	b, _ := strconv.Atoi(r.Form["b"][0])
	var c int
	switch vars["mode"] {
	case "add":
		c = a + b
	case "sub":
		c = a - b
	case "mul":
		c = a * b
	case "div":
		c = a / b
	}
	fmt.Fprintln(w, c)
}
