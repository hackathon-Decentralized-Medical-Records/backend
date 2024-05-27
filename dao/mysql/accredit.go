package mysql

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

// InsertAccredit 新增授权
func InsertAccredit(entity *Accredut) error {
	tx := db.Create(entity)
	return tx.Error
}

// GetAccredit 获取授权
func GetAccredit(entity *Accredut) ([]Case, error) {
	var accredits []Accredut

	query := db.Model(&entity)
	if entity.MedicId != 0 {
		query = query.Where("medic_id=?", entity.MedicId)
	}

	if entity.RegistrationId != 0 {
		query = query.Where("registration_id=?", entity.RegistrationId)
	}

	if entity.CaseId != 0 {
		query = query.Where("case_id = ?", entity.CaseId)
	}

	tx := query.Find(&accredits)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var caseIds []uint
	for i := range accredits {
		caseIds = append(caseIds, accredits[i].CaseId)
	}

	allCases, err := GetCaseById(caseIds)
	if err != nil {
		return nil, err
	}

	return allCases, nil
}
