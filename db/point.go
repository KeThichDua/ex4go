package db

import (
	"errors"
	"fmt"
	"log"
)

// Point la bang diem cua nguoi dung
type Point struct {
	UserId    string `json:"user_id"`
	Points    int64  `json:"points"`
	MaxPoints int64  `json:"max_points"`
}

// InsertPoint de them du lieu point
func (s *Database) InsertPoint(point Point) error {
	c, err := s.Data.Insert(point)
	if c == 0 {
		log.Println("Khong them duoc point")
		return errors.New("Khong them point")
	}
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("Them point thanh cong.")
	return err
}

// UpdatePoint de sua du lieu point
func (s *Database) UpdatePoint(point Point, conditions Point) error {
	c, err := s.Data.Update(point, conditions)
	if c == 0 {
		log.Println("Khong sua duoc point")
		return errors.New("Khong sua point")
	}
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("Cap nhat point thanh cong.")
	return err
}
