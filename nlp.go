package timetable

import (
	"regexp"
	"strings"
	"time"
)

func Get(str string) *TimeTable {
	tt := &TimeTable{}
	query := splitByComma(str)
	for _, q := range query {
		switch identifier(q) {
		case "starting":
			starting(q, tt)
		}
	}
	return tt
}

func starting(query string, tt *TimeTable) *TimeTable {
	tt.StartingFrom(makeTime(query))
	return tt
}

func makeTime(query string) time.Time {
	// re := regexp.MustCompile("(today)")
	// s := re.FindString(query)
	return time.Now()
}

func parseTime(query string) (string, int) {
	re := regexp.MustCompile("(today)")
	if re.Match([]byte(query)) {
		return "days_from_today", 0
	}

	re = regexp.MustCompile("(tomorrow)")
	if re.Match([]byte(query)) {
		return "days_from_today", 1
	}

	return "", 0
}

func splitByComma(str string) []string {
	return strings.Split(str, ",")
}

func identifier(str string) string {
	re := regexp.MustCompile("[a-zA-Z]+")
	s := re.FindString(str)
	switch s {
	case "starting":
		return "starting"
	case "till", "ending":
		return "till"
	case "every":
		return "every"
	default:
		return ""
	}
}
