package models

type User struct {
	UserName string
	PassWord string
	Role     string
}

// 模拟数据集合
var users = []User{
	{UserName: "medicTest", PassWord: "123456", Role: "medic"},
	{UserName: "patientTest", PassWord: "123456", Role: "patient"},
}

// 登录校验
func LoginCheck(userName string, passWord string) bool {
	var ok bool
	for _, user := range users {
		if user.UserName == userName && user.PassWord == passWord {
			ok = true
		}
	}
	return ok
}
