package mis

import (
	"bytes"
	"cqu-backend/src/bo"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"regexp"
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
