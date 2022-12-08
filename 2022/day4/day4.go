package day4

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

// Elf - sections IDs
// Assignments per Elf pair: 2-4, 6-8

type Counter func(chunk []int) int

func FullOverlaps(input io.Reader) int {
	comp := func(c []int) int {
		if (c[0] <= c[2] && c[1] >= c[3]) || (c[0] >= c[2] && c[1] <= c[3]) {
			return 1
		}

		return 0
	}

	return CountSections(input, comp)
}

func PartialOverlaps(input io.Reader) int {
	comp := func(c []int) int {
		// Overlap: a stars before b but b stars before a ends, b starts before a but a starts before b ends
		if (c[0] <= c[2] && c[2] <= c[1]) || (c[2] <= c[0] && c[0] <= c[3]) {
			return 1
		}

		return 0
	}

	return CountSections(input, comp)
}

func CountSections(input io.Reader, comp Counter) (count int) {
	scanner := bufio.NewScanner(input)
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] < '0' || data[i] > '9' {
				return i + 1, data[:i], nil
			}
		}

		if !atEOF {
			return 0, nil, nil
		}
		return 0, data, bufio.ErrFinalToken
	}

	scanner.Split(split)
	chunk := make([]int, 4)
	for i := 0; scanner.Scan(); i = (i + 1) % 4 {
		// Read 4 section IDs, then do the comparison
		v, _ := strconv.Atoi(scanner.Text())
		chunk[i] = v
		if i == 3 {
			count += comp(chunk)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("invalid input: %s", err)
	}

	return
}
