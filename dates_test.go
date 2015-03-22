package timetable

import (
	assert "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestToday(t *testing.T) {
	assert := assert.New(t)
	// Set custom start date
	tt := ListOf(7).StartingFrom(Today())
	assert.Equal(time.Now().Day(), tt.start.time.Day())
}

func TestStartToday(t *testing.T) {
	assert := assert.New(t)

	tt := TimeTable{}

	d := tt.Starting().Today()
	assert.Equal(time.Now().Day(), d.start.time.Day())

	assert.True(true)
}

func TestStartTomorrow(t *testing.T) {
	assert := assert.New(t)

	tt := TimeTable{}

	d := tt.Starting().Tomorrow()
	assert.Equal(time.Now().AddDate(0, 0, 1).Day(), d.start.time.Day())

	assert.True(true)
}

// ListOf(22).Starting().Today().Ending().Next().Month().Every().Monday()

func ExampleDate_Today() {
	ListOf(22).Starting().Today()
}

func ExampleDate_Tomorrow() {
	ListOf(22).Starting().Tomorrow()
}
