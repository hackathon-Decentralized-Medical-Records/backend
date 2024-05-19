package mysql

import "gorm.io/gorm"

// Medic 医生
type Medic struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);not null" json:"name" `
	// 职业
	Profession string `gorm:"type:varchar(100);not null" json:"profession" `
	// 科室
	Department_Id int  `gorm:"type:int;not null" json:"departmentId" `
	User_Id       uint `gorm:"type:int;not null" json:"userId" `
}

func (Medic) TableName() string {
	return "medic"
}

func InsertMedic(m *Medic) (int, string) {
	db.Create(m)
	return 200, "注册成功"
}
