package service

import (
	"gorm.io/gorm"
	"service/global"
)

// Department 科室
type Department struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);not null" json:"name" `
}

func (Department) TableName() string {
	return "department"
}

func GetDepartment() []Department {

	var departments []Department
	global.GVA_DATABASE.Find(&departments)
	return departments
}
