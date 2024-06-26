package mysql

import "gorm.io/gorm"

// 病例
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

// 新增病例
func InsertCase(entity *Case) error {
	tx := db.Create(entity)
	return tx.Error
}

// 查询病例
func GetCase(patientId uint) ([]Case, error) {
	var cases []Case
	tx := db.Preload("Patient").Where("patient_id = ?", patientId).Find(&cases)
	return cases, tx.Error
}

// 根据id查询病例
func GetCaseById(id []uint) ([]Case, error) {
	var c []Case
	tx := db.Preload("Patient").Where("id in (?)", id).First(&c)
	return c, tx.Error
}
