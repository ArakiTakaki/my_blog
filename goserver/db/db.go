package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	_ "github.com/mattn/go-sqlite3"
)

func GetDB() {

	content, err := ioutil.ReadFile("db/queries/users_table.sql")
	if err != nil {
		fmt.Println("======= ERROR =======")
		panic(err)
	}
	fmt.Println(string(content))
	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// テーブル作成
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS "BOOKS" ("ID" INTEGER PRIMARY KEY, "TITLE" VARCHAR(255))`,
	)
	if err != nil {
		panic(err)
	}

	res, err := db.Exec(
		`INSERT INTO BOOKS (ID, TITLE) VALUES (?, ?)`,
		2,
		"title",
	)
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("====== PRINT ID =======")
	fmt.Println(id)
	// fmt.Println(string(item))
}
