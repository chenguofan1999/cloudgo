package cloudgo

import (
	"fmt"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		// t, _ := template.ParseFiles("login.gtpl")
		// log.Println(t.Execute(w, nil))
		fmt.Fprintln(w, loginHTML)
	} else {
		//请求的是登录数据，那么执行登录的逻辑判断
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}
