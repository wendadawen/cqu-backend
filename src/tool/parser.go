package tool

import (
	"fmt"
	"log"
	"regexp"
)

// TODO
func FindFirstSubMatch(r *regexp.Regexp, str string) string {
	match := r.FindStringSubmatch(str)
	if len(match) <= 1 {
		msg := fmt.Sprintf("未能匹配到：%s\n源字符串：%s", r, str)
		log.Printf(msg)
		//go BotMessage(msg)
		return ""
	} else {
		return match[1]
	}
}
