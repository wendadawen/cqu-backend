package model

type Rank struct {
	StuId        string `xorm:"not null pk VARCHAR(255)" json:"stuId,omitempty"`
	Name         string `xorm:"not null VARCHAR(255)" json:"name,omitempty"`
	College      string `xorm:"not null index VARCHAR(255)" json:"college,omitempty"`
	Major        string `xorm:"not null index VARCHAR(255)" json:"major,omitempty"`
	Class        string `xorm:"not null index VARCHAR(255)" json:"class,omitempty"`
	Grade        string `xorm:"not null VARCHAR(20)" json:"grade omitempty"`
	Gpa          string `xorm:"not null VARCHAR(20)" json:"gpa"`        // 绩点
	WeightAvg    string `xorm:"not null VARCHAR(20)" json:"weight_avg"` // 加权平均分
	GradeRanking string `xorm:"VARCHAR(20)" json:"grade_ranking"`       // 年级排名
	MajorRanking string `xorm:"VARCHAR(20)" json:"major_ranking"`       // 专业排名
	ClassRanking string `xorm:"VARCHAR(20)" json:"class_ranking"`       // 班级排名
}
