package service

import (
	"cqu-backend/src/bo"
	"cqu-backend/src/config/setting"
	"cqu-backend/src/dao"
	"cqu-backend/src/model"
	"cqu-backend/src/object"
	"cqu-backend/src/spider"
	"cqu-backend/src/spider/mis"
	"cqu-backend/src/spider/my"
	"cqu-backend/src/tool"
	"github.com/spf13/cast"
	"log"
)

type StudentService struct {
	StudentDao *dao.StudentDao
}

func NewStudentService() *StudentService {
	return &StudentService{StudentDao: dao.NewStudentDao()}
}

func (this *StudentService) GetStudent(UnionId string) *model.Student {
	return this.StudentDao.GetStudentByUnionId(UnionId)
}

func (this *StudentService) BindCas(UnionId, StuId, CasPwd string) (*model.Student, error) {
	student1 := &model.Student{}
	if student1 = this.GetStudent(UnionId); student1 != nil && len(student1.CasPwd) != 0 {
		return student1, nil
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
	if student1 == nil {
		err := this.StudentDao.Insert(student)
		if err != nil {
			log.Printf("[StudentService BindCas Error] Account=%s\n", StuId)
			return nil, err
		}
	} else {
		this.StudentDao.Update(student, "cas_pwd", "uid", "name", "college", "class", "major")
	}
	return student, nil
}

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

func (this *StudentService) ClearCas(UnionId string) {
	student := &model.Student{UnionId: UnionId, CasPwd: ""}
	this.StudentDao.Update(student, "cas_pwd")
}

func (this *StudentService) ClearRoom(UnionId string) {
	student := &model.Student{UnionId: UnionId, Room: ""}
	this.StudentDao.Update(student, "room")
}

func (this *StudentService) ClearAll(UnionId string) {
	student := &model.Student{UnionId: UnionId}
	this.StudentDao.Delete(student)
}

func (this *StudentService) BindRoom(UnionId string, Room string, Campus string, Dom string, RoomNum string) (*bo.ElectricCharge, error) {
	student := this.GetStudent(UnionId)
	if err := tool.CheckCasAccount(student); err != nil {
		return nil, err
	}
	electricFee, err := NewCardService().ElectricByCas(student.StuId, student.CasPwd, Room)
	if err != nil {
		log.Printf("[StudentService BindRoom Error] Account=%s\n", student.StuId)
		return nil, err
	}
	student.ElectricFee = electricFee.Balance
	student.Room = Room
	student.Campus = Campus
	student.Dom = Dom
	student.RoomNum = RoomNum
	this.Update(student)
	return electricFee, nil
}

func (this *StudentService) GetTerm() *bo.TermInfoBo {
	term := setting.CquConfig.GetString("term_name")
	startTime := setting.CquConfig.GetString("term_start_date")
	endTime := setting.CquConfig.GetString("term_end_date")
	currentWeek := tool.CurrentWeek()
	return &bo.TermInfoBo{Term: term, StartTime: startTime, EndTime: endTime, CurrentWeek: cast.ToString(currentWeek)}
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
