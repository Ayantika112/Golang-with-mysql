package main

import (
	"database/sql"
	"fmt"
)

func main1() {
	fmt.Print("Go with MYSQL Tutorial")

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ayantika")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}
