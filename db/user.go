package db

import (
	"errors"
	"fmt"
	"log"
)

// User la nguoi dung
type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Birth     int64  `json:"birth"`
	Created   int64  `json:"created"`
	UpdatedAt int64  `json:"update_at"`
}

// UserGroup nhom nguoi dung
type UserGroup struct {
	Counter int
	User    *User
}

// InsertUser de them du lieu user
func (s *Database) InsertUser(user User) error {
	c, err := s.Data.Insert(user)
	if c == 0 {
		log.Fatal("Khong the them user")
		return errors.New("Loi insert")
	}
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Them user thanh cong.")
	return err
}

// UpdateUser de sua du lieu user
func (s *Database) UpdateUser(user User, conditions User) error {
	c, err := s.Data.Update(user, conditions)
	if c == 0 {
		log.Fatal("Khong tim thay user")
		return errors.New("Khong tim thay user")
	}
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Cap nhat user thanh cong.")
	return err
}

// FindUser tim kiem 1 user
func (s *Database) FindUser(id string) (*User, error) {
	user := &User{Id: id}
	c, err := s.Data.Get(user)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if !c {
		log.Println("Khong tim thay user voi id: ", id)
		return nil, errors.New("Khong tim thay")
	}
	return user, nil
}

// ListUser de liet ke tat ca user
func (s *Database) ListUser() ([]*User, error) {
	var users []*User
	err := s.Data.Desc("id").Find(&users)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if len(users) == 0 {
		log.Println("Loi ko tim thay user")
		return nil, errors.New("Database rong")
	}
	fmt.Println("Lay ve tat ca user thanh cong.")
	return users, nil
}
