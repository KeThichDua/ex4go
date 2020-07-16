package db

// User la nguoi dung
type User struct {
	ID        string
	Name      string
	Birth     int64
	Created   int64
	UpdatedAt int64
}

// CreateTable la ham tao bang
func CreateTable() {
	// err := engine.Sync2(new(User))
}
