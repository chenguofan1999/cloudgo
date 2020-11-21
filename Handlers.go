package cloudgo

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/unrolled/render"
)

func jsonHandler(w http.ResponseWriter, req *http.Request) {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	formatter.JSON(w, http.StatusOK, struct {
		ID      string `json:"id"`
		Content string `json:"content"`
	}{ID: "8675309", Content: "Hello from Go!"})
}

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

		formatter := render.New(render.Options{
			Directory:  "templates",
			Extensions: []string{".html"},
			IndentJSON: true,
		})

		formatter.HTML(w, http.StatusOK, "index", struct {
			Un string `json:"username"`
			Pw string `json:"password"`
		}{Un: r.Form["username"][0], Pw: r.Form["password"][0]})
	}

}

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<!DOCTYPE html>
	<html>
	
	<head>
		<title></title>
	</head>
	
	<body style="color: rgb(0, 0, 0); background-color: rgb(240, 240, 240);">
		<h2 style="text-align: center;"><span style="font-family: Impact, Charcoal, sans-serif; color: rgb(40, 50, 78);">cloudGo</span></h2>
		<p><span style="font-family: Verdana, Geneva, sans-serif; color: rgb(40, 50, 78);">1. 静态文件服务</span></p>
		<p><span style="font-family: Verdana, Geneva, sans-serif; color: rgb(40, 50, 78);"><em>./static/</em></span></p>
		<p><a href="./static/"><span style="color: rgb(71, 85, 119);">转到静态文件服务</span></a></p>
		<p><span style="font-family: Verdana, Geneva, sans-serif; color: rgb(40, 50, 78);">2. 简单js访问</span></p>
		<p><span style="font-family: Verdana, Geneva, sans-serif; color: rgb(40, 50, 78);"><em>./js</em></span></p>
		<p><a href="./js"><span style="color: rgb(71, 85, 119);">转到简单js访问</span></a></p>
		<p><span style="font-family: Verdana, Geneva, sans-serif; color: rgb(40, 50, 78);">3. 表单</span></p>
		<p><span style="font-family: Verdana, Geneva, sans-serif; color: rgb(40, 50, 78);"><em>./login</em></span></p>
		<p><span style="color: rgb(71, 85, 119);"><a href="./login">转到表单</a></span></p>
	</body>
	
	</html>`)
}
