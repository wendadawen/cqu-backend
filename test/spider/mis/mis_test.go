package mis

import (
	"cqu-backend/src/spider"
	"cqu-backend/src/spider/mis"
	"cqu-backend/test"
	"cqu-backend/test/tool"
	"fmt"
	"testing"
)

func TestExamSchedule(t *testing.T) {
	account := spider.SpiderAccount{
		Account:  test.StudId2,
		Password: test.CasPwd2,
	}
	Mis, _ := mis.NewMisByCas(account)
	schedule, err := Mis.ExamSchedule()
	if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		println(tool.ToString(*schedule))
	}
}

func TestRank(t *testing.T) {
	account := spider.SpiderAccount{
		Account:  test.StudId2,
		Password: test.CasPwd2,
	}
	Mis, _ := mis.NewMisByCas(account)
	Rank, err := Mis.Rank()
	if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		println(tool.ToString(*Rank))
	}
}

func TestAllScore(t *testing.T) {
	account := spider.SpiderAccount{
		Account:  test.StudId2,
		Password: test.CasPwd2,
	}
	Mis, _ := mis.NewMisByCas(account)
	AllScore, err := Mis.AllScore()
	if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		println(tool.ToString(*AllScore))
	}
}

func TestCurrentScore(t *testing.T) {
	account := spider.SpiderAccount{
		Account:  test.StudId2,
		Password: test.CasPwd2,
	}
	Mis, _ := mis.NewMisByCas(account)
	CurrentScore, err := Mis.CurrentScore()
	if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		println(tool.ToString(*CurrentScore))
	}
}

func TestClassSchedule(t *testing.T) {
	account := spider.SpiderAccount{
		Account:  test.StudId2,
		Password: test.CasPwd2,
	}
	Mis, _ := mis.NewMisByCas(account)
	ClassSchedule, err := Mis.ClassSchedule()
	if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		println(tool.ToString(*ClassSchedule))
	}
}

func TestStudentInfo(t *testing.T) {
	account := spider.SpiderAccount{
		Account:  test.StudId2,
		Password: test.CasPwd2,
	}
	Mis, _ := mis.NewMisByCas(account)
	StudentInfo, err := Mis.StudentInfo()
	if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		println(tool.ToString(*StudentInfo))
	}
}
