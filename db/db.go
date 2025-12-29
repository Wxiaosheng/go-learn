package db

import (
	"errors"
	"strings"

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

	db.AutoMigrate(&Post{})

	// db.Create(&User{ID: 1, Username: "Victree", Email: "alice@example.com"})
	// var u User
	// db.Find(&u)
	// fmt.Printf("%+v\n", u)
}

/* 判断是否为重复键错误 */
func IsDuplicatedKeyError(err error) bool {
	// 1. 先尝试 GORM 标准错误（未来驱动完善后可生效）
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return true
	}

	// 2. 回退到字符串匹配（兼容当前 SQLite/MySQL/PG）
	if err == nil {
		return false
	}
	s := err.Error()
	return strings.Contains(s, "UNIQUE constraint failed") ||
		strings.Contains(s, "Duplicate entry") ||
		strings.Contains(s, "duplicate key value violates unique constraint") ||
		strings.Contains(s, "violates unique constraint")
}
