package timetable

import (
	c "github.com/sumitasok/timetable/computation"
	"time"
)

const (
	WEEK = "WEEK"
)

const (
	SUNDAY    = c.SUNDAY
	MONDAY    = c.MONDAY
	TUESDAY   = c.TUESDAY
	WEDNESDAY = c.WEDNESDAY
	THURSDAY  = c.THURSDAY
	FRIDAY    = c.FRIDAY
	SATURDAY  = c.SATURDAY
)

const (
	ErrWhileComputation = "this span cannot be computed"
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

func (tt *TimeTable) EndingOn(t time.Time) *TimeTable {
	tt.end = &Date{
		time:      t,
		timetable: tt,
		kind:      END,
	}

	return tt
}

func (tt *TimeTable) Select(collection string, item int) *TimeTable {
	switch collection {
	case WEEK:
		tt.list = c.Weekly(item, tt.start.time, tt.end.time)
		break
	default:
		panic(ErrWhileComputation)
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

func (tt *TimeTable) Ending() *Date {
	tt.end = &Date{
		timetable: tt,
		kind:      END,
	}

	return tt.end
}
