package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Inst4 struct {
	Name    string
	Address string
}

var ins4 Inst4

func main4() {

	db, _ := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ayantika")

	db.Query("Update practice set Name =? where Address=?", "chakraborty", "halisahar")

	sel, _ := db.Query("select * from practice")
	for sel.Next() {
		sel.Scan(&ins4.Name, &ins4.Address)
		log.Printf("Name:- %s and Address:- %s", ins4.Name, ins4.Address)
	}

}
