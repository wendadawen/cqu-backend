package my

const (
	myHost                = "https://my.cqu.edu.cn"
	myExam                = "/api/exam/examTask/get-student-exam-list?"
	myAuthserverAuthorize = "/authserver/oauth/authorize?client_id=personal-prod&response_type=code&scope=all&state=&redirect_uri=https://my.cqu.edu.cn/workspace/token-index"
	myAuthserverCas       = "/authserver/authentication/cas"
	myRank                = "/api/sam/score/student/studentGpaRanking"
	myScore               = "/api/sam/score/student/score"
	myClass               = "/api/timetable/class/stu-course?" // https://my.cqu.edu.cn/workspace/course这个网页的课表
)
