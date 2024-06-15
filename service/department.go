package service

import (
	"service/global"
	. "service/model"
)

func GetDepartment() []Department {

	var departments []Department
	global.GVA_DATABASE.Find(&departments)
	return departments
}
