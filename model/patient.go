package model

import "gorm.io/gorm"

// Patient 患者
type Patient struct {
	gorm.Model
	Name    string `gorm:"type:varchar(100);not null" json:"name" `
	User_Id uint   `gorm:"type:int;not null" json:"userId" `
}

func (Patient) TableName() string {
	return "patient"
}
