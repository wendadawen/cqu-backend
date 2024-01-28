package mis

import (
	"bytes"
	"cqu-backend/src/bo"
	"cqu-backend/src/config"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cast"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// 网页是GBK编码，转化为utf-8编码
func tran2Utf8(res string) string {
	gbkBytes := []byte(res)
	reader := transform.NewReader(bytes.NewReader(gbkBytes), simplifiedchinese.GBK.NewDecoder())
	utf8Bytes, _ := ioutil.ReadAll(reader)
	return string(utf8Bytes)
}

func extractExam(res string) *bo.ExamScheduleBo {
	examList := make([]bo.ExamBo, 0)
	res = tran2Utf8(res)
	document, _ := goquery.NewDocumentFromReader(strings.NewReader(res))
	document.Find("table tr").Next().Next().Each(func(i int, tr *goquery.Selection) {
		item := tr.Find("td").Map(func(idx int, td *goquery.Selection) string {
			return strings.TrimSpace(td.Text())
		})
		if len(item) == 5 {
			// 对时间进行处理
			compile := regexp.MustCompile("\\(.*?\\)")
			shangxiawu := regexp.MustCompile(" {2}(.*?) ")
			t := compile.ReplaceAllString(item[3], "")
			t = shangxiawu.ReplaceAllString(t, " ")
			t = strings.ReplaceAll(t, " －－ ", "-")

			examList = append(examList, bo.ExamBo{
				ExamId:       item[0],
				ExamTitle:    item[1],
				ExamTime:     item[2] + " " + t,
				ExamLocation: item[4],
			})
		}
	})
	return &examList
}

func extractStudentInfo(res string) *bo.StudentInfoBo {
	res = tran2Utf8(res)
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(res))
	student := new(bo.StudentInfoBo)
	student.Type = bo.GraduateStudent
	selection := doc.Find("table.mode19 tr td.mode5")
	student.Uid = strings.TrimSpace(selection.Eq(0).Text())
	student.StudentId = strings.TrimSpace(selection.Eq(1).Text())
	student.DeptName = strings.TrimSpace(selection.Eq(4).Text())
	student.MajorName = strings.TrimSpace(selection.Eq(5).Text())
	student.Grade = strings.TrimSpace(selection.Eq(8).Text())
	selection = doc.Find("table#tab1 table td.mode5")
	student.StudentName = strings.TrimSpace(selection.Eq(1).Text())
	student.IdNumber = strings.TrimSpace(selection.Eq(12).Text())
	return student
}

func extractClassSchedule(res string) *bo.ClassScheduleBo {
	res = tran2Utf8(res)
	classSchedule := make([]bo.ClassInfo, 0)
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(res))
	doc.Find("table tr").Next().Next().Each(func(i int, tr *goquery.Selection) {
		tr.Find("td").Next().Each(func(j int, td *goquery.Selection) {
			html, _ := td.Html()
			html = strings.ReplaceAll(html, `<font color="red">`, "")
			html = strings.ReplaceAll(html, `</font>`, "")
			for _, item := range strings.Split(html, "<br/><br/>") {
				if len(strings.TrimSpace(item)) == 0 {
					continue
				}
				split := strings.Split(item, "<br/>")
				teachClass := "" // 教学班
				id := ""
				title := ""
				week := ""
				teacher := ""
				classroom := "null"
				section := ""
				paltform := "暂未安排"
				meetingID := "暂未安排"
				classQQ := "暂未安排"
				for _, s := range split {
					name := strings.Split(s, "：")
					if name[0] == "班号" {
						teachClass = name[1]
					} else if name[0] == "代码" {
						id = name[1]
					} else if name[0] == "名称" {
						title = name[1]
					} else if name[0] == "周次" {
						week = name[1]
					} else if name[0] == "教师" {
						teacher = name[1]
					} else if name[0] == "教室" {
						classroom = name[1]
					} else if name[0] == "节次" {
						section = name[1]
					} else if name[0] == "平台" {
						paltform = name[1]
					} else if name[0] == "会议ID/房间ID" {
						meetingID = strings.Join(name[1:], ":")
					} else if name[0] == "班级QQ群" {
						classQQ = strings.Join(name[1:], ":")
					}
				}
				start := cast.ToInt(strings.Split(section, "-")[0])
				end := cast.ToInt(strings.Split(section, "-")[1])
				classSchedule = append(classSchedule, bo.ClassInfo{
					Id:         id,
					Title:      title,
					Weeks:      parseWeeks(week),
					Num:        end - start + 1,
					Day:        j,
					Room:       classroom,
					TeachClass: teachClass,
					Start:      start,
					Teacher:    teacher,
					More:       week + section,
					Meeting:    paltform + meetingID + "\n班级QQ群:" + classQQ,
				})
			}
		})
	})
	return &classSchedule
}

func parseWeeks(weekStr string) []int {
	week := make([]int, 0)
	weekStr = SubTill(weekStr, "周")
	for _, str := range strings.Split(weekStr, " ") {
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
	return week
}

func SubTill(str string, endStr string) string {
	rs := []rune(str)
	length := len(rs)
	end := UnicodeIndex(str, endStr)
	if end < 0 || end > length {
		return ""
	}
	return string(rs[:end])
}

func UnicodeIndex(str, substr string) int {
	result := strings.Index(str, substr)
	if result >= 0 {
		prefix := []byte(str)[0:result]
		rs := []rune(string(prefix))
		result = len(rs)
	}
	return result
}

var toNum = map[rune]string{'0': "零", '1': "一", '2': "二", '3': "三", '4': "四", '5': "五", '6': "六", '7': "七", '8': "八", '9': "九"}

func extractAllScore(res string) *bo.MyScoreResultBo {
	res = tran2Utf8(res)
	scores := map[string]bo.MyScoreList{} //一个学期的放在一起
	total := bo.MyScoreList{}             //用于计算全部绩点
	totalScore := []bo.MyScoreListBo{}    //用于存放每个学期的成绩和绩点
	var session string
	totalCredits := map[string]float64{}
	document, _ := goquery.NewDocumentFromReader(strings.NewReader(res))
	document.Find("table#score_sheet tr").Next().Next().Each(func(i int, tr *goquery.Selection) {
		item := tr.Find("td").Map(func(idx int, td *goquery.Selection) string {
			return strings.TrimSpace(td.Text())
		})
		//len(item)==2时，item[2]为综合加权平均成绩，后续可以加上，但需要和本科生成绩统一格式
		if len(item) == 5 {
			s := bo.MyScoreBo{
				StudyNature:        item[0],
				Course:             item[1],
				Credits:            item[3], //未知学年的学分为-
				EffectiveScoreShow: item[4],
				EffectiveScore:     item[4], //都返回
			}
			if v, ok := totalCredits[item[2]]; ok {
				totalCredits[item[2]] = v + cast.ToFloat64(s.Credits)
			} else {
				totalCredits[item[2]] = cast.ToFloat64(s.Credits)
			}
			if cast.ToFloat64(s.Credits) != 0 {
				total = append(total, s) //用于计算总绩点,其中学分项不能缺失
			}
			if list, ok := scores[item[2]]; ok { //每个学期的放一个集合
				scores[item[2]] = append(list, s)
			} else {
				scores[item[2]] = bo.MyScoreList{s}
			}
		}
	})

	for k, _ := range scores {
		r := []rune(k)
		if len(r) != 3 {
			session = "未知学年"
		} else {
			session = fmt.Sprintf("第%s学年%c", toNum[r[0]], r[2])
		}
		totalScore = append(totalScore, bo.MyScoreListBo{
			Session:      session,
			TotalCredits: strconv.FormatFloat(totalCredits[k], 'f', 1, 64),
			GpaFour:      cast.ToString(scores[k].Gpa(bo.FourGpa)),
			GpaFive:      cast.ToString(scores[k].Gpa(bo.FiveGpa)),
			Scores:       scores[k],
		})
	}
	return &bo.MyScoreResultBo{
		TotalScore:   totalScore,
		TotalGpaFour: cast.ToString(total.Gpa(bo.FourGpa)),
		TotalGpaFive: cast.ToString(total.Gpa(bo.FiveGpa)),
	}
}

func extractCurrentScore(res string) *bo.MyScoreListBo {
	res = tran2Utf8(res)
	var session = config.CquConfig.TermCurrentMis
	var scores bo.MyScoreList
	document, _ := goquery.NewDocumentFromReader(strings.NewReader(res))
	document.Find("table#score_sheet tr").Next().Next().Each(func(i int, tr *goquery.Selection) {
		item := tr.Find("td").Map(func(idx int, td *goquery.Selection) string {
			return strings.TrimSpace(td.Text())
		})
		if len(item) == 5 && item[2] == session { //过滤最近学期出的成绩
			s := bo.MyScoreBo{
				StudyNature:    item[0],
				Course:         item[1],
				Credits:        item[3], //未知学年的学分为-
				EffectiveScore: item[4], //都返回
			}
			scores = append(scores, s)
		}
	})
	r := []rune(session)
	return &bo.MyScoreListBo{
		Session: fmt.Sprintf("第%s学年%c", toNum[r[0]], r[2]),
		Scores:  scores,
		GpaFour: cast.ToString(scores.Gpa(bo.FourGpa)),
		GpaFive: cast.ToString(scores.Gpa(bo.FiveGpa)),
	}
}
