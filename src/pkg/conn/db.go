package conn

import (
	"database/sql"
	"os"
	"time"
)

var db *sql.DB

func DbConnect() (err error) {
	db, err = sql.Open("mysql", os.Getenv("DB_URI"))
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(64)
	db.SetConnMaxIdleTime(time.Second * 5)
	return nil
}

func GetDB() *sql.DB {
	return db
}
