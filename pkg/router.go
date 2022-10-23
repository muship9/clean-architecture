package pkg

import (
	"net/http"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type Router interface {
	HandleTodosRequest(w http.ResponseWriter, r *http.Request)
}

func HandleTodosRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case GET:
		GetTodos(Db, w)
	case POST:
		AddTodo(Db, w, r)
	case PUT:
		EditTodo(Db, w, r)
	case DELETE:
		DeleteTodo(Db, w, r)
	default:
		// 指定メソッド以外はアクションを実行しない
		w.WriteHeader(405)
	}
}
