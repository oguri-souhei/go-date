package date

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func ExampleDate() {
	d := New(2023, time.July, 13)

	fmt.Println(d.Year())
	fmt.Println(d.Month())
	fmt.Println(d.Day())

	// Output:
	// 2023
	// July
	// 13
}

func TestNew(t *testing.T) {
	type args struct {
		y int
		m time.Month
		d int
	}
	type want struct {
		y int
		m time.Month
		d int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "init date",
			args: args{y: 2024, m: 7, d: 12},
			want: want{y: 2024, m: 7, d: 12},
		},
		{
			name: "when date not exist",
			args: args{y: 2024, m: 7, d: 32},
			want: want{y: 2024, m: 8, d: 1},
		},
		{
			name: "when leap day",
			args: args{y: 2024, m: 2, d: 29},
			want: want{y: 2024, m: 2, d: 29},
		},
		{
			name: "when 1 day before leap day",
			args: args{y: 2024, m: 2, d: 28},
			want: want{y: 2024, m: 2, d: 28},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			date := New(tt.args.y, tt.args.m, tt.args.d)
			assert.Equal(t, tt.want.y, date.Year())
			assert.Equal(t, tt.want.m, date.Month())
			assert.Equal(t, tt.want.d, date.Day())
		})
	}
}

func TestDate_Year(t *testing.T) {
	tests := []struct {
		name string
		date Date
		want int
	}{
		{
			name: "returns year",
			date: New(2024, 7, 12),
			want: 2024,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.date.Year()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDate_Month(t *testing.T) {
	tests := []struct {
		name string
		date Date
		want time.Month
	}{
		{
			name: "returns month",
			date: New(2024, 7, 12),
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.date.Month()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDate_Day(t *testing.T) {
	tests := []struct {
		name string
		date Date
		want int
	}{
		{
			name: "returns day",
			date: New(2024, 7, 12),
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.date.Day()
			assert.Equal(t, tt.want, got)
		})
	}
}
