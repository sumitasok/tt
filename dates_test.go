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

func TestMinusMonth(t *testing.T) {
	assert := assert.New(t)

	tt := TimeTable{}
	tt.list = []time.Time{time.Now()}
	d := Date{timetable: &tt}
	d.Today().Minus(7).Months()

	assert.Equal(time.Now().AddDate(0, -7, 0).Month().String(), tt.list[0].Month().String())

	assert.True(true)
}

func TestPlusMonth(t *testing.T) {
	assert := assert.New(t)

	tt := TimeTable{}
	tt.list = []time.Time{time.Now()}
	d := Date{timetable: &tt}
	d.Today().Plus(7).Months()

	assert.Equal(time.Now().AddDate(0, 7, 0).Month().String(), tt.list[0].Month().String())

	assert.True(true)
}

func TestMinusDay(t *testing.T) {
	assert := assert.New(t)

	tt := TimeTable{}
	tt.list = []time.Time{time.Now()}
	d := Date{timetable: &tt}
	d.Today().Minus(7).Days()

	assert.Equal(time.Now().AddDate(0, 0, -7).Day(), tt.list[0].Day())

	assert.True(true)
}

func TestPlusDay(t *testing.T) {
	assert := assert.New(t)

	tt := TimeTable{}
	tt.list = []time.Time{time.Now()}
	d := Date{timetable: &tt}
	d.Today().Plus(7).Days()

	assert.Equal(time.Now().AddDate(0, 0, 7).Day(), tt.list[0].Day())

	assert.True(true)
}

// ListOf(22).Starting().Today().Ending().Next().Month().Every().Monday()

func ExampleDate_Today() {
	ListOf(22).Starting().Today()
}

func ExampleDate_Tomorrow() {
	ListOf(22).Starting().Tomorrow()
}
