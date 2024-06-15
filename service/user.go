package service

import (
	"gorm.io/gorm"
	"service/global"
	"service/utils/httputils"
)

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(100);not null" json:"username" `
	PassWord string `gorm:"type:varchar(65);not null" json:"password" `
	Email    string `gorm:"type:varchar(100);not null" json:"email" `
	Role     int    `gorm:"type:int" json:"role"`
	Address  string `gorm:"type:varchar(42);column:address" json:"address"`
}

type Response struct {
	Data interface{} `json:"data"`
}

func (User) TableName() string {
	return "user"
}

// 校验用户是否存在
func CheckUserByEmail(email string) int {
	var user User
	global.GVA_DATABASE.Where("email = ?", email).First(&user)
	if user.ID > 0 {
		return httputils.StatusConflict
	}
	return httputils.StatusOK

}

// 校验登录信息
func GetUserByNameAndPwd(loginId, passWord string) (int, string, Response) {
	var user User
	var response Response
	global.GVA_DATABASE.Where("email = ? AND pass_word = ?", loginId, passWord).First(&user)
	if user.ID <= 0 {
		return httputils.StatusUnauthorized, "登录失败", response
	}
	user.PassWord = ""

	// 如果是医生则查询医生信息

	if user.Role == 0 {
		//只查一条数据
		var medic Medic
		tx := global.GVA_DATABASE.Table(Medic{}.TableName()).Where("user_id = ?", user.ID).Take(&medic)
		if tx.Error != nil {
			return httputils.StatusInternalServerError, "登录失败", response
		}
		response.Data = medic
		return httputils.StatusOK, "登录成功", response
	}

	if user.Role == 1 {
		var patient Patient
		tx := global.GVA_DATABASE.Table(Patient{}.TableName()).Where("user_id = ?", user.ID).Take(&patient)
		if tx.Error != nil {
			return httputils.StatusInternalServerError, "登录失败", response
		}
		response.Data = patient
		return httputils.StatusOK, "登录成功", response
	}
	return httputils.StatusInternalServerError, "登录失败", response

}

// 注册用户信息
func RegisterUser(user *User) (int, string) {
	global.GVA_DATABASE.Create(user)
	return httputils.StatusOK, "注册成功"
}
