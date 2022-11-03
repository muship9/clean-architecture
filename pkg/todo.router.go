package pkg

import (
	"cleanArchitecture/server/domain/entity"
	"cleanArchitecture/server/infrastractor/datastore/postgres/repository"
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
)

// GetTodos DB からデータを全件取得して一覧を返す
func GetTodos(db *sql.DB, w http.ResponseWriter) {

	todos := repository.GetTodos(db)

	// DB から取得した data を Json 構造体 にマッピングする
	var todoResponses []entity.EncodeTodo
	for _, v := range todos {
		todoResponses = append(todoResponses, entity.EncodeTodo{
			TodoId: v.TodoId,
			Title:  v.Title,
			UserId: v.UserId,
		})
	}
	output, _ := json.MarshalIndent(todoResponses, "", "\t\t")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(output)
}

// AddTodo クライアントから送られてきたデータをもとに DB に追加
func AddTodo(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var todo entity.Todos
	var err error
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	var todoRequest entity.TodoRequest

	err = json.Unmarshal(body, &todoRequest)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Connection Failed"))
		return
	}

	if todoRequest.TodoId == "" {
		todoRequest.TodoId = uuid.NewString()
	}

	todo = entity.Todos{
		TodoId: todoRequest.TodoId,
		Title:  todoRequest.Title,
		UserId: todoRequest.UserId,
	}

	err = repository.AddTodos(db, todo)
	if err != nil {
		log.Print(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("success"))

}

// EditTodo クライアントから送られてきたデータをもとに DB を更新する
func EditTodo(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var todos entity.Todos
	var err error
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	var editTodoRequest entity.EditTodoRequest

	err = json.Unmarshal(body, &editTodoRequest)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Connection Failed"))
		return
	}

	todos = entity.Todos{
		TodoId: editTodoRequest.TodoId,
		Title:  editTodoRequest.Title,
	}

	if todos.TodoId == "" {
		w.WriteHeader(400)
		w.Write([]byte("TodoId がないため処理を中断します。"))
		return
	}

	err = repository.EditTodo(db, todos)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Connection failed"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write([]byte("success"))

}

// DeleteTodo 指定データを DB から削除する
func DeleteTodo(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var todo entity.Todos
	var err error
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	var todoRequest entity.TodoRequest

	err = json.Unmarshal(body, &todoRequest)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Connection Failed"))
		return
	}

	todo = entity.Todos{
		TodoId: todoRequest.TodoId,
		Title:  todoRequest.Title,
		UserId: todoRequest.UserId,
	}

	if todo.TodoId == "" {
		w.WriteHeader(400)
		w.Write([]byte("TodoId がないため処理を中断します。"))
		return
	}

	err = repository.DeleteTodo(db, todo)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Connection failed"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write([]byte("success"))

}
