package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"upload_file_handler/config"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := config.Config.DB.DSN
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = connection
}
