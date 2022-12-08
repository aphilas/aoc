package day1_test

import (
	"fmt"
	"strings"
	"testing"

	p "github.com/aphilas/aoc2022/day1"
)

func TestCalCount(t *testing.T) {
	const input = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000

99999
`

	// const want = 24000
	const want = 99999
	got := p.CalCount(strings.NewReader(input))

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestRanker(t *testing.T) {
	inputs := [...]int{5, 6, 3, 1, 9, 23}

	r := p.NewRank()
	fmt.Println(r)

	for _, inp := range inputs {
		// fmt.Printf("input %d\n", inp)
		r.Insert(inp)
		// fmt.Println(r)
	}
}

func TestCalrank(t *testing.T) {
	const input = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

	const want = 45000
	got := p.CalRank(strings.NewReader(input))

	if got != want {
		t.Errorf("got %q, want %d", got, want)
	}
}
