package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/KeThichDua/ex4go/db"
)

// Bai2 gom yeu cau b2
func Bai2() {
	// tạo 1 transaction khi update birth thành công thì cộng 10 điểm vào point sau đó
	// sửa lại name thành $name + "updated " nếu 1 quá trình fail thì rollback, xong commit (CreateSesson)
	fmt.Println("\n Bai 2")
	err = d.Connect("mysql", "root:1@tcp(0.0.0.0:3306)/test")
	defer d.Data.Close()

	id := "temp"
	some := int64(10)
	err = Transaction(id, some)
	fmt.Println(err)
}

// Transaction thuc hien yeu cau
func Transaction(id string, some int64) error {
	session := d.Data.NewSession()
	defer session.Close()

	// add Begin() before any action
	if err = session.Begin(); err != nil {
		// if returned then will rollback automatically
		return err
	}

	// update birth
	user := db.User{Birth: time.Now().Unix()}
	c, err := session.Update(&user, db.User{Id: id})
	if err != nil {
		session.Rollback()
		return err
	} else if c == 0 {
		session.Rollback()
		return errors.New("Khong tim thay user")
	}

	// cong 10 diem
	point := db.Point{UserId: id}
	c1, err := session.Get(&point)
	if err != nil {
		session.Rollback()
		return err
	} else if !c1 {
		session.Rollback()
		return errors.New("Khong tim thay point")
	}
	point.Points = point.Points + some
	c, err = session.Update(&point, db.Point{UserId: id})
	if err != nil {
		session.Rollback()
		return err
	} else if c == 0 {
		session.Rollback()
		return errors.New("Khong tim thay point")
	}

	// sua name thanh name + "updated"
	user = db.User{Id: id}
	if has, err := session.Get(&user); err != nil {
		session.Rollback()
		return err
	} else if !has {
		session.Rollback()
		return errors.New("Khong tim thay user")
	}
	user.Name = user.Name + "updated"
	if c, err = session.Update(&user, db.User{Id: id}); err != nil {
		session.Rollback()
		return err
	} else if c == 0 {
		session.Rollback()
		return errors.New("Khong tim thay user")
	}
	// add Commit() after all actions
	return session.Commit()
}
