package tool

import (
	"cqu-backend/src/config/setting"
	"github.com/spf13/cast"
	"time"
)

func CurrentWeek() int {
	const layout = "2006-01-02"
	termStartDateString := setting.CquConfig.GetString("term_start_date")
	termStartDate, _ := time.Parse(layout, termStartDateString)
	today, _ := time.Parse(layout, time.Now().Format(layout))
	return cast.ToInt(today.Sub(termStartDate).Hours()/(24*7)) + 1
}
