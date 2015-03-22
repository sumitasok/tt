package timetable

import (
	"time"
)

type TimeTable struct {
	count int         // maximum number of datetime to be returned
	start time.Time   // time table to start looking from
	end   time.Time   // time table to stop looking at
	list  []time.Time // the final output list
}

func ListOf(count int) *TimeTable {
	tt := TimeTable{
		count: count,
		start: time.Now(),
	}
	return &tt
}

func (tt *TimeTable) StartingFrom(t time.Time) *TimeTable {
	tt.start = t
	return tt
}
