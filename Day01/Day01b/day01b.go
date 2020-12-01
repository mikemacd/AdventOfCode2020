package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	//	"github.com/davecgh/go-spew/spew"
)

func main() {
	numbers := readInput()

	v := threeSum(numbers, 2020)
	if len(v) > 0 {
		fmt.Printf("%d *%d * %d == %d\n", numbers[v[0]], numbers[v[1]], numbers[v[2]], numbers[v[0]]*numbers[v[1]]*numbers[v[2]])
		os.Exit(0)
	}
	fmt.Printf("Nothing found\n")
	os.Exit(0)

}

func threeSum(numbers []int, target int) []int {
	m := make(map[int]int)
	for idx, num := range numbers {
		if n := twoSum(numbers, target-num); len(n) > 0 {
			return []int{n[0], n[1], idx}
		}
		m[num] = idx
	}
	return []int{}
}

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)
	for idx, num := range numbers {
		if v, found := m[target-num]; found {
			return []int{v, idx}
		}
		m[num] = idx
	}
	return []int{}
}

func readInput() []int {
	var numbers []int

	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return []int{}
	}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}

	lines := bytes.Split(data, []byte("\n"))
	for i, line := range lines {
		num, err := strconv.Atoi(string(line))
		if err != nil {
			log.Fatalf("Can't parse number on line %d: %v\n", i, line)
		}
		numbers = append(numbers, num)
	}

	return numbers
}
