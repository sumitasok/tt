package timetable

import (
	"time"
)

type TimeTable struct {
	count int         // maximum number of datetime to be returned
	start *Date       // time table to start looking from
	end   *Date       // time table to stop looking at
	list  []time.Time // the final output list
}

func ListOf(count int) *TimeTable {
	tt := TimeTable{
		count: count,
	}
	tt.start = &Date{
		time:      time.Now(),
		timetable: &tt,
		kind:      START,
	}
	return &tt
}

func (tt *TimeTable) StartingFrom(t time.Time) *TimeTable {
	tt.start = &Date{
		time:      t,
		timetable: tt,
		kind:      START,
	}

	return tt
}

func (tt *TimeTable) Starting() *Date {
	tt.start = &Date{
		timetable: tt,
		kind:      START,
	}

	return tt.start
}
