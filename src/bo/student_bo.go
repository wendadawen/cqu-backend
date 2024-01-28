package bo

type StudentType int

const (
	UndergraduateStudent StudentType = iota // 本科
	GraduateStudent                         // 研究生
)

// StudentInfoBo 学生教务学籍信息
type StudentInfoBo struct {
	Type                 StudentType // 学生类型
	studentId            string      // 学号
	studentName          string      // 姓名
	gender               string      // 性别
	Grade                string      // 年级
	deptName             string      // 学院
	majorName            string      // 专业
	className            string      // 班级
	IdNumber             string      // 身份证号
	PoliticalStatus      string      // 政治面貌
	Nationality          string      // 民族
	Phone                string      // 电话
	Email                string      //邮箱
	AuthId               string      //统一认证号
	Birthday             string      //生日
	HomeAddress          string      //家庭住址
	Gpa                  string      //绩点
	MajorRanking         string      //专业排名
	Duration             string      //学制
	ObtainSchoolRollTime string      //取得学籍时间
	EnrollmentTime       string      // 入学时间
	DepartureTime        string      // 离校时间
	StuSourceRegion      string      // 生源地
	StuSourceUnit        string      //生源单位
}
