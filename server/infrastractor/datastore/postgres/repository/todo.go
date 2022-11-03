package repository

import (
	"cleanArchitecture/server/domain/entity"
	"database/sql"
	"log"
)

// GetTodos DB から一致する data を取得
func GetTodos(db *sql.DB) *sql.Rows {
	rows, err := db.Query("SELECT * FROM todos WHERE user_id = 'testUser'")

	if err != nil {
		log.Println(err)
	}

	// TODO : Todos モデルにマッピングしたものを返す
	return rows
}

func AddTodos(db *sql.DB, todo entity.Todos) error {
	_, err := db.Exec("INSERT INTO todos (todo_id , title , user_id) VALUES ($1, $2 ,$3)", todo.TodoId, todo.Title, todo.UserId)
	if err != nil {
		return err
	}
	return nil
}

func EditTodo(db *sql.DB, todos entity.Todos) error {
	_, err := db.Exec("UPDATE todos SET title = $1 WHERE todo_id = $2", todos.Title, todos.TodoId)

	if err != nil {
		return err
	}

	return nil
}

func DeleteTodo(db *sql.DB, todo entity.Todos) error {
	_, err := db.Exec("DELETE FROM todos WHERE todo_id = $1", todo.TodoId)

	if err != nil {
		return err
	}

	return nil
}
