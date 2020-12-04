package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

type position struct {
	X int
	Y int
}

type forest struct {
	trees  map[position]bool
	width  int
	height int
}

func main() {
	data := readInput()

	// data.print()

	n := data.countTrees(3, 1)
	fmt.Printf("Number of trees: %d\n", n)

}

func readInput() (f forest) {
	f = forest{height: 0, width: 0, trees: map[position]bool{}}

	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}
	lines := bytes.Split(data, []byte("\n"))

	if len(lines) > f.height {
		f.height = len(lines)
	}
	for y, line := range lines {
		if len(line) > f.width {
			f.width = len(line)
		}

		for x, char := range line {
			var tree bool
			switch char {
			case '.':
				tree = false
			case '#':
				tree = true
			default:
				continue
			}
			f.trees[position{X: x, Y: y}] = tree
		}
	}

	return
}

func (f forest) countTrees(x, y int) (num int) {
	current := position{0, 0}

	// Loop until we hit the bottom
	for current.Y <= f.height+1 {
		if f.trees[current] {
			num++
		}

		// move according to our slope
		current.X += x
		current.Y += y

		// check if we need to scroll around to the left hand side
		current.X = current.X % f.width
	}

	return
}

func (f forest) print() {
	fmt.Printf("dimensions %d x %d\n", f.width, f.height)

	fmt.Printf("    ")
	for x := 0; x < f.width; x++ {
		fmt.Printf("%d", x/10)
	}
	fmt.Println()
	fmt.Printf("    ")
	for x := 0; x < f.width; x++ {
		fmt.Printf("%d", x%10)
	}
	fmt.Println()

	for y := 0; y < f.height; y++ {
		fmt.Printf("%3d ", y)
		for x := 0; x < f.width; x++ {
			t := '.'
			if f.trees[position{X: x, Y: y}] {
				t = 'T'
			}
			fmt.Printf("%s", string(t))
		}
		fmt.Println("")
	}
}
