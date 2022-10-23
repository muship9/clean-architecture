package repository

import (
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
