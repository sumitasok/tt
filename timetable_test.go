package timetable

import (
	assert "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestListOfInit(t *testing.T) {
	assert := assert.New(t)
	tt := ListOf(7)

	// Set Number of Maximum output
	assert.Equal(7, tt.count)

	// Set Start date as today to start with
	// assert.Equal(time.Now().Day(), tt.start.Day())

	assert.True(true)
}

func TestStartingFrom(t *testing.T) {
	assert := assert.New(t)

	// Set custom start date
	tt := ListOf(7).StartingFrom(time.Now().AddDate(0, 0, 2))
	assert.Equal(time.Now().AddDate(0, 0, 2).Day(), tt.start.time.Day())
}

// starting today till next week get every wednesday
// starting today till 14 days get every friday
func TestSelect(t *testing.T) {
	assert := assert.New(t)

	for _, i := range []int{1, 2, 4} {
		dateOffset := 7 * i
		tt := ListOf(7).Starting().Today().EndingOn(time.Now().AddDate(0, 0, dateOffset)).Select(WEEK, THURSDAY)
		assert.Equal(i, len(tt.list))
		assert.Equal("Thursday", tt.list[i-1].Weekday().String())

		tt = ListOf(7).StartingFrom(time.Now().AddDate(0, 0, 2)).EndingOn(time.Now().AddDate(0, 0, 16)).Select(WEEK, WEDNESDAY)
		assert.Equal(2, len(tt.list))
		assert.Equal("Wednesday", tt.list[1].Weekday().String())
	}

	tt := ListOf(7).Starting().Today().EndingOn(time.Now().AddDate(0, 0, 16)).Select(WEEK, int(time.Now().Weekday())+1)
	assert.Equal(3, len(tt.list))
	assert.Equal(time.Now().AddDate(0, 0, 1).Weekday().String(), tt.list[1].Weekday().String())

	assert.True(true)
}

// starting today till Jan 20, 1985 get every week friday minus 30 days
// Starting().Today().EndingOn(7 days from now).every("week friday").minus("30 days")

func TestMinusDays(t *testing.T) {
	assert := assert.New(t)

	minus30Days := ListOf(7).Starting().Today().EndingOn(time.Now().AddDate(0, 0, 28)).Select(WEEK, FRIDAY).Minus(60).Days()
	fridays := ListOf(7).Starting().Today().EndingOn(time.Now().AddDate(0, 0, 28)).Select(WEEK, FRIDAY).list

	layout := "Mon Jan 2 15:04:05 -0700 MST 2006"
	for i, f := range fridays {
		assert.Equal(f.AddDate(0, 0, -60).Format(layout), minus30Days.list[i].Format(layout))
		println(f.AddDate(0, 0, -60).Format(layout), minus30Days.list[i].Format(layout))
	}

	assert.True(true)
}
