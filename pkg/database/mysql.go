package database

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	defaultDB *gorm.DB
	lock      = sync.Mutex{}
)

func NewDB() *gorm.DB {
	lock.Lock()
	defer lock.Unlock()
	if defaultDB == nil {
		//连接MYSQL, 获得DB类型实例
		dsn := "root:root@tcp(127.0.0.1:3306)/boss?charset=utf8&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("数据库连接错误" + err.Error())
		}
		defaultDB = db
	}
	return defaultDB
}
