package model

import "gorm.io/gorm"

// Medic 医生
type Medic struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);not null"  binding:"omitempty" form:"name" json:"name" `
	// 身份
	Profession string `gorm:"type:varchar(100);not null"  binding:"omitempty" form:"profession" json:"profession" `
	// 开始工作时间
	WorkTime string `gorm:"type:datetime;not null"  binding:"omitempty" form:"workTime" json:"workTime" `
	// 结束工作时间
	EndTime string `gorm:"type:datetime;not null"  binding:"omitempty"  form:"endTime" json:"endTime" `
	// 价格
	Price float64 `gorm:"type:float;not null"  binding:"omitempty" form:"price" json:"price" `
	// 简介
	Introduce    string `gorm:"type:varchar(1000);not null"  binding:"omitempty" json:"introduce" `
	DepartmentID uint   `gorm:"not null"   binding:"omitempty"  json:"department_id"` // 外键字段
	// 科室
	Department Department `gorm:"foreignKey:DepartmentID" binding:"omitempty" json:"department"`
	UserID     uint       `gorm:"not null" binding:"omitempty" json:"user_id"` // 外键字段
	// user
	User User   `gorm:"foreignKey:UserID" binding:"omitempty" json:"user"`
	Sort string `json:"sort" binding:"omitempty" `
}

func (Medic) TableName() string {
	return "medic"
}
