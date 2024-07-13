package date

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func ExampleDate_String() {
	d := New(2023, time.July, 13)

	fmt.Println(d.String())
	// Output: 2023-07-13
}

func ExampleDate_Format() {
	d := New(2023, time.July, 13)

	fmt.Println(d.Format("02-01-2006"))
	// Output: 13-07-2023
}

func ExampleParse() {
	d, err := Parse("13-07-2023")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(d)
	}
	// Output: 2023-07-13
}

func TestDate_String(t *testing.T) {
	tests := []struct {
		name string
		date Date
		want string
	}{
		{
			name: "",
			date: New(2024, 7, 13),
			want: "2024-07-13",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.date.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDate_Format(t *testing.T) {
	tests := []struct {
		name   string
		date   Date
		layout string
		want   string
	}{
		{
			name:   "returns formatted string",
			date:   New(2024, 7, 13),
			layout: "02 Jan 2006",
			want:   "13 Jul 2024",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.date.Format(tt.layout)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    Date
		wantErr bool
	}{
		{
			name: "parse date-only format",
			args: "2024-07-12",
			want: New(2024, 7, 12),
		},
		{
			name:    "invalid format",
			args:    "hogehuga",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args)
			if !tt.wantErr {
				assert.Equal(t, tt.want, got)
				assert.Nil(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
