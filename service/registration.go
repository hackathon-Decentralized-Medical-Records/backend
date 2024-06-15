package service

import (
	"gorm.io/gorm"
	"service/global"
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

func InsertRegistration(entity *Registration) {
	tx := global.GVA_DATABASE.Create(entity)
	if tx.Error != nil {
		panic(tx.Error)
	}
}

func UpdateStatus(id uint, status int) {
	tx := global.GVA_DATABASE.Model(&Registration{}).Where("id=?", id).Update("status", status)
	if tx.Error != nil {
		panic(tx.Error)
	}
}

// GetRecordByMedic 根据医生id获取挂号记录
func GetRecordByMedic(medicId uint) []Registration {
	var records []Registration
	tx := global.GVA_DATABASE.Preload("Patient").Preload("Medic").Preload("Medic.Department").Where("medic_id=?", medicId).Find(&records)
	if tx.Error != nil {
		panic(tx.Error)
	}

	return records
}

// GetRecordByPatient 根据患者id获取挂号记录
func GetRecordByPatient(patientId uint) []Registration {
	var records []Registration
	tx := global.GVA_DATABASE.Preload("Medic").Preload("Medic.Department").Preload("Patient").Where("patient_id=?", patientId).Find(&records)
	if tx.Error != nil {
		panic(tx.Error)
	}
	return records
}

// GetList 获取列表
func GetList(registration *Registration) []Registration {
	var records []Registration
	tx := global.GVA_DATABASE.Where(&registration).Find(&records)
	if tx.Error != nil {
		panic(tx.Error)
	}
	return records
}
