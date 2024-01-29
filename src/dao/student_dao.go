package dao

import (
	"cqu-backend/src/config/datasource"
	"cqu-backend/src/model"
	"github.com/go-xorm/xorm"
	"log"
)

type StudentDao struct {
	Engine *xorm.Engine
}

func NewStudentDao() *StudentDao {
	return &StudentDao{Engine: datasource.InstanceMaster()}
}

func (this *StudentDao) GetStudentByUnionId(unionId string) *model.Student {
	student := &model.Student{UnionId: unionId}
	get, err := this.Engine.Get(student)
	if err != nil {
		log.Printf("[StudentDao GetStudentByUnionId Error] %+v\n", err)
		return nil
	}
	if get {
		return student
	}
	return nil
}

func (this *StudentDao) Update(student *model.Student, cols ...string) {
	_, err := this.Engine.ID(student.UnionId).MustCols(cols...).Update(student)
	if err != nil {
		log.Printf("[StudentDao Update Error] %+v\n", err)
	}
}

func (this *StudentDao) Insert(student *model.Student) error {
	_, err := this.Engine.Insert(student)
	if err != nil {
		log.Printf("[StudentDao Insert Error] %+v\n", err)
		return err
	}
	return nil
}

func (this *StudentDao) Delete(student *model.Student) {
	_, err := this.Engine.Delete(student)
	if err != nil {
		log.Printf("[StudentDao Delete Error] %+v\n", err)
	}
}
