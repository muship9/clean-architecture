package repository

import (
	"cleanArchitecture/server/domain/entity"
	"database/sql"
)

type TodoRepository interface {
	GetTodos(db *sql.DB) *sql.Rows
	AddTodos(db *sql.DB, todos entity.Todos) error
	EditTodo(db *sql.DB, todos entity.Todos) error
	DeleteTodo(db *sql.DB, todos entity.Todos) error
}
