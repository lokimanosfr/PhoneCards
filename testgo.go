package main

import (
	"fmt"
	"net/http"
)

func indexHandler(r http.ResponseWriter, req *http.Request) {
	r.Write([]byte("hello world"))
}

func main() {
	fmt.Println("listen port 9090")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":9090", nil)
}
