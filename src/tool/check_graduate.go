package tool

func CheckGraduate(StuId, CasPwd string) bool {
	if len(StuId) >= 11 {
		return true
	}
	return false
}
