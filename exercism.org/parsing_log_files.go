package main

import (
	"fmt"
	"log"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(ERR|TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[-=~*]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)
	ans := 0
	for _, line := range lines {
		if re.MatchString(line) {
			ans++
		}
	}
	return ans
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)
	ret := []string{}
	for _, line := range lines {
		founds := re.FindStringSubmatch(line)
		if founds != nil {
			line = fmt.Sprintf("[USR] %s %s", founds[1], line)
		}
		ret = append(ret, line)
	}
	return ret
}

func main() {
	log.Println(IsValidLine("[ERR] A good error here"))
	log.Println(IsValidLine("Any old [ERR] text"))
	log.Println(IsValidLine("[BOB] Any old text"))
	log.Println(SplitLogLine("section 1<*>section 2<~~~>section 3"))
}
