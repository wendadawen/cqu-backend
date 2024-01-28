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

func TestRank(t *testing.T) {
	account := spider.SpiderAccount{
		Account:  test.StudId1,
		Password: test.CasPwd1,
	}
	My, _ := my.NewMyByCas(account)
	Rank, err := My.Rank()
	if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		println(tool.ToString(*Rank))
	}
}

func TestAllScore(t *testing.T) {
	account := spider.SpiderAccount{
		Account:  test.StudId1,
		Password: test.CasPwd1,
	}
	My, _ := my.NewMyByCas(account)
	AllScore, err := My.AllScore()
	if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		println(tool.ToString(*AllScore))
	}
}

func TestCurrentScore(t *testing.T) {
	account := spider.SpiderAccount{
		Account:  test.StudId1,
		Password: test.CasPwd1,
	}
	My, _ := my.NewMyByCas(account)
	CurrentScore, err := My.CurrentScore()
	if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		println(tool.ToString(*CurrentScore))
	}
}
