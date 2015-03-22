package computation

import (
	"time"
)

const (
	SUNDAY    = 0
	MONDAY    = 1
	TUESDAY   = 2
	WEDNESDAY = 3
	THURSDAY  = 4
	FRIDAY    = 5
	SATURDAY  = 6
)

func Weekly(item int, starting time.Time, ending time.Time) []time.Time {
	var times []time.Time

	timeCheck := starting
	firstLoop := true

	for timeCheck.Before(ending) {
		for firstLoop {
			if int(timeCheck.Weekday()) == item {
				times = append(times, timeCheck)
				firstLoop = false
				timeCheck = timeCheck.AddDate(0, 0, 7)
				break
			}
			timeCheck = timeCheck.AddDate(0, 0, 1)
		}

		if timeCheck.Before(ending) {
			times = append(times, timeCheck)
			timeCheck = timeCheck.AddDate(0, 0, 7)
		}
	}

	return times
}
