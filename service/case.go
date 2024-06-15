package service

import (
	"service/global"
	. "service/model"
)

// 新增病例
func InsertCase(entity *Case) error {
	tx := global.GVA_DATABASE.Create(entity)
	return tx.Error
}

// 查询病例
func GetCase(patientId uint) ([]Case, error) {
	var cases []Case
	tx := global.GVA_DATABASE.Preload("Patient").Where("patient_id = ?", patientId).Find(&cases)
	return cases, tx.Error
}

// 根据id查询病例
func GetCaseById(id []uint) ([]Case, error) {
	var c []Case
	tx := global.GVA_DATABASE.Preload("Patient").Where("id in (?)", id).First(&c)
	return c, tx.Error
}
