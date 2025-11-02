package Database

import (
	"database/sql"
	"diary/Schema"
	"diary/Util"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func GetSqlDB() *sql.DB {
	return db
}

func InitializeDatabase() error {

	dbPath := "../Database"
	if err := Util.CreateFolder(dbPath); err != nil {
		return err
	}

	var dbError error
	db, dbError = sql.Open("sqlite3", dbPath+"/diary.db")

	if dbError != nil {
		return dbError
	}

	if err := db.Ping(); err != nil {
		return err
	}

	fmt.Println("Database connected successfully.")

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(2)

	db.Exec("PRAGMA journal_mode=WAL;")
	db.Exec("PRAGMA synchronous=NORMAL;")
	db.Exec("PRAGMA cache_size=-64000;")
	db.Exec("PRAGMA temp_store=MEMORY;")
	db.Exec("PRAGMA mmap_size=268435456;")

	if err := Schema.CreateContentTable(db); err != nil {
		return err
	}

	return nil
}
