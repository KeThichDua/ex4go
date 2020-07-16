package db

import (
	"log"

	"xorm.io/xorm"
)

// Database bao gom Data la database ket noi mysql
type Database struct {
	Data *xorm.Engine
}

// Connect tao ket noi database
func (s *Database) Connect(driverName string, dataSourceName string) {
	engine, err := xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		log.Println(err)
	} else {
		log.Println()
		s.Data = engine
	}
}
