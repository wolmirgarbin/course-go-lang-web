package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectaComBancoDeDados() *sql.DB {
	// conexao := "user=MainUser dbname=EXEMPLO_GO_LANG pass=MainPassword host=localhost sslmode=disable"
	db, err := sql.Open("mysql", "MainUser:MainPassword@/EXEMPLO_GO_LANG")
	if err != nil {
		panic(err)
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
