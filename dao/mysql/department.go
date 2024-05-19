package mysql

import "gorm.io/gorm"

// Department 科室
type Department struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);not null" json:"name" `
}

func (Department) TableName() string {
	return "department"
}
