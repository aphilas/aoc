package day6

import (
	"fmt"
	"strings"
)

func PacketMarker(input string) int {
	return FindMarker(input, 4-1)
}

func MessageMarker(input string) int {
	return FindMarker(input, 14-1)
}

// TODO: Fix tracking duplicates—fails part 1 in current state.
// n - Number of prev characters to store
func FindMarker(input string, n int) int {
	// dup stores the index of the second to last instance of the most recent duplicate char
	c, dup := byte(0), -1

	// Use a list for indexing and a hashmap for fast look ups
	w, h := make([]byte, n), make(map[byte]int)

	for i := 0; i < len(input); i++ {
		c = input[i]
		if _, ok := h[c]; !ok && i-n > dup {
			// char has no duplicate in the window and the window has no duplicates
			// fmt.Printf("%s %c \n", ListString(w), c)
			return i + 1
		}

		// Do not delete head's duplicate from map
		if h[w[i%n]] < dup && !(dup > -1 && dup >= i-n && w[i%n] == input[dup]) {
			delete(h, w[i%n])
		}

		if j, ok := h[c]; ok && j > dup {
			dup = j
		}

		// Replace head with current char
		h[c], w[i%n] = i, c
	}

	return 0
}

func ListString(c []byte) string {
	var s strings.Builder
	s.WriteString("[")
	for _, v := range c {
		if v != 0 {
			s.WriteString(fmt.Sprintf("%c ", v))
		}
	}
	s.WriteString("]")
	return s.String()
}

func MapString(c map[byte]int) string {
	var s strings.Builder
	s.WriteString("{")
	for k := range c {
		s.WriteString(fmt.Sprintf("%c ", k))
	}
	s.WriteString("}")
	return s.String()
}
