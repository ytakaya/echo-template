package controller

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
)

const (
	DBUser     = "mysql_user"
	DBPassword = "mysql_password"
	DBHost     = "tcp(db:3306)"
	DB         = "sample"
)

func OpenMySqlConnection() *gorm.DB {
	connect := fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4", DBUser, DBPassword, DBHost, DB)
	for count := 0; count < 30; count++ {
		db, err := gorm.Open("mysql", connect)
		if err == nil {
			db.LogMode(true)
			return db
		}
		log.Println("Not ready. Retry connecting...")
		time.Sleep(time.Second)
	}

	panic("failed to connect database.")
}
