package day7_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/aphilas/aoc2022/day7"
)

const sampleInput = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func TestParseOutput(t *testing.T) {
	root := day7.ParseFS(strings.NewReader(sampleInput))
	fmt.Println(day7.SmallDirSum(root))
}
