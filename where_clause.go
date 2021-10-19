package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Inst5 struct {
	Name    string
	Address string
}

var ins5 Inst5

func main5() {

	db, _ := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ayantika")
	NAME := "ayantika chakravborty"

	db.QueryRow("select Name,Address from practice where Name=?", NAME).Scan(&ins5.Name, &ins5.Address)
	if ins.Name == NAME {
		fmt.Println("present")
	} else {
		fmt.Println("not present")
	}

}
