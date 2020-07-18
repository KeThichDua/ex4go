package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/KeThichDua/ex4go/db"
)

// Bai3 gom yeu cau b3
func Bai3() {
	// b3: insert 100 bản ghi vào user: sau đó viết 1 workerpool scantableuser lấy ra tên của các user in
	// ra màn hình (Dùng scan theo row) dùng 2 worker và thiết lập bộ đếm ${counter} - ${id} - ${name}
	fmt.Println("\n	Bai 3")
	err = d.Connect("mysql", "root:1@tcp(0.0.0.0:3306)/test")
	defer d.Data.Close()

	for i := 0; i < 100; i++ {
		// insert user
		temp := time.Now().UnixNano()
		user := db.User{Id: strconv.Itoa(i), Name: strconv.Itoa(i), Birth: temp, Created: temp, UpdatedAt: temp}
		err = d.InsertUser(user)
	}

	rows, err := d.Data.Rows(&db.User{})
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	size := 30
	mes := make(chan db.UserGroup, size)
	defer close(mes)
	wg := &sync.WaitGroup{}
	stop := make([]chan bool, 2)

	for i := 0; i < 2; i++ {
		stop[i] = make(chan bool)
		go Worker(mes, wg, stop[i])
	}

	index := 0
	bean := new(db.User)
	var userGroup db.UserGroup
	for rows.Next() {
		err = rows.Scan(bean)
		if err != nil {
			log.Fatal(err)
		}
		wg.Add(1)
		index++
		userGroup = db.UserGroup{Counter: index, User: bean}
		mes <- userGroup
	}

	wg.Wait()
	for i := 0; i < 2; i++ {
		stop[i] <- true
	}
}

// Worker Pool
func Worker(mes <-chan db.UserGroup, wg *sync.WaitGroup, stop chan bool) {
	for {
		select {
		case temp := <-mes:
			fmt.Println(temp.Counter, " - ", temp.User.Id, " - ", temp.User.Name)
			wg.Done()
		case <-stop:
			log.Println("Đã xóa worker")
			break
		}
	}
}
