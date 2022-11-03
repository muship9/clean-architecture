package repository

import (
	"cleanArchitecture/server/domain/entity"
	"database/sql"
	"log"
)

// GetTodos DB から一致する data を取得
func GetTodos(db *sql.DB) []entity.Todos {
	rows, err := db.Query("SELECT * FROM todos WHERE user_id = 'testUser'")
	if err != nil {
		log.Println(err)
	}
	r := toTodoModel(rows)
	return r
}

// AddTodos DB に data を追加
func AddTodos(db *sql.DB, todo entity.Todos) error {
	_, err := db.Exec("INSERT INTO todos (todo_id , title , user_id) VALUES ($1, $2 ,$3)", todo.TodoId, todo.Title, todo.UserId)
	if err != nil {
		return err
	}
	return nil
}

// EditTodo DB から一致する data を更新
func EditTodo(db *sql.DB, todos entity.Todos) error {
	_, err := db.Exec("UPDATE todos SET title = $1 WHERE todo_id = $2", todos.Title, todos.TodoId)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTodo DB から一致する data を削除
func DeleteTodo(db *sql.DB, todo entity.Todos) error {
	_, err := db.Exec("DELETE FROM todos WHERE todo_id = $1", todo.TodoId)
	if err != nil {
		return err
	}
	return nil
}

// toTodoModel　TodoModel への変換を行う
func toTodoModel(rows *sql.Rows) []entity.Todos {
	var todoList []entity.Todos
	for rows.Next() {
		todo := entity.Todos{}
		err := rows.Scan(&todo.TodoId, &todo.Title, &todo.UserId, &todo.Created_at, &todo.Updated_at)
		if err != nil {
			return nil
		}
		todoList = append(todoList, todo)
	}
	return todoList
}
