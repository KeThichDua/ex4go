package main

import (
	"errors"
	"time"

	"github.com/KeThichDua/ex4go/db"
)

// Transaction tu dong cap nhat
func Transaction(x db.Database) error {
	session := x.Data.NewSession()
	defer session.Close()
	var err error

	// add Begin() before any action
	if err = session.Begin(); err != nil {
		// if returned then will rollback automatically
		return err
	}

	// update birth
	y := db.UserTable{Birth: time.Now().Unix()}
	c, err := session.Where("i_d = ?", "17:29:06").Update(&y)
	if err != nil {
		session.Rollback()
		return err
	} else if c == 0 {
		session.Rollback()
		return errors.New("Khong tim thay user")
	}

	// cong 10 diem
	z := &db.PointTable{UserID: "17:29:06"}
	c1, err := session.Get(z)
	if err != nil {
		session.Rollback()
		return err
	} else if !c1 {
		session.Rollback()
		return errors.New("Khong tim thay point")
	}
	z.Points = z.Points + 10
	c, err = session.Where("user_i_d = ?", "17:29:06").Update(z)
	if err != nil {
		session.Rollback()
		return err
	} else if c == 0 {
		session.Rollback()
		return errors.New("Khong the them point")
	}

	// sua name thanh name + "updated"
	y = db.UserTable{ID: "17:29:06"}
	if has, err := session.Get(y); err != nil {
		session.Rollback()
		return err
	} else if !has {
		session.Rollback()
		return errors.New("Khong tim thay user")
	}
	y.Name = y.Name + "updated"
	if c, err = session.Where("i_d = ?", "17:29:06").Update(y); err != nil {
		session.Rollback()
		return err
	} else if c == 0 {
		session.Rollback()
		return errors.New("Khong the sua user")
	}
	// add Commit() after all actions
	return session.Commit()
}
