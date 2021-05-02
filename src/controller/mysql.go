package controller

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

const (
	DBUser     = "mysql_user"
	DBPassword = "mysql_password"
	DBHost     = "db"
	DB         = "sample"
)

func OpenMySqlConnection() *gorm.DB {
	connect := fmt.Sprintf("%s:%s@%s/%s", DBUser, DBPassword, DBHost, DB)
	db, err := gorm.Open("mysql", connect)
	if err != nil {
		panic("failed to connect database.")
	}

	db.LogMode(true)
	return db
}
