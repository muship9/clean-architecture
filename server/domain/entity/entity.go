package entity

import "time"

type Todos struct {
	TodoId     string
	Title      string
	UserId     string
	Created_at time.Time
	Updated_at time.Time
}

type EncodeTodo struct {
	TodoId string `json:"id"`
	Title  string `json:"title"`
	UserId string `json:"todo"`
}

type TodoRequest struct {
	TodoId string `json:"todoId"`
	Title  string `json:"title"`
	UserId string `json:"user_id"`
}

type EditTodoRequest struct {
	TodoId string `json:"todoId"`
	Title  string `json:"title"`
}
