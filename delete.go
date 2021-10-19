package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Inst3 struct {
	Name    string
	Address string
}

var ins3 Inst3

func main3() {

	db, _ := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ayantika")

	db.Query("delete from practice where Name=?", "ayantika")

	sel, _ := db.Query("select * from practice")
	for sel.Next() {
		sel.Scan(&ins3.Name, &ins3.Address)
		log.Printf("Name:- %s and Address:- %s", ins3.Name, ins3.Address)
	}

}
