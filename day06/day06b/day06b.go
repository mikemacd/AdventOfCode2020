package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

type votes []votingRecord

type votingRecord struct {
	group   int
	person  int
	answers []string
}

func main() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:            true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
	})

	data := readInput()

	votesByGroupByPersonByQuestion := map[int]map[int]map[string]int{}
	votesByGroupByQuestion := map[int]map[string]int{}
	for _, r := range data {
		for _, a := range r.answers {
			if _, ok := votesByGroupByPersonByQuestion[r.group]; !ok {
				votesByGroupByPersonByQuestion[r.group] = map[int]map[string]int{}
			}
			if _, ok := votesByGroupByPersonByQuestion[r.group][r.person]; !ok {
				votesByGroupByPersonByQuestion[r.group][r.person] = map[string]int{}
			}
			votesByGroupByPersonByQuestion[r.group][r.person][a]++

			if _, ok := votesByGroupByQuestion[r.group]; !ok {
				votesByGroupByQuestion[r.group] = map[string]int{}
			}

			votesByGroupByQuestion[r.group][a]++

		}

	}

	sum := 0
	for g, gv := range votesByGroupByPersonByQuestion {
		l := len(gv)
		for q := range votesByGroupByQuestion[g] {
			if votesByGroupByQuestion[g][q] == l {
				sum++
			}
		}
	}
	fmt.Printf("\nSum:%d\n\n", sum)
}

func readInput() (data votes) {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}
	f, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Panicf("can't read file: %s", os.Args[1])
	}

	contents := strings.ReplaceAll(string(f), "\r\n", "\n")

	lines := strings.Split(contents, "\n")

	group := 0
	person := 0
	for _, line := range lines {

		// new group starting
		if len(line) == 0 {
			data = append(data, votingRecord{group: group, person: person, answers: strings.Split(line, "")})
			person = 0
			group++
			continue
		}
		data = append(data, votingRecord{group: group, person: person, answers: strings.Split(line, "")})

		person++
	}

	return
}
