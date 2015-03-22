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
