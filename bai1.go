package main

import (
	"fmt"
	"time"

	"github.com/KeThichDua/ex4go/db"
)

// Bai1 gom cac yeu cau b1
func Bai1(d db.Database, err error) {
	fmt.Println("\n	Bai 1")
	// anh xa bang
	err = d.CreateTable()
	err = d.Sync2()

	// insert user
	temp := time.Now().UnixNano()
	user := db.User{Id: "temp", Name: "temp", Birth: temp, Created: temp, UpdatedAt: temp}
	err = d.InsertUser(user)
	user = db.User{Id: "temp1", Name: "temp1", Birth: temp, Created: temp, UpdatedAt: temp}
	err = d.InsertUser(user)

	// cap nhat user theo id
	user = db.User{UpdatedAt: time.Now().UnixNano()} // cac truong cap nhat user
	userCondi := db.User{Id: "temp"}                 // dieu kien cap nhat user
	err = d.UpdateUser(user, userCondi)

	//tim user user id = temp
	user1, err := d.FindUser("temp")
	fmt.Println(user1.Id, "|", user1.Name, "|", user1.Birth, "|", user1.Created, "|", user1.UpdatedAt)

	// liet ke tat ca user
	arrY, err := d.ListUser()
	for i := range arrY {
		j := arrY[i]
		fmt.Println(j.Id, "|", j.Name, "|", j.Birth, "|", j.Created, "|", j.UpdatedAt)
	}

	// insert user_id vào user_point với số điểm 10.
	for i := range arrY {
		j := arrY[i]
		point := db.Point{UserId: j.Id, Points: 10, MaxPoints: 100}
		err = d.InsertPoint(point)
	}
}
