package timetable

import (
	"strconv"
	"strings"
)

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
