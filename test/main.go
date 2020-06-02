package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbDriverName = "sqlite3"
	dbName       = "./data.db"
)

type Link struct {
	url  string
	code string
}

func main() {
	db, err := sql.Open(dbDriverName, dbName)
	if checkErr(err) {
		return
	}
	// err = createTable(db)
	// if checkErr(err) {
	// 	return
	// }
	err = insertData(db, Link{"www.baidu.com", "123"})
	if checkErr(err) {
		return
	}
	res, err := queryData(db, "123")
	if checkErr(err) {
		return
	}
	fmt.Println(res)

}

func createTable(db *sql.DB) error {
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
	_, err := db.Exec(sql)
	return err
}

func insertData(db *sql.DB, link Link) error {
	sql := `insert into links (url, code) values(?,?)`
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(link.url, link.code)
	return err
}

func queryData(db *sql.DB, code string) (list []Link, e error) {
	sql := `select  * from links where(code = ? ) order by id desc`
	stmt, err := db.Prepare(sql)
	rows, err := stmt.Query(code)

	var result = make([]Link, 0)
	for rows.Next() {
		var link, code string
		rows.Scan(&link, &code)
		result = append(result, Link{link, code})
	}
	return result, err

}

func checkErr(e error) bool {
	if e != nil {
		log.Fatal(e)
		return true
	}
	return false
}
