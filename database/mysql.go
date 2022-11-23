package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/belajar_database?parseTime=True")
	if err != nil {
		panic(err)
	}
	return db
}
// defer db.Close()
// db.SetMaxIdleConns(10)
// db.SetMaxOpenConns(100)
// db.SetConnMaxIdleTime(5 * time.Minute)
// db.SetConnMaxLifetime(60 * time.Minute)
