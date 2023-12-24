package dao

import (
	"cqu-backend/src/dao/db"
	"cqu-backend/src/dao/model"
	"github.com/go-xorm/xorm"
	"log"
)

type StudentDao struct {
	engine *xorm.Engine
}

func NewStudentDao() *StudentDao {
	return &StudentDao{engine: db.InstanceMaster()}
}

func (this *StudentDao) GetStudentByUnionId(unionId string) *model.Student {
	student := &model.Student{UnionId: unionId}
	get, err := this.engine.Get(student)
	if err != nil {
		log.Printf("[StudentDao GetStudentByUnionId] %+v\n", err.Error())
		return nil
	}
	if get {
		return student
	}
	return nil
}

func (this *StudentDao) Update(student *model.Student, cols ...string) {
	_, err := this.engine.ID(student.UnionId).MustCols(cols...).Update(student)
	if err != nil {
		log.Printf("[StudentDao Update] %+v\n", err.Error())
	}
}
