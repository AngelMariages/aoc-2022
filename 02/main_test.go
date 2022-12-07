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
			want: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSumOfScores(testInput, 1); got != tt.want {
				t.Errorf("getSumOfScores() = %v, want %v", got, tt.want)
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
			want: 12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSumOfScores(testInput, 2); got != tt.want {
				t.Errorf("getSumOfScores() = %v, want %v", got, tt.want)
			}
		})
	}
}
