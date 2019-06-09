/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-9
* Time: 上午9:10
* */
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/",hello)
	if err := http.ListenAndServe(":9090", nil);err==nil {
		fmt.Println("http://127.0.0.1:9090  runing")
	}
}

func hello(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.PostForm.Get("hello"))

}
