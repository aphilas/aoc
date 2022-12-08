package day3_test

import (
	"strings"
	"testing"

	"github.com/aphilas/aoc2022/day3"
)

type Test struct {
	input string
	want  byte
}

var tests = []Test{
	{"vJrwpWtwJgWrhcsFMMfFFhFp", 16},
	{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", 38},
	{"PmmdzqPrVvPwwTWBwg", 42},
	{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", 22},
	{"ttgJtRGJQctTZtZT", 20},
	{"CrZsJsPPZsGzwwsLwLmpwMDw", 19},
}

func TestRucksackDuplicates(t *testing.T) {
	for _, tt := range tests {
		got := day3.RucksackDuplicates(tt.input)

		if got != tt.want {
			t.Errorf("got %q, want %d", got, tt.want)
		}
	}
}

func TestCampDuplicates(t *testing.T) {
	var input strings.Builder
	for _, r := range tests {
		_, err := input.WriteString(r.input + "\n")
		if err != nil {
			t.Error(err)
		}
	}

	const want = 157
	got := day3.CampDuplicates(strings.NewReader(input.String()))

	if got != want {
		t.Errorf("got %q, want %d", got, want)
	}
}

func TestGroupBadge(t *testing.T) {
	got := day3.CharPriority(day3.GroupBadge(tests[0].input, tests[1].input, tests[2].input))
	want := 18

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCampBadges(t *testing.T) {
	var input strings.Builder
	for _, r := range tests {
		_, err := input.WriteString(r.input + "\n")
		if err != nil {
			t.Error(err)
		}
	}

	got := day3.CampBadges(strings.NewReader(input.String()))
	want := 70

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
