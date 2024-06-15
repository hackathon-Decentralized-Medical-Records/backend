package model

import "gorm.io/gorm"

type Case struct {
	gorm.Model
	// 患者id
	PatientId uint    `gorm:"type:int;not null" json:"patientId" `
	Patient   Patient `gorm:"foreignKey:PatientId" binding:"omitempty" json:"patient"`
	Content   string  `gorm:"type:text;not null" json:"content" `
}

func (Case) TableName() string {
	return "case"
}
