package timetable

import (
	"time"
)

const (
	START = "START"
	END   = "END"
	EVERY = "EVERY"
)

type Date struct {
	time      time.Time  // the actual time computed
	timetable *TimeTable // the actual time table structs reference
	kind      string     // start | end | every - reference for time computation
}

func Today() time.Time {
	return time.Now()
}

func (d *Date) Today() *TimeTable {
	d.time = time.Now()
	return d.timetable
}

func (d *Date) Tomorrow() *TimeTable {
	d.time = time.Now().AddDate(0, 0, 1)
	return d.timetable
}
