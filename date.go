package date

import "time"

type Date struct {
	t time.Time
}

// New creates a new Date instance for the given year, month, and day.
// The time is set to midnight UTC.
func New(y int, m time.Month, d int) Date {
	t := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	return Date{t: t}
}

func (d Date) Year() int {
	return d.t.Year()
}

func (d Date) Month() time.Month {
	return d.t.Month()
}

func (d Date) Day() int {
	return d.t.Day()
}
