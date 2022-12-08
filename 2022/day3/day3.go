package day3

import (
	"bufio"
	"fmt"
	"io"
)

// firstHalf|secondHalf
// a-z 1-26, A-Z 27-52
// Find item appearing in both compartments

func RucksackDuplicates(line string) byte {
	aMap := make(map[byte]bool)
	for i := 0; i < len(line)/2; i++ {
		aMap[line[i]] = true
	}

	for i := len(line) / 2; i < len(line); i++ {
		if _, ok := aMap[line[i]]; ok {
			return line[i]
		}
	}

	return 0
}

func CharPriority(c byte) int {
	if c > 'Z' {
		return int(c) - 'a' + 1
	} else {
		return int(c) - 'A' + 27
	}
}

func CampDuplicates(input io.Reader) int {
	total := 0
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		text := scanner.Text()
		total += CharPriority(RucksackDuplicates(text))
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("invalid: %s", err)
	}

	return total
}

// Groups of 3
// All three Elves have shared badge
// < 3 Elves have an additional item

func GroupBadge(lines ...string) byte {
	set := make(map[rune]bool)
	flag := false

	for i, l := range lines {
		for _, r := range l {
			v, ok := set[r]

			if i == len(lines)-1 && v && ok {
				return byte(r)
			}

			if !flag || ok {
				set[r] = flag
			}
		}

		flag = !flag
	}

	return 0
}

func CampBadges(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	chunk := make([]string, 3)
	total := 0

	for i := 0; scanner.Scan(); i = (i + 1) % 3 {
		text := scanner.Text()
		chunk[i] = text

		if i == 2 {
			// Assume input is always in chunks of 3
			total += CharPriority(GroupBadge(chunk...))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("invalid: %s", err)
	}

	return total
}
