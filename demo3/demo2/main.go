/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-9
* Time: 上午9:42
* */
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/",ccs)
	http.ListenAndServe(":8085",nil)
}

func ccs(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()
	fmt.Println("username: ",template.HTMLEscapeString(r.PostForm.Get("username")))
	fmt.Println("password: ",template.HTMLEscapeString(r.PostForm.Get("password")))
	template.HTMLEscape(w,[]byte(r.PostForm.Get("username")))
}
