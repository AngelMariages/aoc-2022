package main

import (
	_ "embed"
	"testing"
)

//go:embed test_input.txt
var testInput string

func Test_part1(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "example result for part 1",
			want: 24000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHighestCalories(getElfsCalories(testInput)); got != tt.want {
				t.Errorf("getHighestCalories() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "example result for part 2",
			want: 45000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSumOfHighest3Calories(getElfsCalories(testInput)); got != tt.want {
				t.Errorf("getSumOfHighest3Calories() = %v, want %v", got, tt.want)
			}
		})
	}
}
