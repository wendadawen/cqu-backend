package my

import (
	"cqu-backend/src/bo"
	"cqu-backend/src/dao/model"
	"fmt"
	"github.com/tidwall/gjson"
)

var days = map[string]string{"1": "一", "2": "二", "3": "三", "4": "四", "5": "五", "6": "六", "7": "日"}

func exactExams(json string) *bo.ExamScheduleBo {
	parse := gjson.Parse(json)
	contents := parse.Get("data.content").Array()
	exams := bo.ExamScheduleBo{}
	for _, content := range contents {
		time := fmt.Sprintf("%s[ %s周 星期%s ] %s ~ %s",
			content.Get("examDate"),
			content.Get("week"),
			days[content.Get("weekDay").String()],
			content.Get("startTime").String(),
			content.Get("endTime").String())
		exam := bo.ExamBo{
			ExamTitle:    content.Get("courseName").String(),
			ExamTime:     time,
			ExamSeat:     content.Get("seatNum").String(),
			ExamLocation: content.Get("roomName").String(),
		}
		exams = append(exams, exam)
	}
	return &exams
}

func exacRank(json string) *model.Rank {
	data := gjson.Get(json, "data")
	rank := &model.Rank{
		Gpa:          data.Get("gpa").String(),
		WeightAvg:    data.Get("weightedAvg").String(),
		GradeRanking: data.Get("gradeRanking").String(),
		MajorRanking: data.Get("majorRanking").String(),
		ClassRanking: data.Get("classRanking").String(),
	}
	return rank
}
