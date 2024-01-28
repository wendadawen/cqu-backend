package service

import (
	"cqu-backend/src/bo"
	"cqu-backend/src/dao"
	"cqu-backend/src/model"
	"cqu-backend/src/object"
	"cqu-backend/src/spider"
	"cqu-backend/src/spider/mis"
	"cqu-backend/src/spider/my"
	"cqu-backend/src/tool"
	"log"
)

//type StudentService interface {
//	GetStudent(UnionId string) *model.Student
//	GetAccount(UnionId string) *model.Student
//	BindCas(UnionId, StuId, CasPwd string) (*model.Student, error)
//	Update(student *model.Student)
//}

type StudentService struct {
	StudentDao *dao.StudentDao
}

func NewStudentService() *StudentService {
	return &StudentService{StudentDao: dao.NewStudentDao()}
}

// GetStudent 获取学生信息
func (this *StudentService) GetStudent(UnionId string) *model.Student {
	return this.StudentDao.GetStudentByUnionId(UnionId)
}

func (this *StudentService) BindCas(UnionId, StuId, CasPwd string) (*model.Student, error) {
	if student := this.GetStudent(UnionId); student != nil {
		return student, nil
	}
	account := spider.SpiderAccount{
		Account:  StuId,
		Password: CasPwd,
	}
	info := &bo.StudentInfoBo{}
	if tool.CheckGraduate(StuId, CasPwd) {
		Mis, err := mis.NewMisByCas(account)
		if err != nil {
			log.Printf("[StudentService BindCas Error] Account=%s\n", StuId)
			return nil, err
		}
		info, err = Mis.StudentInfo()
		if err != nil {
			log.Printf("[StudentService BindCas Error] Account=%s\n", StuId)
			return nil, err
		}
	} else {
		My, err := my.NewMyByCas(account)
		if err != nil {
			log.Printf("[StudentService BindCas Error] Account=%s\n", StuId)
			return nil, err
		}
		info, err = My.StudentInfo()
		if err != nil {
			log.Printf("[StudentService BindCas Error] Account=%s\n", StuId)
			return nil, err
		}
	}
	if len(info.StudentId) == 0 {
		return nil, object.CasIdNoError
	}
	student := &model.Student{
		UnionId: UnionId,
		CasPwd:  CasPwd,
		StuId:   info.StudentId,
		Uid:     info.Uid,
		Name:    info.StudentName,
		College: info.DeptName,
		Class:   info.ClassName,
		Major:   info.MajorName,
	}
	err := this.StudentDao.Insert(student)
	if err != nil {
		log.Printf("[StudentService BindCas Error] Account=%s\n", StuId)
		return nil, err
	}
	return student, nil
}

// Update 更新学生表的信息
func (this *StudentService) Update(student *model.Student) {
	if student.Uid == "" {
		student.Uid = student.StuId
	}
	this.StudentDao.Update(student)
}

func (this *StudentService) GetAccount(UnionId string) *model.Student {
	student := this.GetStudent(UnionId)
	if student == nil {
		return nil
	}
	Trans2Account(student)
	return student
}

func Trans2Account(student *model.Student) {
	checkIsBound := func(pwd *string) {
		if len(*pwd) > 0 {
			*pwd = "true"
		} else {
			*pwd = "false"
		}
	}
	checkIsBound(&student.CasPwd)
	checkIsBound(&student.JwcPwd)
	checkIsBound(&student.LibPwd)
	checkIsBound(&student.CardPwd)
}
