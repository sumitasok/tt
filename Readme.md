A Package that makes playing with date easy.

#### Usage

Starting().Today().EndingOn(7 days from now).every("week friday").minus("30 days")

If you want to find all Wednesdays between today and 14 days from then, use

```
import(
	t "github.com/sumitasok/timetable"
)
```

```
t.ListOf(7).Starting().Today().EndingOn(time.Now().AddDate(0,0,14)).
	Select(t.WEEK, t.WEDNESDAY).Get()
```
returns a list of time `[]time.Time`

To find 30 days prior to every wednesday in a time span.

```
t.ListOf(7).Starting().Today().EndingOn(time.Now().AddDate(0,0,14)).
	Select(t.WEEK, t.WEDNESDAY).Minus(30).Days().Get()
```

TO-DO: make sure start, till and every is defined before query executes
starting 14 days from now, till 3 weeks from `then|that`, every wednesday
starting 14 days from now, till 3 weeks from `then|that`, 30 days before every wednesday