package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"service/global"
)

func Gorm() *gorm.DB {
	db, err := gorm.Open(mysql.Open(global.GVA_SETTING.Database.Dsn()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func RegisterTables() {
	db := global.GVA_DATABASE
	// TODO: 表的完善
	err := db.AutoMigrate()

	if err != nil {
		global.GVA_LOGRUS.Panicf("register table error: %v", err)
		os.Exit(0)
	}
	global.GVA_LOGRUS.Info("register table success")
}
