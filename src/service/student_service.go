package service

import (
	"cqu-backend/src/dao"
	"cqu-backend/src/dao/model"
)

type StudentService struct {
	StuDao *dao.StudentDao
}

func NewStudentService() *StudentService {
	return &StudentService{StuDao: dao.NewStudentDao()}
}

// GetStudent 获取学生信息
func (this *StudentService) GetStudent(unionId string) *model.Student {
	return this.StuDao.GetStudentByUnionId(unionId)
}

// Update 更新学生表的信息
func (this *StudentService) Update(student *model.Student) {
	if student.Uid == "" {
		student.Uid = student.StuId
	}
	this.StuDao.Update(student)
}
