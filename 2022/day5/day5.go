package day5

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Bytes per item
const size = 4

type Crane = func(a [][]byte, from, to, count int)

func MoveCrates(input io.ReadSeeker, cols, rows int, c Crane) string {
	a, chunk, rowSize := make([][]byte, cols), make([]byte, size), cols*size
	for i := range a {
		a[i] = make([]byte, 0)
	}

	for cursor := rowSize*rows - size; cursor > -1; cursor -= size {
		input.Seek(int64(cursor), io.SeekStart)
		input.Read(chunk)
		// row := (cursor + 1) / rowSize
		col := ((cursor + 1) % rowSize) / size
		if chunk[1] != ' ' {
			a[col] = append(a[col], chunk[1])
		}
	}

	// Skip stacks, stack numbering, and a new line
	input.Seek(int64(rowSize*(rows+1)+1), io.SeekStart)
	var count, from, to int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		_, err := fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &count, &from, &to)
		if err != nil {
			panic(err)
		}
		c(a, from-1, to-1, count)
	}

	var (
		tops strings.Builder
		char byte
	)
	for _, s := range a {
		if len(s) > 0 {
			char = s[len(s)-1]
		} else {
			char = ' '
		}
		err := tops.WriteByte(char)
		if err != nil {
			panic(err)
		}
	}

	return tops.String()
}

func MoveCratesSingly(input io.ReadSeeker, cols, rows int) string {
	mover := func(a [][]byte, from, to, count int) {
		for ; count > 0; count-- {
			item := a[from][len(a[from])-1]
			a[from] = a[from][:len(a[from])-1]
			a[to] = append(a[to], item)
		}
	}
	return MoveCrates(input, cols, rows, mover)
}

func MoveCratesGrouped(input io.ReadSeeker, cols, rows int) string {
	mover := func(a [][]byte, from, to, count int) {
		a[to] = append(a[to], a[from][len(a[from])-count:]...)
		a[from] = a[from][:len(a[from])-count]
	}
	return MoveCrates(input, cols, rows, mover)
}

func PrintCargo(a [][]byte) {
	for _, s := range a {
		for _, c := range s {
			fmt.Printf("%c ", c)
		}
		fmt.Printf("\n")
	}
}
