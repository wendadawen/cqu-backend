package bo

import (
	"github.com/spf13/cast"
	"math"
)

type MyScoreResultBo struct {
	TotalScore   []MyScoreListBo //全部成绩
	TotalGpaFour string          //  综合绩点
	TotalGpaFive string          //辅修综合绩点
	MinorGpaFour string
	MinorGpaFive string
}

type MyScoreListBo struct {
	Session      string //学期
	TotalCredits string //总学分
	GpaFour      string //绩点
	GpaFive      string
	MinorGpaFour string //辅修绩点
	MinorGpaFive string
	Scores       MyScoreList //各科成绩
}
type MyScoreList []MyScoreBo
type MyScoreBo struct {
	Course             string //课程名
	Credits            string //学分
	Code               string //课程编号
	Instructor         string //教师名
	EffectiveScoreShow string // 合格、良、优
	EffectiveScore     string //分数
	CourseNature       string `json:",omitempty"` // 类别 必修、选修
	StudyNature        string `json:",omitempty"` // 先修
	Remark             string `json:"-"`          // 备注
}

const (
	FourGpa = iota // 四分制 GPA
	FiveGpa        // 五分制 GPA
)

// 四分制 GPA 计算
func calFour(grade float64) float64 {
	g := math.Min(90.0, grade)
	if g < 60 {
		return 0.0
	} else {
		return (g - 50) / 10.0
	}
}

// 五分制 GPA 计算
func calFive(g float64) float64 {
	if g < 60 {
		return 0.0
	} else {
		return (g - 50) / 10.0
	}
}

func (this MyScoreList) Gpa(gpaType int) float64 {
	var cal func(grade float64) float64
	if gpaType == FourGpa {
		cal = calFour
	} else {
		cal = calFive
	}
	totalCredits := 0.0
	totalGrades := 0.0
	for _, score := range this {
		credits := cast.ToFloat64(score.Credits)
		totalCredits += credits
		switch score.EffectiveScore {
		case "优秀", "优":
			totalGrades += 4.0 * credits
		case "良好", "良":
			totalGrades += 3.5 * credits
		case "中等", "中":
			totalGrades += 2.5 * credits
		case "及格":
			totalGrades += 1.5 * credits
		case "不及格":
			totalGrades += 0 * credits
		case "合格":
			totalGrades += 3.5 * credits
		case "不合格":
			totalGrades += 0 * credits
		case "未录入":
		case "待评教":
		case "":
			totalCredits -= credits
		default:
			grade := cast.ToFloat64(score.EffectiveScore)
			totalGrades += cal(grade) * credits
		}
	}
	if totalCredits == 0.0 {
		return cal(100.0)
	} else {
		// 保留 4 位有效数字
		return math.Round(totalGrades/totalCredits*10000) / 10000
	}

}
