package controller

import (
	"fmt"

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
	connect := fmt.Sprintf("%s:%s@%s/%s", DBUser, DBPassword, DBHost, DB)
	fmt.Println(connect)
	db, err := gorm.Open("mysql", connect)
	if err != nil {
		panic("failed to connect database.")
	}

	db.LogMode(true)
	return db
}
