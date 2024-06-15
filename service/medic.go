package service

import (
	"gorm.io/gorm"
	"service/global"
)

// Medic 医生
type Medic struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);not null"  binding:"omitempty" form:"name" json:"name" `
	// 身份
	Profession string `gorm:"type:varchar(100);not null"  binding:"omitempty" form:"profession" json:"profession" `
	// 开始工作时间
	WorkTime string `gorm:"type:datetime;not null"  binding:"omitempty" form:"workTime" json:"workTime" `
	// 结束工作时间
	EndTime string `gorm:"type:datetime;not null"  binding:"omitempty"  form:"endTime" json:"endTime" `
	// 价格
	Price float64 `gorm:"type:float;not null"  binding:"omitempty" form:"price" json:"price" `
	// 简介
	Introduce    string `gorm:"type:varchar(1000);not null"  binding:"omitempty" json:"introduce" `
	DepartmentID uint   `gorm:"not null"   binding:"omitempty"  json:"department_id"` // 外键字段
	// 科室
	Department Department `gorm:"foreignKey:DepartmentID" binding:"omitempty" json:"department"`
	UserID     uint       `gorm:"not null" binding:"omitempty" json:"user_id"` // 外键字段
	// user
	User User   `gorm:"foreignKey:UserID" binding:"omitempty" json:"user"`
	Sort string `json:"sort" binding:"omitempty" `
}

func (Medic) TableName() string {
	return "medic"
}

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
