package initial

import (
	"errors"
	"go-learn/global"
	"go-learn/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/** Gorm 初始化 */
func InitGorm() {

	if global.BLOGS_CONFIG.System.DbType == "mysql" {
		initMysqlGorm()
	}

}

/** MySQL Gorm 初始化 */
func initMysqlGorm() {
	m := global.BLOGS_CONFIG.Mysql

	// 没有配置数据库名则不初始化
	if m.DBName == "" {
		panic(errors.New("数据库名配置为空"))
	}

	db, err := gorm.Open(mysql.Open(m.DSN()), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)

	// 将 gorm.DB 对象赋值给全局变量
	global.BLOGS_DB = db
}

/** 自动迁移数据库表 */
func InitTables() {
	db := global.BLOGS_DB

	// 自动迁移数据库表
	err := db.AutoMigrate(

		/* 用户管理 */
		&model.User{},

		/* 文章管理 */
		&model.Post{},
	)

	if err != nil {
		panic(err)
	}
}
