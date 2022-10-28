package repository

import (
	"cleanArchitecture/server/domain/entity"
	"database/sql"
)

type TodoRepository interface {
	GetTodos(db *sql.DB) *sql.Rows
	AddTodos(db *sql.DB, todos entity.Todos) error
}
