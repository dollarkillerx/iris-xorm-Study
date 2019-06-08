/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-8
* Time: 下午10:51
* */
package main

import "net/http"

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world"))
	})

	http.ListenAndServe(":8580",nil)
}
