package my

import (
	"cqu-backend/src/bo"
	"cqu-backend/src/config"
	"cqu-backend/src/model"
	"fmt"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"regexp"
	"strconv"
	"strings"
)

var days = map[string]string{"1": "一", "2": "二", "3": "三", "4": "四", "5": "五", "6": "六", "7": "日"}

func extractExam(json string) *bo.ExamScheduleBo {
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

func extractRank(json string) *model.Rank {
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

func extractAllScore(json string) *bo.MyScoreResultBo {
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

func extractCurrentScore(json string) *bo.MyScoreListBo {
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

func extractClassSchedule(json string) *bo.ClassScheduleBo {
	classes := gjson.Get(json, "data").Array()
	schedule := make(bo.ClassScheduleBo, 0)
	for _, class := range classes {
		title := class.Get("courseName").String()
		id := class.Get("courseCode").String()
		teachClass := class.Get("classNbr").String()
		classTimetableVOList := class.Get("classTimetableVOList").Array()
		teachers := class.Get("instList").Array()
		teacher_t := ""
		for i, x := range teachers {
			if i == len(teachers)-1 {
				teacher_t += x.String()
			} else {
				teacher_t += x.String() + "，"
			}
		}
		teacher := regexp.MustCompile(`\[[0-9]+]`).ReplaceAllString(teacher_t, "")
		for _, dayClass := range classTimetableVOList {
			room := dayClass.Get("roomName").String()
			day := dayClass.Get("weekDayFormat").String()
			teachingWeekFormat := dayClass.Get("teachingWeekFormat").String()
			period := dayClass.Get("periodFormat").String()
			startSection, num := 0, 0
			if strings.Contains(period, "-") {
				split := strings.Split(period, "-")
				startSection, _ = strconv.Atoi(split[0])
				endSection, _ := strconv.Atoi(split[1])
				num = endSection - startSection + 1
			} else if period == "" { // 没有分配时间的
				startSection, num = 0, 0
			} else { // 只有一节课的
				startSection, _ = strconv.Atoi(period)
				num = 1
			}
			info := bo.ClassInfo{
				Title:      title,
				Id:         id,
				TeachClass: teachClass,
				Day:        chineseToNumber(day),
				Weeks:      parseWeek(teachingWeekFormat),
				Room:       room,
				Start:      startSection,
				Num:        num,
				Teacher:    teacher,
				More:       teachingWeekFormat,
				Content:    "",
				CampusId:   room,
			}
			schedule = append(schedule, info)
		}
	}
	return &schedule
}

func chineseToNumber(day string) int {
	switch day {
	case "一":
		return 0
	case "二":
		return 1
	case "三":
		return 2
	case "四":
		return 3
	case "五":
		return 4
	case "六":
		return 5
	case "日":
		return 6
	default:
		return -1 // 非法输入，返回-1表示错误
	}
}

func parseWeek(weekStr string) (week []int) {
	week = make([]int, 0)
	for _, str := range strings.Split(weekStr, ",") {
		if strings.Index(str, "-") != -1 {
			split := strings.Split(str, "-")
			start, end := cast.ToInt(split[0]), cast.ToInt(split[1])
			for i := start; i <= end; i++ {
				week = append(week, i)
			}
		} else {
			w, _ := strconv.Atoi(str)
			week = append(week, w)
		}
	}
	return
}

func extractStudentInfo(json string) *bo.StudentInfoBo {
	data := gjson.Get(json, "data")
	return &bo.StudentInfoBo{
		Type:                 bo.UndergraduateStudent,
		StudentId:            data.Get("studentId").String(),
		StudentName:          data.Get("studentName").String(),
		Gender:               data.Get("gender").String(),
		Grade:                data.Get("grade").String(),
		DeptName:             data.Get("deptName").String(),
		MajorName:            data.Get("majorName").String(),
		ClassName:            data.Get("className").String(),
		IdNumber:             data.Get("idNumber").String(),
		PoliticalStatus:      data.Get("politicalStatus").String(),
		Nationality:          data.Get("nationality").String(),
		Phone:                data.Get("phone").String(),
		Email:                data.Get("email").String(),
		AuthId:               data.Get("authId").String(),
		Birthday:             data.Get("birthday").String(),
		HomeAddress:          data.Get("homeAddress").String(),
		Gpa:                  data.Get("gpa").String(),
		MajorRanking:         data.Get("majorRanking").String(),
		Duration:             data.Get("duration").String(),
		ObtainSchoolRollTime: data.Get("obtainSchoolRollTime").String(),
		EnrollmentTime:       data.Get("enrollmentTime").String(),
		DepartureTime:        data.Get("departureTime").String(),
		StuSourceRegion:      data.Get("stuSourceRegion").String(),
		StuSourceUnit:        data.Get("stuSourceUnit").String(),
	}
}
