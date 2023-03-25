package day7

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type Node struct {
	name     string
	size     int
	children map[string]*Node // If nil, node is a file
	parent   *Node            // If nil, node is root
}

func NewDir(name string, parent *Node, children ...*Node) *Node {
	n := Node{name: name, parent: parent}

	n.children = make(map[string]*Node)
	for _, n := range children {
		n.children[n.name] = n
	}

	return &n
}

func NewFile(name string, parent *Node, size int) *Node {
	return &Node{name: name, parent: parent, size: size}
}

func ParseFS(input io.Reader) *Node {
	scanner := bufio.NewScanner(input)
	root := NewDir("", nil)
	wd := root

	for scanner.Scan() {
		var cmd, arg, sizeOrDir, name string

		txt := scanner.Text()
		if txt[0] == '$' {
			fmt.Sscanf(txt, "$ %s %s", &cmd, &arg)
		}

		switch cmd {
		case "cd":
			switch arg {
			case "/":
			case "..":
				wd = wd.parent
			default:
				for k, n := range wd.children {
					if k == arg {
						wd = n
						break
					}
				}
			}
		case "ls":
		default:
			fmt.Sscanf(txt, "%s %s", &sizeOrDir, &name)

			if sizeOrDir == "dir" {
				d := NewDir(name, wd)
				wd.children[name] = d
			} else {
				size, _ := strconv.Atoi(sizeOrDir)
				f := NewFile(name, wd, size)
				wd.children[name] = f
			}
		}

	}

	return root
}

// TODO: Memoize?
func DirSize(n *Node) int {
	if n.children == nil {
		return n.size
	}

	sum := 0
	for _, c := range n.children {
		sum += DirSize(c)
	}

	return sum
}

// Find total size per directory
// Find all directories with total size <= 10e5, find their sum
func SmallDirSum(n *Node) int {
	sum := 0

	// BFS traversal

	queue := []*Node{n}
	for len(queue) > 0 {
		v := queue[0]
		size := DirSize(v)
		if size < 10e5 {
			sum += size
		}

		for _, c := range v.children {
			if c.children != nil {
				// Directories only
				queue = append(queue, c)
			}
		}

		// queue[0] = nil
		queue = queue[1:]
	}

	return sum
}

func SpaceSaver(input io.Reader) int {
	root := ParseFS(input)
	return SmallDirSum(root)
}
