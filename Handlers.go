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
	}{ID: "18342008", Content: "Contents from the back-end"})
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
