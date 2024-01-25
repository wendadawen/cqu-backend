package cas

const (
	// 统一认证的网址
	authUrl = "https://sso.cqu.edu.cn"
	// 统一认证登录
	loginUrl = "/login"
	// 获取是否需要输入验证码
	validCodeUrl = "/authserver/needCaptcha.html?username="
	// 研究生 MIS 系统重定向
	misReferUrl = "?service=http%3A%2F%2Fmis.cqu.edu.cn%2Fmis%2Fsso_entry.jsp%3FuserType%3Dstudent"
)
