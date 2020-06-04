package db

import (
	"database/sql"
	"fmt"

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
	defer db.Close()
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
		log.Fatal(err)
	}

	fmt.Println(url)

	return url
}

// // Read
// func (db *sql.DB) Read() {
// 	rows, err := db.Query("SELECT * FROM users")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		p := new(Link)
// 		err := rows.Scan(&p.id, &p.url, &p.code)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Println(p.id, p.url, p.code)
// 	}
// }

// // UPDATE
// func (db *sql.DB) Update() {
// 	stmt, err := db.Prepare("UPDATE users SET code = ? WHERE id = ?")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	result, err := stmt.Exec(10, 1)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	affectNum, err := result.RowsAffected()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("update affect rows is ", affectNum)
// }

// // DELETE
// func (db *sql.DB) Delete() {
// 	stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	result, err := stmt.Exec(1)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	affectNum, err := result.RowsAffected()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("delete affect rows is ", affectNum)
// }

// // // Mysqlite3 - sqlite3 CRUD
// // func main() {
// // 	c, err := connectDB("sqlite3", "abc.db")
// // 	if err != "" {
// // 		print(err)
// // 	}

// // sql_table := `
// // CREATE TABLE IF NOT EXISTS "userinfo" (
// // 	"uid" INTEGER PRIMARY KEY AUTOINCREMENT,
// // 	"userurl" VARCHAR(64) NULL,
// // 	"departurl" VARCHAR(64) NULL,
// // 	"created" TIMESTAMP default (datetime('now', 'localtime'))
// // );
// // CREATE TABLE IF NOT EXISTS "userdeatail" (
// // 	"uid" INT(10) NULL,
// // 	"intro" TEXT NULL,
// // 	"profile" TEXT NULL,
// // 	PRIMARY KEY (uid)
// // );
// // 	`
// // 	c.Exec(sql_table)

// // 	c.Create()
// // 	fmt.Println("add action done!")

// // 	c.Read()
// // 	fmt.Println("get action done!")

// // 	c.Update()
// // 	fmt.Println("update action done!")

// // 	c.Delete()
// // 	fmt.Println("delete action done!")
// // }
