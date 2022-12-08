package day2

import (
	"bufio"
	"fmt"
	"io"
)

// Score/round - { 1 A | 2 B | 3 C } + { 0 L | 3 D | 6 W }
// Your wins - C A, B C, A B

func FixedChoice(input io.Reader) int {
	score := 0
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		text := scanner.Text()

		opp, you := text[0]-'A'+1, text[2]-'X'+1 // Convert to 1, 2, 3
		sig := opp*10 + you                      // Concatenate opp,you

		if opp == you {
			score += 3
		} else if sig == 31 || sig == 23 || sig == 12 {
			score += 6
		}

		score += int(you)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("invalid: %s", err)
	}

	return score
}

// Score/round - { 1 A | 2 B | 3 C } + { 0 L | 3 D | 6 W }
// Your wins - C A, B C, A B

func FixedEnding(input io.Reader) int {
	score := 0
	scanner := bufio.NewScanner(input)

	// Loss moves for opp moves: 1, 2, 3
	loss := [...]byte{3, 1, 2}

	for scanner.Scan() {
		text := scanner.Text()

		opp, you := text[0]-'A'+1, text[2]-'X'+1 // Convert to 1, 2, 3
		score += int(you-1) * 3                  // X,Y,Z = 1,2,3 = L,D,W = 0,3,6 pts

		if you == 2 {
			// Draw
			score += int(opp)
		} else if you == 1 {
			// TODO: Replace with modular arithmetic
			score += int(loss[opp-1])
		} else {
			// Win hierachy: 1, 2, 3
			score += int(opp%3) + 1
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("invalid: %s", err)
	}

	return score
}
