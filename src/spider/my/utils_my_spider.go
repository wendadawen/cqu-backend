package my

import (
	"cqu-backend/src/bo"
	"cqu-backend/src/config"
	"cqu-backend/src/dao/model"
	"fmt"
	"github.com/spf13/cast"
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

func exactRank(json string) *model.Rank {
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

func exactAllScore(json string) *bo.MyScoreResultBo {
	data := gjson.Get(json, "data")
	totalScores := []bo.MyScoreListBo{}               //每个学期的成绩
	totalGpaScores := bo.MyScoreList{}                //用于计算综合绩点
	totalMinorScores := bo.MyScoreList{}              //用于计算辅修课程绩点
	data.ForEach(func(key, value gjson.Result) bool { //某一学期成绩
		scores := bo.MyScoreList{}      // 学期成绩表
		minorScores := bo.MyScoreList{} //用于计算辅修课程绩点
		gpaScores := bo.MyScoreList{}   //用于计算学期绩点
		courses := value.Get("stuScoreHomePgVoS").Array()
		for _, course := range courses { //成绩列表
			s := bo.MyScoreBo{
				Course:             course.Get("courseName").String(),
				Credits:            course.Get("courseCredit").String(),
				Code:               course.Get("courseCode").String(),
				Instructor:         course.Get("instructorName").String(),
				EffectiveScore:     course.Get("effectiveScore").String(),
				EffectiveScoreShow: course.Get("effectiveScoreShow").String(), //合格、良、优
				CourseNature:       course.Get("courseNature").String(),       //必修、选修
				StudyNature:        course.Get("studyNature").String(),        //初修
			}
			scores = append(scores, s)
			if s.Course == "大学英语(国家四级)" || s.Course == "大学英语(国家六级)" || s.Code == "5200080" || s.Code == "5200081" {
				continue
			} else if s.CourseNature == "二专" || s.CourseNature == "辅修" { //计入辅修表
				minorScores = append(minorScores, s)
				totalMinorScores = append(totalMinorScores, s)
			} else { // 计入绩点表
				gpaScores = append(gpaScores, s)
				totalGpaScores = append(totalGpaScores, s)
			}
		}
		totalScores = append(totalScores, bo.MyScoreListBo{
			Session:      key.String(),                      //学期
			TotalCredits: value.Get("totalCredit").String(), //总学分
			Scores:       scores,                            //成绩表
			GpaFour:      cast.ToString(gpaScores.Gpa(bo.FourGpa)),
			GpaFive:      cast.ToString(gpaScores.Gpa(bo.FiveGpa)),
			MinorGpaFour: cast.ToString(minorScores.Gpa(bo.FourGpa)),
			MinorGpaFive: cast.ToString(minorScores.Gpa(bo.FiveGpa)),
		})
		return true
	})
	return &bo.MyScoreResultBo{
		totalScores,
		cast.ToString(totalGpaScores.Gpa(bo.FourGpa)),
		cast.ToString(totalGpaScores.Gpa(bo.FiveGpa)),
		cast.ToString(totalMinorScores.Gpa(bo.FourGpa)),
		cast.ToString(totalMinorScores.Gpa(bo.FiveGpa)),
	}
}

func exactCurrentScore(json string) *bo.MyScoreListBo {
	data := gjson.Get(json, "data")
	var myScoreList bo.MyScoreList
	data.ForEach(func(key, value gjson.Result) bool { //某一学期成绩
		if key.String() == config.CquConfig.TermCurrentMy {
			courses := value.Get("stuScoreHomePgVoS").Array()
			for _, course := range courses { //成绩列表
				s := bo.MyScoreBo{
					Course:             course.Get("courseName").String(),
					Credits:            course.Get("courseCredit").String(),
					Code:               course.Get("courseCode").String(),
					Instructor:         course.Get("instructorName").String(),
					EffectiveScore:     course.Get("effectiveScore").String(),
					EffectiveScoreShow: course.Get("effectiveScoreShow").String(), //合格、良、优
					CourseNature:       course.Get("courseNature").String(),       //必修、选修
					StudyNature:        course.Get("studyNature").String(),        //初修
				}
				myScoreList = append(myScoreList, s)
			}
			return true
		}
		return true
	})
	return &bo.MyScoreListBo{
		Session: config.CquConfig.TermCurrentMy,
		Scores:  myScoreList,
		GpaFour: cast.ToString(myScoreList.Gpa(bo.FourGpa)),
		GpaFive: cast.ToString(myScoreList.Gpa(bo.FiveGpa)),
	}
}
