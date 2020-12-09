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

	votesByGroupByQuestion := map[int]map[string]int{}

	for _, r := range data {
		for _, v := range r.answers {
			if _, ok := votesByGroupByQuestion[r.group]; !ok {
				votesByGroupByQuestion[r.group] = map[string]int{}
			}
			votesByGroupByQuestion[r.group][v]++
		}

	}
	sum := 0
	for _, gv := range votesByGroupByQuestion {
		sum += len(gv)
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
	//	re := regexp.MustCompile("(?P<G>(([^\r\n]+)\r?\n)+)(\r?\n)(?P<G2>([^\r\n]+))")

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
