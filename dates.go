package timetable

import (
	c "github.com/sumitasok/timetable/computation"
	"time"
)

const (
	START = "START"
	END   = "END"
	EVERY = "EVERY"

	MINUS = "MINUS"
	NEXT  = "NEXT"
	PLUS  = "PLUS"
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

func (d *Date) Next() *Date {
	d.kind = NEXT
	return d
}

func (d *Date) Week() *TimeTable {
	t := d.timetable
	start := t.start.time
	end := t.end.time
	timeChk := start

	switch d.kind {
	case NEXT:
		if len(t.list) == 0 {
			list := c.Weekly(int(start.Weekday()), start, end)
			d.timetable.list = list
			return t
		} else {
			var list []time.Time
			for i := range t.list {
				timeChk = t.list[i].AddDate(0, 0, 7)
				list = append(list, timeChk)
				t.list = list
			}
			return t
		}
	}
	return t
}

func (d *Date) Days() *TimeTable {
	switch d.kind {
	case MINUS:
		var list []time.Time
		for i := range d.timetable.list {
			list = append(list, d.timetable.list[i].AddDate(0, 0, -(d.n)))
		}
		d.timetable.list = list
	case PLUS:
		var list []time.Time
		for i := range d.timetable.list {
			list = append(list, d.timetable.list[i].AddDate(0, 0, (d.n)))
		}
		d.timetable.list = list
	}
	return d.timetable
}

func (d *Date) Months() *TimeTable {
	switch d.kind {
	case MINUS:
		var list []time.Time
		for i := range d.timetable.list {
			list = append(list, d.timetable.list[i].AddDate(0, -(d.n), 0))
		}
		d.timetable.list = list
	case PLUS:
		var list []time.Time
		for i := range d.timetable.list {
			list = append(list, d.timetable.list[i].AddDate(0, (d.n), 0))
		}
		d.timetable.list = list
	}
	return d.timetable
}
