package day4_test

import (
	"strings"
	"testing"

	"github.com/aphilas/aoc2022/day4"
)

const input = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`

// Find number of supersets

func TestOverlappingSections(t *testing.T) {
	const want = 2
	got := day4.FullOverlaps(strings.NewReader(input))

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCountOverlappingSections(t *testing.T) {
	const want = 4
	got := day4.PartialOverlaps(strings.NewReader(input))

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
