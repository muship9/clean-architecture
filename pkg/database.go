package pkg

import (
	"database/sql"
	"log"

	// postgres ドライバ
	_ "github.com/lib/pq"
)

var Db *sql.DB

func NewDb() error {
	var err error
	Db, err = sql.Open("postgres", "host=db user=admin password=password dbname=cleanArchitectureDb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return err
}
