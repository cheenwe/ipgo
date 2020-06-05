package db

import (
	"database/sql"

	"log"

	_ "github.com/mattn/go-sqlite3" // sqlite3 dirver
)

const (
	dbDriverName = "sqlite3"
	dbName       = "./db/data.db"
)

func checkErr(e error) bool {
	if e != nil {
		log.Fatal(e)
		return true
	}
	return false
}

func ConnectDB(driverName string, dbName string) (db *sql.DB) {
	db, err := sql.Open(dbDriverName, dbName)
	if err != nil {
		return nil
	}
	return db
}

func CreateLinkTable(db *sql.DB) error {
	sql := `
	CREATE TABLE IF NOT EXISTS "links" (
	"id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"url" VARCHAR(256) NULL,
	"code" VARCHAR(24) NULL,
	"created_at" datetime DEFAULT CURRENT_TIMESTAMP
	);
	CREATE INDEX  "index_links_on_code" ON "links" (
		"code"	ASC
	);
	`
	// defer db.Close()
	_, err := db.Exec(sql)
	return err
}

// Create
func InsertLink(db *sql.DB, url string, code string) error {
	sql := `insert into links (url, code) values(?,?)`
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(url, code)
	return err
}

func QueryLink(db *sql.DB, code string) (url string) {
	stmt, err := db.Prepare("select  url from links where code = ?  ORDER BY id DESC;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(code).Scan(&url)

	if err != nil {
		url = "0"
	}
	return url
}
