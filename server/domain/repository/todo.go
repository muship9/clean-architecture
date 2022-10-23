package repository

import "database/sql"

type TodoRepository interface {
	GetTodos(db *sql.DB) *sql.Rows
}
