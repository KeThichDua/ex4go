package main

import (
	"fmt"

	"github.com/KeThichDua/ex4go/db"
	_ "github.com/go-sql-driver/mysql"
)

var d db.Database
var err error

func main() {
	Bai1()
	Bai2()
	Bai3()

	fmt.Println("Done")
}
