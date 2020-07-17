package db

import (
	"errors"
	"fmt"

	"xorm.io/xorm"
)

// Database bao gom Data, bang User, bang Point
type Database struct {
	Data *xorm.Engine
}

// Connect tao ket noi database
func (s *Database) Connect(driverName string, dataSourceName string) error {
	engine, err := xorm.NewEngine(driverName, dataSourceName)
	s.Data = engine
	if err != nil || engine != nil {
		fmt.Println("Loi ket noi db.")
		return errors.New("Loi db")
	}
	fmt.Println("Dang ket noi db.")
	return nil
}
