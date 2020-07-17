package db

import (
	"errors"
	"fmt"
	"log"

	"xorm.io/xorm"
)

// PointTable la bang diem cua nguoi dung
type PointTable struct {
	UserID    string `json:"user_id"`
	Points    int64  `json:"points"`
	MaxPoints int64  `json:"max_points"`
}

// CreateTable la phuong thuc tao bang Point
func (s *PointTable) CreateTable(data *xorm.Engine) error {
	err := data.CreateTables(PointTable{})
	err = data.Sync2(new(PointTable))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Tao bang point thanh cong.")
	return nil
}

// Insert de them du lieu point
func (s *PointTable) Insert(data *xorm.Engine, point PointTable) error {
	c, err := data.Insert(point)
	if c == 0 || err != nil {
		log.Println("Khong them duoc point")
		return errors.New("Khong them point")
	}
	fmt.Println("Them point thanh cong.")
	return err
}

// Update de sua du lieu point
func (s *PointTable) Update(data *xorm.Engine, point PointTable, conditions PointTable) error {
	c, err := data.Update(point, conditions)
	if c == 0 || err != nil {
		log.Println("Khong sua duoc point")
		return errors.New("Khong sua point")
	}
	fmt.Println("Cap nhat point thanh cong.")
	return err
}
