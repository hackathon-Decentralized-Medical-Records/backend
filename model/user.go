package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(100);not null" json:"username" `
	PassWord string `gorm:"type:varchar(65);not null" json:"password" `
	Email    string `gorm:"type:varchar(100);not null" json:"email" `
	Role     int    `gorm:"type:int" json:"role"`
	Address  string `gorm:"type:varchar(42);column:address" json:"address"`
}

type Response struct {
	Data interface{} `json:"data"`
}

func (User) TableName() string {
	return "user"
}
