package day5_test

import (
	"bytes"
	"testing"

	"github.com/aphilas/aoc2022/day5"
)

const (
	test1 = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`
	test2 = `[M]                     [N] [Z]    
[F]             [R] [Z] [C] [C]    
[C]     [V]     [L] [N] [G] [V]    
[W]     [L]     [T] [H] [V] [F] [H]
[T]     [T] [W] [F] [B] [P] [J] [L]
[D] [L] [H] [J] [C] [G] [S] [R] [M]
[L] [B] [C] [P] [S] [D] [M] [Q] [P]
[B] [N] [J] [S] [Z] [W] [F] [W] [R]
 1   2   3   4   5   6   7   8   9 

move 5 from 3 to 6
move 2 from 2 to 5
move 1 from 9 to 1
move 1 from 3 to 1
move 5 from 7 to 5
move 2 from 9 to 8
move 1 from 2 to 8
move 1 from 4 to 2
`
)

func TestMoveCratesSingly(t *testing.T) {
	// day5.OrganizeCrates(*bytes.NewReader([]byte(test2)), 9, 8)

	const want = "CMZ"
	got := day5.MoveCratesSingly(bytes.NewReader([]byte(test1)), 3, 3)
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestMoveCratesGrouped(t *testing.T) {
	const want = "MCD"
	got := day5.MoveCratesGrouped(bytes.NewReader([]byte(test1)), 3, 3)
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
