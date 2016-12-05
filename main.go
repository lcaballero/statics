package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lcaballero/statics/handling"
	"net/http"
)

func main() {
	server()
}

func server() {
	ip := ":5555"
	root := ".www"
	index := "index.html"
	prefix := "/assets/"

	router := mux.NewRouter()
	handling.BasicStaticRoute(
		root,
		prefix,
		index,
		router.NewRoute())

	fmt.Printf("binding sever to %s\n", ip)

	http.ListenAndServe(ip, router)
}
