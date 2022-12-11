package day6_test

import (
	"testing"

	"github.com/aphilas/aoc2022/day6"
)

type test struct {
	input         string
	packetMarker  int
	messageMarker int
}

var tests = []test{
	{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7, 19},
	{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5, 23},
	{"nppdvjthqldpwncqszvftbrmjlhg", 6, 23},
	{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10, 29},
	{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11, 26},
}

func TestPacketMarker(t *testing.T) {
	for i, tt := range tests {
		got := day6.PacketMarker(tt.input)
		if got != tt.packetMarker {
			t.Errorf("got %v want %v case %d", got, tt.packetMarker, i+1)
		}
	}
}

func TestMessageMarker(t *testing.T) {
	for _, tt := range tests {
		got := day6.MessageMarker(tt.input)
		if got != tt.messageMarker {
			t.Errorf("got %v want %v case %s...", got, tt.messageMarker, tt.input[:3])
		}
	}
}
