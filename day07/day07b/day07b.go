package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

type lineSet []string
type ruleSet map[string]map[string]int
type reverseRuleSet map[string]lineSet

func main() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:            true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
	})

	data := readInput()

	rules,ruletree := data.extractRules()

	c := ruletree.GetChildren("shiny gold")
spew.Dump(rules["shiny gold"])
	// -1 because the shiny gold bag can't be in itself
	fmt.Printf("Unique colours: %d\n", len(unique(c))-1)
}

func (tree reverseRuleSet) GetChildren(name string) (children []string) {
	children = []string{name}
	if len(tree[name]) == 0 {
		return
	}
	for _, v := range tree[name] {
		children = append(children, tree.GetChildren(v)...)
	}
	return
}
func unique(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
func (tree reverseRuleSet) Print() {

	keys := make([]string, 0, len(tree))
	for k := range tree {
		keys = append(keys, k)
	}
	sort.Strings(keys)

}
func (data lineSet) extractRules() (ruleSet,reverseRuleSet) {
	r := ruleSet{}
	rr := reverseRuleSet{}

	for _, v := range data {
		outerBagRE := regexp.MustCompile("((?P<name>.*?) bags? contain)+")
		innerBagRE := regexp.MustCompile("((?P<count>[0-9]+) (?P<name>.*?) bags?)+")

		outerBagNameIdx := outerBagRE.SubexpIndex("name")
		outerBagNameMatches := outerBagRE.FindStringSubmatch(v)
		if len(outerBagNameMatches) == 0 {
			fmt.Printf("No match for: %v\n", v)
		}

		outerBagName := outerBagNameMatches[outerBagNameIdx]

		submatches := innerBagRE.FindAllStringSubmatch(v, -1)

		if len(submatches) > 0 {
			for _, v := range submatches {
				count, err := strconv.Atoi(v[2])
				if err != nil {
					log.Errorf("could not convert '%s' to an int", v[2])
				}
				name := v[3]

				if r[outerBagName] == nil {
					r[outerBagName] = map[string]int{}
				}

				r[outerBagName][name] = count
				if _, ok := rr[name]; !ok {
					rr[name] = []string{}
				}

				rr[name] = append(rr[name], outerBagName)
			}
		}
	}

	return r,rr
}

func readInput() (data lineSet) {
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

	return lines
}
