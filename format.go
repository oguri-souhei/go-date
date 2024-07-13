package date

import (
	"fmt"
	"time"
)

// Implements fmt.Stringer
func (d Date) String() string {
	return d.Format(time.DateOnly)
}

func (d Date) Format(layout string) string {
	return d.t.Format(layout)
}

var dateFormats = []string{
	time.DateOnly,
	"02-01-2006",
	"01/02/2006",
	"2006/01/02",
	"Jan 2, 2006",
	"2 Jan 2006",
	"02 Jan 2006",
	time.RFC3339,
	time.RFC1123,
}

func Parse(v string) (Date, error) {
	for i := 0; i < len(dateFormats); i++ {
		d, err := time.Parse(dateFormats[i], v)
		if err == nil {
			return New(d.Year(), d.Month(), d.Day()), nil
		}
	}
	return Date{}, fmt.Errorf("failed to parse string %+v to date", v)
}
