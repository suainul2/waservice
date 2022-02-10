package databases

import (
	"fmt"
	"waservice/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Model interface{}
}

var Db *gorm.DB

func MysqlRun() {
	dsn := configs.Mysql.Address
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if Db == nil {
		fmt.Println("databse berjalan")
		Db = db
		Migrate()
	}
}
