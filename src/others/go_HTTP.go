// go_HTTP
package main

import (
	//"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("hello go"))
	})
	http.ListenAndServe("127.0.0.1:8080", nil)
	//在浏览器输入 localhost:8080/hello
}
