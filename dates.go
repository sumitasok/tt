package timetable

import (
	"time"
)

const (
	START = "START"
	END   = "END"
	EVERY = "EVERY"

	MINUS = "MINUS"
)

type Date struct {
	time      time.Time  // the actual time computed
	timetable *TimeTable // the actual time table structs reference
	kind      string     // start | end | every - reference for time computation
	n         int
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

func (d *Date) Days() *TimeTable {
	switch d.kind {
	case MINUS:
		var list []time.Time
		for i := range d.timetable.list {
			list = append(list, d.timetable.list[i].AddDate(0, 0, -(d.n)))
		}
		d.timetable.list = list
	}
	return d.timetable
}
