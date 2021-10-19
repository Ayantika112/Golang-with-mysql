package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Inst1 struct {
	Name    string
	Address string
}

var ins1 Inst1

func main2() {

	db, _ := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ayantika")

	db.Query("insert into practice values(?,?)", "ayantika", "Halisahar")

	sel, _ := db.Query("select * from practice")
	for sel.Next() {
		sel.Scan(&ins.Name, &ins.Address)
		log.Printf("Name:- %s and Address:- %s", ins.Name, ins.Address)
	}

}
