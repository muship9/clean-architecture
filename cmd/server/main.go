package main

import (
	"log"
	"net/http"
	"simple-todo-api/pkg"
)

// メイン関数
func main() {

	err := pkg.NewDb()
	if err != nil {
		log.Println(err)
	}

	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/todos", pkg.HandleTodosRequest)
	server.ListenAndServe()
}
