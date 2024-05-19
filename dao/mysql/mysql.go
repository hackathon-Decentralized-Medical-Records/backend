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
	// user 用户
	db.Table(User{}.TableName()).AutoMigrate(&User{})
	// medic 医生
	db.Table(Medic{}.TableName()).AutoMigrate(&Medic{})
	// patient 患者
	db.Table(Patient{}.TableName()).AutoMigrate(&Patient{})
	// department 科室
	db.Table(Department{}.TableName()).AutoMigrate(&Department{})
	// registration 挂号
	db.Table(Registration{}.TableName()).AutoMigrate(&Registration{})

}
