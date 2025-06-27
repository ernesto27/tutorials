package main

import (
	"goapi/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", controllers.HelloHandler)
	http.HandleFunc("/version", controllers.VersionHandler)
	http.ListenAndServe(":8080", nil)
}