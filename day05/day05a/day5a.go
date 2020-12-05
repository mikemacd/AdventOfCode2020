package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:            true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
	})

	data := readInput()
	passes := map[int]struct {
		r    int
		c    int
		pass string
	}{}

	for _, v := range data {
		v = strings.TrimSpace(v)
		r, c, err := decodeBoardingPass(v)
		if err != nil {
			log.Error(err)
		}
		passes[r*8+c] = struct {
			r    int
			c    int
			pass string
		}{r, c, v}
	}

	maxID := 0
	for i, _ := range passes {
		if i > maxID {
			maxID = i
		}
	}

	fmt.Printf("\nMax pass ID: %v\n\n", maxID)

}

func decodeBoardingPass(pass string) (row, col int, err error) {
	rowBSP := pass[0:7]
	colBSP := pass[7:10]

	row = bsp(128, rowBSP)
	col = bsp(8, colBSP)

	return
}

func bsp(size int, bsp string) (pos int) {
	lower := 0
	upper := size
	for _, v := range bsp {
		if v == 'F' || v == 'L' {
			upper = upper - ((upper + 1 - lower) / 2)
		} else {
			lower = lower + ((upper + 1 - lower) / 2)
		}
	}
	rv := lower
	if upper-lower == 2 {
		rv = lower + 1
	}

	return rv
}

func readInput() (data []string) {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}
	f, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}
	lines := bytes.Split(f, []byte("\n"))

	for _, v := range lines {
		line := strings.TrimSpace(string(v))
		data = append(data, line)
	}

	return
}
