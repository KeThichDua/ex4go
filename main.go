package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/KeThichDua/ex4go/db"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var err error
	var x db.Database
	var y db.UserTable
	var z db.PointTable
	var arrY []*db.UserTable
	err = x.Connect("mysql", "root:1@tcp(0.0.0.0:3306)/test")
	defer x.Data.Close()

	// anh xa bang
	err = y.CreateTable(x.Data)
	err = z.CreateTable(x.Data)

	// insert user
	temp := time.Now().Format("15:04:05")
	temp1 := time.Now().UnixNano()
	y = db.UserTable{ID: temp, Name: temp, Birth: temp1, Created: temp1, UpdatedAt: temp1}
	err = y.Insert(x.Data, y)

	// cap nhat user theo id
	y = db.UserTable{UpdatedAt: temp1}    // cac truong cap nhat user
	temp3 := db.UserTable{ID: "17:29:06"} // dieu kien cap nhat user
	err = y.Update(x.Data, y, temp3)

	//liet ke 3 user theo id = 17:29:06
	arrY, err = y.Find(x.Data, "17:29:06", 3)
	for i := range arrY {
		j := arrY[i]
		fmt.Println(j.ID, "|", j.Name, "|", j.Birth, "|", j.Created, "|", j.UpdatedAt)
	}

	// liet ke tat ca user
	arrY, err = y.List(x.Data)
	for i := range arrY {
		j := arrY[i]
		fmt.Println(j.ID, "|", j.Name, "|", j.Birth, "|", j.Created, "|", j.UpdatedAt)
	}

	// insert user_id vào user_point với số điểm 10.
	for i := range arrY {
		j := arrY[i]
		z = db.PointTable{UserID: j.ID, Points: 10, MaxPoints: 100}
		err = z.Insert(x.Data, z)
	}
	fmt.Println("Xong bai 1")

	// tạo 1 transaction khi update birth thành công thì cộng 10 điểm vào point sau đó
	// sửa lại name thành $name + "updated " nếu 1 quá trình fail thì rollback, xong commit (CreateSesson)
	err = Transaction(x)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Xong bai 2")

	// b3: insert 100 bản ghi vào user: sau đó viết 1 workerpool scantableuser lấy ra tên của các user in
	// ra màn hình (Dùng scan theo row) dùng 2 worker và thiết lập bộ đếm ${counter} - ${id} - ${name}
	for i := 0; i < 100; i++ {
		// insert user
		temp = time.Now().Format("15:04:05")
		temp1 = time.Now().UnixNano()
		y = db.UserTable{ID: temp, Name: temp, Birth: temp1, Created: temp1, UpdatedAt: temp1}
		err = y.Insert(x.Data, y)
	}

	rows, err := x.Data.Rows(&db.UserTable{})
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	size := 40
	mes := make(chan db.UserGroup, size)
	defer close(mes)
	wg := &sync.WaitGroup{}
	stop := make([]chan bool, 2)

	for i := 0; i < 2; i++ {
		stop[i] = make(chan bool)
		go Worker(mes, wg, stop[i])
	}

	index := 0
	bean := new(db.UserTable)
	var userGroup db.UserGroup
	for rows.Next() {
		err = rows.Scan(bean)
		if err != nil {
			log.Fatal(err)
		}
		wg.Add(1)
		index++
		userGroup = db.UserGroup{Counter: index, UserTable: bean}
		mes <- userGroup
	}

	wg.Wait()
	for i := 0; i < 2; i++ {
		stop[i] <- true
	}
	fmt.Println("Xong bai 3")
}
