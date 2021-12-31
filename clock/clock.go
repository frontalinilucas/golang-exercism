package clock

import "time"

type Clock struct {
	hour, minute int
}

func getTime(h, m int) time.Time {
	t := time.Time{}
	return time.Date(t.Year(), t.Month(), t.Day(), h, m, t.Second(), t.Nanosecond(), t.Location())
}

func New(h, m int) Clock {
	t := getTime(h, m)
	return Clock{
		hour:   t.Hour(),
		minute: t.Minute(),
	}
}

func (c Clock) Add(m int) Clock {
	t := getTime(c.hour, c.minute).Add(time.Duration(m) * time.Minute)
	return Clock{
		hour:   t.Hour(),
		minute: t.Minute(),
	}
}

func (c Clock) Subtract(m int) Clock {
	t := getTime(c.hour, c.minute).Add(time.Duration(m) * time.Minute * -1)
	return Clock{
		hour:   t.Hour(),
		minute: t.Minute(),
	}
}

func (c Clock) String() string {
	return getTime(c.hour, c.minute).Format("15:04")
}
