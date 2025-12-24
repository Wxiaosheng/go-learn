package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var sqlDB *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB = db

	db.AutoMigrate(&User{})

	// db.Create(&User{ID: 1, Username: "Victree", Email: "alice@example.com"})
	// var u User
	// db.Find(&u)
	// fmt.Printf("%+v\n", u)
}
