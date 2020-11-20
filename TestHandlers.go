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

var testPageHTML = `<!DOCTYPE html>
<html>

<head>
    <title></title>
</head>

<body style="color: rgb(0, 0, 0); background-color: rgb(240, 240, 240);">
    <p style="text-align: center;"><span style="font-family: Impact, Charcoal, sans-serif; font-size: 28px; color: rgb(40, 50, 78);">CloudGo测试</span></p>
    <h3><strong><span style="color: rgb(71, 85, 119); font-size: 24px;">云服务(员)</span></strong></h3>
    <div><span style="font-family: Verdana, Geneva, sans-serif; color: rgb(40, 50, 78); font-size: 18px;">./hello/[your name]</span></div>
    <p><span style="font-family: Tahoma, Geneva, sans-serif; color: rgb(204, 204, 204); font-size: 18px;"><a href="./hello/Dave">example: Dave</a></span></p>
    <p><br></p>
    <h3><strong><span style="color: rgb(71, 85, 119); font-size: 24px;">云计算(器)</span></strong></h3>
    <div><span style="font-family: Verdana, Geneva, sans-serif; color: rgb(40, 50, 78); font-size: 18px;">./calc/{add/sub/mul/div}/?a=[?]&amp;b=[?]</span></div>
    <p><span style="font-family: Verdana, Geneva, sans-serif; color: rgb(40, 50, 78); font-size: 18px;"><span style="font-family: Verdana, Geneva, sans-serif; color: rgb(40, 50, 78); font-size: 18px;"><span style="font-family: Tahoma, Geneva, sans-serif; color: rgb(204, 204, 204); font-size: 18px;"><a href="./calc/mul/?a=4&b=6">example:4*6</a></span></span></span></p>
    <p><br></p>
    <h3><strong><span style="color: rgb(71, 85, 119); font-size: 24px;">查询GPA😅</span></strong></h3>
    <div><span style="font-size: 18px;"><span style="font-family: Verdana, Geneva, sans-serif; color: rgb(40, 50, 78);">./GPA/[your name]</span></span></div>
    <p><span style="font-family: Tahoma, Geneva, sans-serif; color: rgb(204, 204, 204); font-size: 18px;"><a href="./GPA/18342000">example:18342000</a></span></p>
    <p><br></p>
    <h3><strong><span style="color: rgb(71, 85, 119); font-size: 24px;">爬虫功能</span></strong></h3>
    <div><span style="font-family: Verdana, Geneva, sans-serif; color: rgb(40, 50, 78); font-size: 18px;">./crawl/?url=[your URL]</span></div>
    <p><span style="font-family: Verdana, Geneva, sans-serif; color: rgb(40, 50, 78); font-size: 18px;"><span style="font-family: Verdana, Geneva, sans-serif; color: rgb(40, 50, 78); font-size: 18px;"><span style="font-family: Tahoma, Geneva, sans-serif; color: rgb(204, 204, 204); font-size: 18px;"><a href="./crawl/?url=https%3A//pmlpml.gitee.io/service-computing/post/index-2020/">example:服务计算课程页面</a></span></span></span></p>
    <h3><br></h3>
</body>

</html>`

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
