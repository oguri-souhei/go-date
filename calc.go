package date

func (d1 Date) Before(d2 Date) bool {
	return d1.t.Before(d2.t)
}

func (d1 Date) After(d2 Date) bool {
	return d1.t.After(d2.t)
}

func (d1 Date) Equal(d2 Date) bool {
	y := d1.Year() == d2.Year()
	m := d1.Month() == d2.Month()
	d := d1.Day() == d2.Day()
	return y && m && d
}

func (d Date) Add(y, m, day int) Date {
	t := d.t.AddDate(y, m, day)
	return Date{t: t}
}

func (d Date) AddYear(y int) Date {
	return d.Add(y, 0, 0)
}

func (d Date) AddMonth(m int) Date {
	return d.Add(0, m, 0)
}

func (d Date) AddDay(day int) Date {
	return d.Add(0, 0, day)
}
