package main

import (
	"./externals/github"
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	github.Get("Nicovideo", "VideoCollection", 5)
	fmt.Fprintf(writer, "Ok")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":4000", nil)
}
