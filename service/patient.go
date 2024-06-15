package service

import (
	"service/global"
	. "service/model"
)

func InsertPatient(entity *Patient) {
	tx := global.GVA_DATABASE.Create(entity)
	if tx.Error != nil {
		panic(tx.Error)
	}
}
