package main

import (
	"fmt"

	"github.com/KeThichDua/ex4go/db"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("HI")
	var db db.Database
	db.Connect("mysql", "root:1@tcp(0.0.0.0.3306)/test")
	defer db.Data.Close()

	
}
