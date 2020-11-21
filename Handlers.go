package cloudgo

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/unrolled/render"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		// 此地址相对于 main 程序位置
		t, _ := template.ParseFiles("assets/testInput/login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		//请求的是登录数据，那么执行登录的逻辑判断
		r.ParseForm()
		fmt.Println(r.Form)
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func jsTestHandler(w http.ResponseWriter, req *http.Request) {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	formatter.JSON(w, http.StatusOK, struct {
		ID      string `json:"id"`
		Content string `json:"content"`
	}{ID: "8675309", Content: "Hello from Go!"})
}

func jsHandler(w http.ResponseWriter, req *http.Request) {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	formatter.JSON(w, http.StatusOK, struct {
		ID      string `json:"id"`
		Content string `json:"content"`
	}{ID: "8675309", Content: "Hello from Go!"})
}
