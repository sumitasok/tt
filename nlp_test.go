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
var (
	layout = "Jan 2, 2006 (MST)"
)

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
	timeTable := Get(query)
	assert.Equal(3, len(timeTable.list))
	assert.Equal(time.Now().AddDate(0, 0, 1).Weekday().String(), timeTable.list[1].Weekday().String())
	printList(timeTable.list)
	assert.Equal(time.Now().AddDate(0, 0, 1).Format(layout), timeTable.list[0].Format(layout))
	assert.Equal(time.Now().AddDate(0, 0, 8).Format(layout), timeTable.list[1].Format(layout))
	assert.Equal(time.Now().AddDate(0, 0, 15).Format(layout), timeTable.list[2].Format(layout))

	query = "starting today, till 2 weeks from now, every " + weekDayList[(int(time.Now().Weekday())+2)]
	println(query)
	timeTable = Get(query)
	assert.Equal(2, len(timeTable.list))
	printList(timeTable.list)
	assert.Equal(time.Now().AddDate(0, 0, 2).Weekday().String(), timeTable.list[1].Weekday().String())
	assert.Equal(time.Now().AddDate(0, 0, 1+1).Format(layout), timeTable.list[0].Format(layout))
	assert.Equal(time.Now().AddDate(0, 0, 8+1).Format(layout), timeTable.list[1].Format(layout))

	query = "starting 2 weeks from now, till 4 weeks from now, every " + weekDayList[(int(time.Now().Weekday())+2)]
	println(query)
	timeTable = Get(query)
	assert.Equal(2, len(timeTable.list))
	printList(timeTable.list)
	assert.Equal(time.Now().AddDate(0, 0, 2).Weekday().String(), timeTable.list[1].Weekday().String())
	assert.Equal(time.Now().AddDate(0, 0, 1+14+1).Format(layout), timeTable.list[0].Format(layout))
	assert.Equal(time.Now().AddDate(0, 0, 8+14+1).Format(layout), timeTable.list[1].Format(layout))

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

	str = "30 days before every wednesday"
	assert.Equal("every", identifier(str))

	assert.True(true)
}

func TestGetByQuery(t *testing.T) {
	assert := assert.New(t)

	tt := Get("starting today")
	assert.Equal(time.Now().Day(), tt.start.time.Day())

	tt = Get("till tomorrow")
	assert.Equal(time.Now().AddDate(0, 0, 1).Day(), tt.end.time.Day())

	assert.True(true)
}

func TestStarting(t *testing.T) {
	assert := assert.New(t)

	tt := &TimeTable{}
	tt = starting("starting today", tt)
	assert.Equal(time.Now().Day(), tt.start.time.Day())

	assert.True(true)
}

func TestTill(t *testing.T) {
	assert := assert.New(t)

	tt := &TimeTable{}
	tt = till("till next week", tt)
	assert.Equal(time.Now().AddDate(0, 0, 7).Day(), tt.end.time.Day())

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

	id, offset = parseTime("till next week")
	assert.Equal("days_from_today", id)
	assert.Equal(7, offset)

	id, offset = parseTime("123 days from now")
	assert.Equal("days_from_today", id)
	assert.Equal(123, offset)

	id, offset = parseTime("3 weeks from now")
	assert.Equal("days_from_today", id)
	assert.Equal(21, offset)

	id, offset = parseTime("30 days before")
	assert.Equal("days_from_today", id)
	assert.Equal(-30, offset)

	assert.True(true)
}

func TestStringToInt(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(123, stringToInt("123"))

	assert.True(true)
}

func TestMakeTime(t *testing.T) {
	assert := assert.New(t)

	ti := makeTime("starting today")
	assert.Equal(ti.Year(), time.Now().Year())
	assert.Equal(ti.Month(), time.Now().Month())
	assert.Equal(ti.Day(), time.Now().Day())

	ti = makeTime("till tomorrow")
	assert.Equal(ti.Year(), time.Now().AddDate(0, 0, 1).Year())
	assert.Equal(ti.Month(), time.Now().AddDate(0, 0, 1).Month())
	assert.Equal(ti.Day(), time.Now().AddDate(0, 0, 1).Day())

	ti = makeTime("starting yesterday")
	assert.Equal(ti.Year(), time.Now().AddDate(0, 0, -1).Year())
	assert.Equal(ti.Month(), time.Now().AddDate(0, 0, -1).Month())
	assert.Equal(ti.Day(), time.Now().AddDate(0, 0, -1).Day())

	ti = makeTime("starting next week")
	assert.Equal(ti.Year(), time.Now().AddDate(0, 0, 7).Year())
	assert.Equal(ti.Month(), time.Now().AddDate(0, 0, 7).Month())
	assert.Equal(ti.Day(), time.Now().AddDate(0, 0, 7).Day())

	assert.True(true)
}

func printList(timeList []time.Time) {
	for i := range timeList {
		println(timeList[i].Format(layout))
	}
}
