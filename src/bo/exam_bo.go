package bo

type ExamScheduleBo = []ExamBo

type ExamBo struct {
	ExamId       string `json:"-"` // Id
	ExamTitle    string ``         // 考试名称
	ExamTime     string ``         // 考试时间
	ExamSeat     string ``         // 座位号
	ExamLocation string ``         // 考场
	ExamCredits  string ``         // 学分
	ExamCategory string ``         // 考试类型
	ExamStyle    string ``         // 考试方式
}
