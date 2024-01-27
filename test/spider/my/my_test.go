package my

import (
	"cqu-backend/src/spider"
	"cqu-backend/src/spider/my"
	"cqu-backend/test"
	"cqu-backend/test/tool"
	"fmt"
	"testing"
)

func TestExamSchedule(t *testing.T) {
	account := spider.SpiderAccount{
		Account:  test.StudId1,
		Password: test.CasPwd1,
	}
	My, _ := my.NewMyByCas(account)
	schedule, err := My.ExamSchedule()
	if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		println(tool.ToString(*schedule))
	}
}
