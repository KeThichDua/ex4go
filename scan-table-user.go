package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/KeThichDua/ex4go/db"
)

// Worker Pool
func Worker(mes <-chan db.UserGroup, wg *sync.WaitGroup, stop chan bool) {
	for {
		select {
		case temp := <-mes:
			fmt.Println(temp.Counter, " - ", temp.UserTable.ID, " - ", temp.UserTable.Name)
			wg.Done()
		case <-stop:
			log.Println("Đã xóa worker")
			break
		}
	}
}
