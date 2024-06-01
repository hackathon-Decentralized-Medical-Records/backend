package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"service/config"
)

var db *gorm.DB

// Initialize mysql
func Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.MysqlUser,
		config.MysqlPassword,
		config.MysqlHost,
		config.MysqlPort,
		config.MysqlDbname,
	)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection test failed:", err)
	}
	db.Table(User{}.TableName()).AutoMigrate(&User{})
	db.Table(Department{}.TableName()).AutoMigrate(&Department{})
	db.AutoMigrate(&Medic{}, &Department{}, &User{})
	db.Table(Patient{}.TableName()).AutoMigrate(&Patient{})
	db.AutoMigrate(Registration{}, Medic{})
	db.AutoMigrate(&Case{}, &Patient{})
	db.AutoMigrate(&Accredut{})

}
