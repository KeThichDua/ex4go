package main

import (
	"fmt"

	"github.com/KeThichDua/ex4go/db"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var d db.Database
	err := d.Connect("mysql", "root:1@tcp(0.0.0.0:3306)/test")
	defer d.Data.Close()

	Bai1(d, err)
	Bai2(d, err)
	Bai3(d, err)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Done")
}
