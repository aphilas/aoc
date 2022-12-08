package day1

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func CalCount(input io.Reader) int {
	var max, curr int

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			if curr > max {
				// fmt.Printf("new max: %d\n", curr)
				max = curr
			}

			// fmt.Printf("total: %d\n", curr)
			curr = 0
			continue
		}

		value, err := strconv.Atoi(text)
		if err != nil {
			fmt.Printf("cannot parse %q as int\n", text)
			continue
		}

		curr += value
		// fmt.Printf("current: %d\n", curr)
	}

	// Last total after line without subsequent empty line
	if curr != 0 && curr > max {
		max = curr
		// fmt.Printf("current: %d\n", curr)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("invalid: %s", err)
	}

	return max
}

type Rank struct {
	items [3]int
}

func NewRank() *Rank {
	r := Rank{
		items: [3]int{-1, -1, -1},
	}
	return &r
}

func (r *Rank) Insert(n int) {
	if n < r.items[2] {
		return
	}

	for i := 0; i < 3; i++ {
		// Insert n at position i
		// and shift the rest of elements right
		if n >= r.items[i] {
			temp := n
			for j := i; j < 3; j++ {
				// Swap item and temp
				r.items[j], temp = temp, r.items[j]
			}
			break
		}
	}
}

func (r *Rank) String() string {
	return fmt.Sprintf("%v", r.items)
}

// Idea: After counting the calories in each bag, use modified insertion sort

func CalRank(input io.Reader) int {
	var curr int

	r := NewRank()
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			// fmt.Printf("curr: %d\n", curr)
			r.Insert(curr)
			curr = 0
			continue
		}

		value, err := strconv.Atoi(text)
		if err != nil {
			fmt.Printf("cannot parse %q as int\n", text)
			continue
		}

		// fmt.Printf("value: %d\n", value)
		curr += value
	}

	// Last total after line without subsequent empty line
	if curr != 0 {
		// fmt.Printf("curr: %d\n", curr)
		r.Insert(curr)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("invalid: %s", err)
	}

	// fmt.Println(r)
	sum := 0
	for _, r := range r.items {
		if r == -1 {
			continue
		}

		sum += r
	}

	return sum
}
