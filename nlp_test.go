package timetable

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

/*
starting today, till 7 days from now, every wednesday
starting 7 days from now, till next month, every wed
starting today, till next week, every alternate day
*/

func TestNlp(t *testing.T) {
	assert := assert.New(t)

	tt := ListOf(7).Starting().Today().EndingOn(time.Now().AddDate(0, 0, 16)).Select(WEEK, int(time.Now().Weekday())+1)
	assert.Equal(3, len(tt.list))
	assert.Equal(time.Now().AddDate(0, 0, 1).Weekday().String(), tt.list[1].Weekday().String())

	// Now replicate the same result using NLP

	weekDayList := map[int]string{
		0: "sunday",
		1: "monday",
		2: "tuesday",
		3: "wednesday",
		4: "thursday",
		5: "friday",
		6: "saturday",
	}

	query := "starting today, till 16 days from now, every " + weekDayList[(int(time.Now().Weekday())+1)]
	println(query)
	// timeTable := Get(query)
	// assert.Equal(3, len(timeTable.list))
	// assert.Equal(time.Now().AddDate(0, 0, 1).Weekday().String(), timeTable.list[1].Weekday().String())

	assert.True(true)
}

func TestIdentifier(t *testing.T) {
	assert := assert.New(t)

	str := "starting today"
	assert.Equal("starting", identifier(str))

	str = "till today"
	assert.Equal("till", identifier(str))

	str = "ending today"
	assert.Equal("till", identifier(str))

	str = "every day"
	assert.Equal("every", identifier(str))

	str = "    starting today"
	assert.Equal("starting", identifier(str))

	str = "    till today"
	assert.Equal("till", identifier(str))

	str = "    ending today"
	assert.Equal("till", identifier(str))

	str = "    every day"
	assert.Equal("every", identifier(str))

	assert.True(true)
}

func TestParseTime(t *testing.T) {
	assert := assert.New(t)

	id, offset := parseTime("starting today")
	assert.Equal("days_from_today", id)
	assert.Equal(0, offset)

	id, offset = parseTime("starting tomorrow")
	assert.Equal("days_from_today", id)
	assert.Equal(1, offset)

	id, offset = parseTime("till yesterday")
	assert.Equal("days_from_today", id)
	assert.Equal(-1, offset)

	assert.True(true)
}
