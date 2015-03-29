package timetable

import (
	"regexp"
	"strconv"
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
		case "till":
			till(q, tt)
		case "every":
			every(q, tt)
		}
	}
	return tt
}

func starting(query string, tt *TimeTable) *TimeTable {
	tt.StartingFrom(makeTime(query))
	return tt
}

func till(query string, tt *TimeTable) *TimeTable {
	tt.EndingOn(makeTime(query))
	return tt
}

var (
	weekPointer = map[string]int{
		"sunday":    0,
		"monday":    1,
		"tuesday":   2,
		"wednesday": 3,
		"thursday":  4,
		"friday":    5,
		"saturday":  6,
	}
)

func every(query string, tt *TimeTable) *TimeTable {
	for weekName, item := range weekPointer {
		re := regexp.MustCompile("(" + weekName + ")")
		if re.Match([]byte(query)) {
			tt.Select(WEEK, item)
		}
	}
	return tt
}

func makeTime(query string) time.Time {
	id, offset := parseTime(query)
	switch id {
	case "days_from_today":
		return time.Now().AddDate(0, 0, offset)
	}
	return time.Time{}
}

func parseTime(query string) (string, int) {
	days := []string{"today", "tomorrow", "yesterday", "next week"}
	for _, item := range days {
		re := regexp.MustCompile("(" + item + ")")
		if re.Match([]byte(query)) {
			return "days_from_today", daysFromTodayConstants(item)
		}
	}

	re := regexp.MustCompile("(days from now)")
	if re.Match([]byte(query)) {
		re1 := regexp.MustCompile("[0-9]+")
		if re1.Match([]byte(query)) {
			number := re1.FindString(query)
			return "days_from_today", stringToInt(number)
		}
		return "", 0
	}

	return "", 0
}

func daysFromTodayConstants(str string) int {
	daysFromToday := map[string]int{
		"today":     0,
		"tomorrow":  1,
		"yesterday": -1,
		"next week": 7,
	}
	return daysFromToday[str]
}

func stringToInt(str string) int {
	if i, err := strconv.ParseInt(str, 0, 64); err == nil {
		return int(i)
	} else {
		return 0
	}
}

func splitByComma(str string) []string {
	return strings.Split(str, ",")
}

func identifier(str string) string {
	re := regexp.MustCompile("[a-zA-Z]+")
	s := re.FindString(str)
	switch s {
	case "starting", "from":
		return "starting"
	case "till", "ending":
		return "till"
	case "every":
		return "every"
	default:
		return ""
	}
}
