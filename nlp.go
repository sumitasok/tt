package timetable

import (
	"regexp"
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
			date := tt.Select(WEEK, item)

			re = regexp.MustCompile("(days before)")
			if re.Match([]byte(query)) {
				id, number := parseTime(query)
				if id == "days_before_today" {
					date.Minus(number).Days()
				}
			}

			re = regexp.MustCompile("(days after)")
			if re.Match([]byte(query)) {
				id, number := parseTime(query)
				if id == "days_from_today" {
					date.Plus(number).Days()
				}
			}

			re = regexp.MustCompile("(months before)")
			if re.Match([]byte(query)) {
				id, number := parseTime(query)
				if id == "months_before_today" {
					date.Minus(number).Months()
				}
			}

			re = regexp.MustCompile("(months after)")
			if re.Match([]byte(query)) {
				id, number := parseTime(query)
				if id == "months_from_today" {
					date.Plus(number).Months()
				}
			}

		}
	}

	return tt
}

func makeTime(query string) time.Time {
	id, offset := parseTime(query)
	switch id {
	case "days_from_today":
		return time.Now().AddDate(0, 0, offset)
	case "days_before_today":
		return time.Now().AddDate(0, 0, -offset)
	case "months_from_today":
		return time.Now().AddDate(0, offset, 0)
	case "months_before_today":
		return time.Now().AddDate(0, -offset, 0)
	}
	return time.Time{}
}

func parseTime(query string) (string, int) {
	re := regexp.MustCompile("(days from now)")
	if re.Match([]byte(query)) {
		re1 := regexp.MustCompile("[0-9]+")
		if re1.Match([]byte(query)) {
			number := re1.FindString(query)
			return "days_from_today", stringToInt(number)
		}
		return "", 0
	}

	re = regexp.MustCompile("(months from now)")
	if re.Match([]byte(query)) {
		re1 := regexp.MustCompile("[0-9]+")
		if re1.Match([]byte(query)) {
			number := re1.FindString(query)
			return "months_from_today", stringToInt(number)
		}
		return "", 0
	}

	re = regexp.MustCompile("(days before today)")
	if re.Match([]byte(query)) {
		re1 := regexp.MustCompile("[0-9]+")
		if re1.Match([]byte(query)) {
			number := re1.FindString(query)
			return "days_before_today", stringToInt(number)
		}
		return "", 0
	}

	re = regexp.MustCompile("(months before today)")
	if re.Match([]byte(query)) {
		re1 := regexp.MustCompile("[0-9]+")
		if re1.Match([]byte(query)) {
			number := re1.FindString(query)
			return "months_before_today", stringToInt(number)
		}
		return "", 0
	}

	re = regexp.MustCompile("(days before now)")
	if re.Match([]byte(query)) {
		re1 := regexp.MustCompile("[0-9]+")
		if re1.Match([]byte(query)) {
			number := re1.FindString(query)
			return "days_before_today", stringToInt(number)
		}
		return "", 0
	}

	re = regexp.MustCompile("(weeks from now)")
	if re.Match([]byte(query)) {
		re1 := regexp.MustCompile("[0-9]+")
		if re1.Match([]byte(query)) {
			number := re1.FindString(query)
			return "days_from_today", stringToInt(number) * 7
		}
		return "", 0
	}

	// Keep this matching here to avoid this becoming a wild card and escaping earlier matches
	days := []string{"today", "tomorrow", "yesterday", "next week"}
	for _, item := range days {
		re := regexp.MustCompile("(" + item + ")")
		if re.Match([]byte(query)) {
			return "days_from_today", daysFromTodayConstants(item)
		}
	}

	// has to be towards the last, as this is for `every` identifier
	re = regexp.MustCompile("(days before)")
	if re.Match([]byte(query)) {
		re1 := regexp.MustCompile("[0-9]+")
		if re1.Match([]byte(query)) {
			number := re1.FindString(query)
			return "days_before_today", stringToInt(number)
		}
		return "", 0
	}

	re = regexp.MustCompile("(days after)")
	if re.Match([]byte(query)) {
		re1 := regexp.MustCompile("[0-9]+")
		if re1.Match([]byte(query)) {
			number := re1.FindString(query)
			return "days_from_today", stringToInt(number)
		}
		return "", 0
	}

	re = regexp.MustCompile("(months before)")
	if re.Match([]byte(query)) {
		re1 := regexp.MustCompile("[0-9]+")
		if re1.Match([]byte(query)) {
			number := re1.FindString(query)
			return "months_before_today", stringToInt(number)
		}
		return "", 0
	}

	re = regexp.MustCompile("(months after)")
	if re.Match([]byte(query)) {
		re1 := regexp.MustCompile("[0-9]+")
		if re1.Match([]byte(query)) {
			number := re1.FindString(query)
			return "months_from_today", stringToInt(number)
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
		re := regexp.MustCompile("(every)")
		if re.Match([]byte(str)) {
			return "every"
		}
		return ""
	}
}
