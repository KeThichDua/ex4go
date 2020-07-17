package db

import (
	"errors"
	"fmt"
	"log"

	"xorm.io/xorm"
)

// UserTable la nguoi dung
type UserTable struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Birth     int64  `json:"birth"`
	Created   int64  `json:"created"`
	UpdatedAt int64  `json:"update_at"`
}

// UserGroup nhom nguoi dung
type UserGroup struct {
	Counter   int
	UserTable *UserTable
}

// CreateTable la phuong thuc tao bang User
func (s *UserTable) CreateTable(data *xorm.Engine) error {
	err := data.CreateTables(UserTable{})
	err = data.Sync2(new(UserTable))
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("Tao bang user thanh cong.")
	return nil
}

// Insert de them du lieu user
func (s *UserTable) Insert(data *xorm.Engine, user UserTable) error {
	c, err := data.Insert(user)
	if c == 0 || err != nil {
		log.Println("Khong the them user")
		return errors.New("Loi insert")
	}
	fmt.Println("Them user thanh cong.")
	return err
}

// Update de sua du lieu user
func (s *UserTable) Update(data *xorm.Engine, user UserTable, conditions UserTable) error {
	c, err := data.Update(user, conditions)
	if c == 0 || err != nil {
		log.Println("Khong tim thay user")
		return errors.New("Khong tim thay user")
	}
	fmt.Println("Cap nhat user thanh cong.")
	return err
}

// Find de tim user theo so luong mong muon
func (s *UserTable) Find(data *xorm.Engine, id string, amount int) ([]*UserTable, error) {
	var users []*UserTable
	err := data.Where("i_d = ?", id).Limit(amount, 0).Find(&users)
	if err != nil || len(users) == 0 {
		log.Println("Khong tim thay user")
		return nil, err
	}
	fmt.Println("Da tim thay user voi id = ", id)
	return users, nil
}

// List de liet ke tat ca user
func (s *UserTable) List(data *xorm.Engine) ([]*UserTable, error) {
	var users []*UserTable
	err := data.Find(&users)
	if err != nil || len(users) == 0 {
		log.Println("Loi ko tim thay user.")
		return nil, err
	}
	fmt.Println("Lay ve tat ca user thanh cong.")
	return users, nil
}
