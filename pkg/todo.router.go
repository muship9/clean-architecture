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

	rows := repository.GetTodos(db)

	var data []entity.Todos

	// TODOにEntityをマッピングし、返却用のスライスに追加
	for rows.Next() {
		todo := entity.Todos{}
		err := rows.Scan(&todo.TodoId, &todo.Title, &todo.UserId, &todo.Created_at, &todo.Updated_at)
		if err != nil {
			log.Print(err)
			return
		}
		data = append(data, todo)
	}

	// DB から取得した data を Json 構造体 にマッピングする
	var todoResponses []entity.EncodeTodo
	for _, v := range data {
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

	_, err = db.Exec("INSERT INTO todos (todo_id , title , user_id) VALUES ($1, $2 ,$3)", todo.TodoId, todo.Title, todo.UserId)
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

	_, err = db.Exec("UPDATE todos SET title = $1 WHERE todo_id = $2", todos.Title, todos.TodoId)

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

	_, err = db.Exec("DELETE FROM todos WHERE todo_id = $1", todo.TodoId)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Connection failed"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write([]byte("success"))

}
