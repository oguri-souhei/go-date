package date

import (
	"database/sql/driver"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDate_Value(t *testing.T) {
	tests := []struct {
		name string
		date Date
		want driver.Value
	}{
		{
			name: "returns driver.Value",
			date: New(2024, 7, 13),
			want: "2024-07-13",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.date.Value()
			assert.Equal(t, tt.want, got)
			assert.Nil(t, err)
		})
	}
}

func TestDate_Scan(t *testing.T) {
	tests := []struct {
		name    string
		date    Date
		args    any
		want    Date
		wantErr bool
	}{
		{
			name: "when src is string",
			date: Date{},
			args: "2024-07-13",
			want: New(2024, 7, 13),
		},
		{
			name:    "when src is string and invalid format",
			date:    Date{},
			args:    "hoge",
			wantErr: true,
		},
		{
			name: "when src is time",
			date: Date{},
			args: time.Date(2024, 7, 14, 0, 0, 0, 0, time.UTC),
			want: New(2024, 7, 14),
		},
		{
			name: "when src is []byte",
			date: Date{},
			args: []byte("2024-07-17"),
			want: New(2024, 7, 17),
		},
		{
			name:    "when src is []byte and inavlid format",
			date:    Date{},
			args:    []byte("hoge"),
			wantErr: true,
		},
		{
			name:    "when src is int",
			date:    Date{},
			args:    10,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.date.Scan(tt.args)
			if !tt.wantErr {
				assert.Equal(t, tt.want, tt.date)
				assert.Nil(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
