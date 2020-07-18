package db

import (
	"fmt"
	"log"

	"xorm.io/xorm"
)

// Database bao gom Data, bang User, bang Point
type Database struct {
	Data *xorm.Engine
}

// Connect tao ket noi database
func (s *Database) Connect(driverName string, dataSourceName string) error {
	engine, err := xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		log.Println("Loi", err)
		return err
	}
	fmt.Println("Da ket noi db")
	s.Data = engine
	return nil
}

// CreateTable la phuong thuc tao bang User
func (s *Database) CreateTable() error {
	err := s.Data.CreateTables(User{})
	err = s.Data.CreateTables(Point{})
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("Tao bang user, point thanh cong")
	return nil
}

// Sync2 de anh xa bang
func (s *Database) Sync2() error {
	err := s.Data.Sync2(new(User))
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("Anh xa bang thanh cong")
	return nil
}
