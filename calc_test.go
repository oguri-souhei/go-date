package date

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func ExampleDate_Before() {
	d1 := New(2023, time.July, 13)
	d2 := New(2023, time.August, 13)

	fmt.Println(d1.Before(d2))
	// Output: true
}

func ExampleDate_After() {
	d1 := New(2023, time.July, 13)
	d2 := New(2023, time.June, 13)

	fmt.Println(d1.After(d2))
	// Output: true
}

func ExampleDate_Equal() {
	d1 := New(2023, time.July, 13)
	d2 := New(2023, time.July, 13)
	d3 := New(2023, time.July, 14)

	fmt.Println(d1.Equal(d2))
	fmt.Println(d1.Equal(d3))
	// Output:
	// true
	// false
}

func ExampleDate_Add() {
	d := New(2023, time.July, 13)
	d2 := d.Add(1, 1, 1)

	fmt.Println(d2)
	// Output: 2024-08-14
}

func ExampleDate_AddYear() {
	d := New(2023, time.July, 13)
	d2 := d.AddYear(1)

	fmt.Println(d2)
	// Output: 2024-07-13
}

func ExampleDate_AddMonth() {
	d := New(2023, time.July, 13)
	d2 := d.AddMonth(1)

	fmt.Println(d2)
	// Output: 2023-08-13
}

func ExampleDate_AddDay() {
	d := New(2023, time.July, 13)
	d2 := d.AddDay(1)

	fmt.Println(d2)
	// Output: 2023-07-14
}

func TestDate_Before(t *testing.T) {
	tests := []struct {
		name  string
		date1 Date
		date2 Date
		want  bool
	}{
		{
			name:  "date1 is before date2",
			date1: New(2024, 7, 12),
			date2: New(2024, 7, 13),
			want:  true,
		},
		{
			name:  "date1 is after date2",
			date1: New(2024, 7, 12),
			date2: New(2024, 7, 11),
			want:  false,
		},
		{
			name:  "date1 is same date2",
			date1: New(2024, 7, 12),
			date2: New(2024, 7, 12),
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.date1.Before(tt.date2)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDate_After(t *testing.T) {
	tests := []struct {
		name  string
		date1 Date
		date2 Date
		want  bool
	}{
		{
			name:  "date1 is before date2",
			date1: New(2024, 7, 12),
			date2: New(2024, 7, 13),
			want:  false,
		},
		{
			name:  "date1 is after date2",
			date1: New(2024, 7, 12),
			date2: New(2024, 7, 11),
			want:  true,
		},
		{
			name:  "date1 is same date2",
			date1: New(2024, 7, 12),
			date2: New(2024, 7, 12),
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.date1.After(tt.date2)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDate_Equal(t *testing.T) {
	tests := []struct {
		name  string
		date1 Date
		date2 Date
		want  bool
	}{
		{
			name:  "date1 is before date2",
			date1: New(2024, 7, 12),
			date2: New(2024, 7, 13),
			want:  false,
		},
		{
			name:  "date1 is after date2",
			date1: New(2024, 7, 12),
			date2: New(2024, 7, 11),
			want:  false,
		},
		{
			name:  "date1 is same date2",
			date1: New(2024, 7, 12),
			date2: New(2024, 7, 12),
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.date1.Equal(tt.date2)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDate_Add(t *testing.T) {
	type args struct {
		y, m, d int
	}
	type want struct {
		y int
		m time.Month
		d int
	}
	tests := []struct {
		name string
		date Date
		args args
		want want
	}{
		{
			name: "add dates",
			date: New(2024, 7, 12),
			args: args{y: 1, m: 2, d: 3},
			want: want{y: 2025, m: 9, d: 15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.date.Add(tt.args.y, tt.args.m, tt.args.d)
			assert.Equal(t, tt.want.y, got.Year())
			assert.Equal(t, tt.want.m, got.Month())
			assert.Equal(t, tt.want.d, got.Day())
		})
	}
}

func TestDate_AddYear(t *testing.T) {
	type want struct {
		y int
		m time.Month
		d int
	}
	tests := []struct {
		name string
		date Date
		args int
		want want
	}{
		{
			name: "add dates",
			date: New(2024, 7, 12),
			args: 1,
			want: want{y: 2025, m: 7, d: 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.date.AddYear(tt.args)
			assert.Equal(t, tt.want.y, got.Year())
			assert.Equal(t, tt.want.m, got.Month())
			assert.Equal(t, tt.want.d, got.Day())
		})
	}
}

func TestDate_AddMonth(t *testing.T) {
	type want struct {
		y int
		m time.Month
		d int
	}
	tests := []struct {
		name string
		date Date
		args int
		want want
	}{
		{
			name: "add dates",
			date: New(2024, 7, 12),
			args: 1,
			want: want{y: 2024, m: 8, d: 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.date.AddMonth(tt.args)
			assert.Equal(t, tt.want.y, got.Year())
			assert.Equal(t, tt.want.m, got.Month())
			assert.Equal(t, tt.want.d, got.Day())
		})
	}
}

func TestDate_AddDay(t *testing.T) {
	type want struct {
		y int
		m time.Month
		d int
	}
	tests := []struct {
		name string
		date Date
		args int
		want want
	}{
		{
			name: "add dates",
			date: New(2024, 7, 12),
			args: 1,
			want: want{y: 2024, m: 7, d: 13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.date.AddDay(tt.args)
			assert.Equal(t, tt.want.y, got.Year())
			assert.Equal(t, tt.want.m, got.Month())
			assert.Equal(t, tt.want.d, got.Day())
		})
	}
}
