package service

import (
	"service/global"
	. "service/model"
)

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
