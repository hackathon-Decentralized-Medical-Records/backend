package model

import "gorm.io/gorm"

// 授权
type Accredut struct {
	gorm.Model
	// 医生id
	MedicId uint `gorm:"type:int;not null" json:"medicId" `
	// 挂号id
	RegistrationId uint `gorm:"type:int;not null" json:"registrationId" `
	// 病例id
	CaseId uint `gorm:"type:int;not null" json:"caseId" `
}

type AccredutCase struct {
	Accredut `json:"accredit"`
	Case     Case `json:"case"`
}

func (Accredut) TableName() string {
	return "accredit"
}
