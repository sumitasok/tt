A Package that makes playing with date easy.

#### How to use
```
import(
	"github.com/sumitasok/timetable"
)
```

```
query := "starting today, till 2 weeks from now, every wednesday"
list := timetable.Get(query)
```

other queries you might like

```
starting tomorrow, till 40 days from now, every thursday
starting yesterday, till 4 weeks from now, 14 days before every wednesday
starting 2 weeks from now, till 4 weeks from now, 21 days after every monday
starting 21 days before now, till 4 weeks from now, every tuesday
starting 21 days before today, till 4 weeks from now, every tuesday

```


#### Usage Old School

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
Starting().Today().EndingOn(7 days from now).every("week friday").minus("30 days")
