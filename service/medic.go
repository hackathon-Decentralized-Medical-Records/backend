package service

import (
	"gorm.io/gorm"
	"service/global"
	. "service/model"
)

func InsertMedic(m *Medic) (int, string) {
	global.GVA_DATABASE.Create(m)
	return 200, "注册成功"
}

func GetMedicList(m *Medic) []Medic {
	var medicList []Medic

	var order string

	if m.Sort != "" {
		order = m.Sort
	}

	query := global.GVA_DATABASE.Model(&Medic{})
	if m.Name != "" {
		query = query.Where("name LIKE ?", "%"+m.Name+"%")
	}

	if m.Profession != "" {
		query = query.Where("profession = ?", m.Profession)
	}

	if m.DepartmentID != 0 {
		query = query.Where("department_id = ?", m.DepartmentID)
	}

	if m.Price != 0 {
		query = query.Where("price >= ? ", m.Price)
	}

	if m.WorkTime != "" {
		query = query.Where("work_time between ? and ? ", m.WorkTime, m.EndTime)
	}

	query.Preload("Department").Preload("User", func(db *gorm.DB) *gorm.DB {
		return global.GVA_DATABASE.Select("id, user_name, email, role")
	}).Find(&medicList).Order(order)
	return medicList
}
