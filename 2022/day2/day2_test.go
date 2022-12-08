package day2_test

import (
	"strings"
	"testing"

	p "github.com/aphilas/aoc2022/day2"
)

func TestFixedChoice(t *testing.T) {
	const input = `A Y
B X
C Z
`

	const want = 15
	got := p.FixedChoice(strings.NewReader(input))

	if got != want {
		t.Errorf("got %q, want %d", got, want)
	}
}

func TestFixedEnding(t *testing.T) {
	const input = `A Y
B X
C Z
`

	want := 12
	got := p.FixedEnding(strings.NewReader(input))

	if got != want {
		t.Errorf("got %q, want %d", got, want)
	}
}
