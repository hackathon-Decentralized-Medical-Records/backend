package model

import (
	"gorm.io/gorm"
	"time"
)

// Registration 挂号
type Registration struct {
	gorm.Model
	PatientId uint    `gorm:"type:int;not null" json:"patientId" `
	Patient   Patient `gorm:"foreignKey:PatientId" json:"patient"`
	MedicId   uint    `gorm:"type:int;not null" json:"medicId" `
	Medic     Medic   `gorm:"foreignKey:MedicId" json:"medic"`
	// 挂号状态 1-已挂号 2-已就诊 3-已退号
	Status       int       `gorm:"type:int;not null" json:"status" `
	RegisterDate time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP" json:"registerDate"`
}

func (Registration) TableName() string {
	return "registration"
}
