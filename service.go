package cloudgo

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
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
	router.HandleFunc("/calc/{mode}/", calc)
	router.PathPrefix("/").HandlerFunc(defaultHandler)

	n := negroni.Classic()
	n.UseHandler(router)
	return n
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Fprintln(w, `<p style="text-align: center;"><span style="font-family: Impact, Charcoal, sans-serif; font-size: 28px; color: rgb(40, 50, 78);">CloudGo</span></p>
	<h3><strong><span style="color: rgb(71, 85, 119); font-size: 24px;">云服务(员)</span></strong></h3>
	<div><span style="font-family: Verdana, Geneva, sans-serif; color: rgb(40, 50, 78); font-size: 18px;">./hello/[your name]</span></div>
	<h3><strong><span style="color: rgb(71, 85, 119); font-size: 24px;">查询GPA功能😅</span></strong></h3>
	<div><span style="font-size: 18px;"><span style="font-family: Verdana, Geneva, sans-serif; color: rgb(40, 50, 78);">./GPA/[your name]</span></span></div>
	<h3><strong><span style="color: rgb(71, 85, 119); font-size: 24px;">爬虫功能</span></strong></h3>
	<div><span style="font-family: Verdana, Geneva, sans-serif; color: rgb(40, 50, 78); font-size: 18px;">./crawl/?url=[your URL]</span></div>
	<h3><strong><span style="color: rgb(71, 85, 119); font-size: 24px;">云计算(器)</span></strong></h3>
	<div><span style="font-family: Verdana, Geneva, sans-serif; color: rgb(40, 50, 78); font-size: 18px;">./calc/{add/sub/mul/div}/?a=[?]&amp;b=[?]</span></div>`)

	for k, v := range r.Form {
		fmt.Println(k, " : ", strings.Join(v, ""))
	}
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
