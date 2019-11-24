package zellers

import (
	"fmt"
	"testing"
)

func TestCongruence(t *testing.T) {
	tests := []struct {
		month int
		day   int
		year  int
		want  string
	}{
		{
			month: 99,
			day:   99,
			year:  9999,
			want:  "Saturday",
		},
		{
			month: 12,
			day:   23,
			year:  1951,
			want:  "Sunday",
		},
		{
			month: 10,
			day:   28,
			year:  2019,
			want:  "Monday",
		},
		{
			month: 1,
			day:   15,
			year:  2041,
			want:  "Tuesday",
		},
		{
			month: 2,
			day:   23,
			year:  1853,
			want:  "Wednesday",
		},
		{
			month: 1,
			day:   30,
			year:  2025,
			want:  "Thursday",
		},
		{
			month: 8,
			day:   3,
			year:  2096, // leap year
			want:  "Friday",
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("test-%d-%d-%d", tc.month, tc.day, tc.year), func(t *testing.T) {
			got := Congruence(tc.month, tc.day, tc.year)
			if got != tc.want {
				t.Errorf("fail got %v want %v", got, tc.want)
			}
		})
	}
}
