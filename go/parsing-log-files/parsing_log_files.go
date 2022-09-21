package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func escape_squares(s string) string {
	for _, square := range []string{"[", "]"} {
		s = strings.Replace(s,
			square,
			`\`+square,
			-1)
	}

	return s
}

func IsValidLine(text string) bool {

	legit_start := []string{
		"[TRC]",
		"[DBG]",
		"[INF]",
		"[WRN]",
		"[ERR]",
		"[FTL]",
	}

	for _, start := range legit_start {
		re, _ := regexp.Compile(
			fmt.Sprintf(
				"^%s",
				escape_squares(start)))
		if re.MatchString(text) {
			return true
		}
	}

	return false
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`\<[~\*=-]*\>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	total := 0
	re, _ := regexp.Compile(`(?i)".*password.*"`)
	for _, line := range lines {
		if re.MatchString(line) {
			total++
		}
	}
	return total
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	if re.MatchString(text) {
		return re.ReplaceAllString(text, "")
	} else {
		return text
	}
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User +([a-zA-Z0-9]+)`)
	for idx, line := range lines {
		sl := re.FindStringSubmatch(line)
		if len(sl) == 0 {
			continue
		}
		user := sl[1]
		lines[idx] = fmt.Sprintf("[USR] %s %s", user, line)
	}
	return lines
}
