package service

import (
	"gorm.io/gorm"
	"service/global"
)

// Patient 患者
type Patient struct {
	gorm.Model
	Name    string `gorm:"type:varchar(100);not null" json:"name" `
	User_Id uint   `gorm:"type:int;not null" json:"userId" `
}

func (Patient) TableName() string {
	return "patient"
}

func InsertPatient(entity *Patient) {
	tx := global.GVA_DATABASE.Create(entity)
	if tx.Error != nil {
		panic(tx.Error)
	}
}
