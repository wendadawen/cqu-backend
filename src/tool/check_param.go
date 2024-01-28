package tool

import (
	"cqu-backend/src/model"
	"cqu-backend/src/object"
)

// CheckIsEmpty 检查参数是否为空
func CheckIsEmpty(s ...string) bool {
	for _, str := range s {
		if str == "" {
			return true
		}
	}
	return false
}

// CheckCasAccount 检查学生统一认证信息绑定情况
func CheckCasAccount(student *model.Student) error {
	if student == nil || student.StuId == "" || student.CasPwd == "" {
		return object.CasNotBound // 返回统一认证未绑定的认证
	}
	return nil
}

// CheckCardAccount 检查一卡通是否绑定
func CheckCardAccount(student *model.Student) error {
	if err := CheckCasAccount(student); err != nil {
		return err
	}
	if student == nil || student.CardPwd == "" {
		return object.CardNotBoundError // 一卡通没有绑定
	}
	return nil
}

// CheckRoomAccount 检查房间是否绑定
func CheckRoomAccount(student *model.Student) error {
	if err := CheckCasAccount(student); err != nil {
		return err
	}
	if student == nil || student.Room == "" {
		return object.RoomNotBoundError // 房间没有绑定
	}
	return nil
}
