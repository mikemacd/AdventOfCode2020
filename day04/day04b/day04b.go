package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type passports []passport

type passport struct {
	fields map[string]string
}

func main() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:            true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
	})
	//	log.SetReportCaller(true)

	data := readInput()

	n := data.countValidPassports()
	fmt.Printf("\nValid passports:%d out of %d \n\n", n, len(data))

	//	spew.Dump(data)
}
func (p passports) countValidPassports() (num int) {
	for _, passport := range p {
		if passport.isValid() {
			num++
		}
	}
	return
}

func (p passport) hasAllRequiredFields() bool {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, f := range requiredFields {
		if _, ok := p.fields[f]; !ok {
			// log.Debugf("Missing required field %s in %v\n", f, p )
			return false
		}
	}

	return true
}

func (p passport) hasOnlyPermittedFields() bool {
	permittedFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}

	for k, _ := range p.fields {
		permitted := false
		for _, f := range permittedFields {
			if k == f {
				permitted = true
			}
		}
		if !permitted {
			// log.Debugf("Non permitted field %s in %v\n", k, p )
			return false
		}
	}

	return true
}

func (p passport) isValid() bool {
	if !p.hasAllRequiredFields() {
		log.Debugf("Does not have All permitted fields: %v\n", p)
		return false
	}
	if !p.hasOnlyPermittedFields() {
		log.Debugf("Does not have only permitted fields: %v\n", p)
		return false
	}

	for k, v := range p.fields {
		v = strings.TrimSpace(v)
		switch k {
		case "byr":
			{
				yr, err := strconv.Atoi(v)
				if err != nil {
					log.Debugf("ERROR %v\n", err)
					return false
				}
				if !(yr >= 1920 && yr <= 2002) {
					log.Debugf("Bad byr: %v %v\n", yr, v)
					return false
				}
			}
		case "iyr":
			{
				yr, err := strconv.Atoi(v)
				if err != nil {
					log.Debugf("ERROR %v\n", err)
					return false
				}
				if !(yr >= 2010 && yr <= 2020) {
					log.Debugf("Bad iyr: %v %v\n", yr, v)
					return false
				}
			}
		case "eyr":
			{
				yr, err := strconv.Atoi(v)
				if err != nil {
					log.Debugf("ERROR %v\n", err)
					return false
				}
				if !(yr >= 2020 && yr <= 2030) {
					log.Debugf("Bad eyr: %v %v\n", yr, v)
					return false
				}

			}
		case "hgt":
			{
				re := regexp.MustCompile(`(?P<height>\d+)(?P<unit>(cm|in))`)
				matches := re.FindStringSubmatch(v)
				if len(matches) == 0 {
					log.Debugf("BAD HGT: cant parse %v\n", v)
					return false
				}
				heightIdx := re.SubexpIndex("height")
				unitIdx := re.SubexpIndex("unit")
				if heightIdx < 0 {
					log.Debugf("Bad hgt hgtIdx: m:'%v' hi:'%v' ui:'%v' v:'%v'\n", matches, heightIdx, unitIdx, v)
					return false

				}
				if unitIdx < 0 {
					log.Debugf("Bad hgt unitIdx: m:'%v' hi:'%v' ui:'%v' v:'%v'\n", matches, heightIdx, unitIdx, v)
					return false
				}

				height, err := strconv.Atoi(strings.TrimSpace(matches[heightIdx]))
				if err != nil {
					log.Debugf("ERROR %v\n", err)
					return false
				}
				switch matches[unitIdx] {
				case "cm":
					{
						if !(height >= 150 && height <= 193) {
							log.Debugf("Bad cm hgt: %v \n", height)
							return false
						}
					}
				case "in":
					{
						if !(height >= 59 && height <= 76) {
							log.Debugf("Bad in hgt: %v \n", height)
							return false
						}

					}
				}

			}
		case "hcl":
			{
				re := regexp.MustCompile(`^(#[0-9a-f]{6})$`)
				matches := re.FindString(v)
				if len(matches) == 0 {
					log.Debugf("Bad hcl: '%v' %v \n", matches, v)
					return false
				}
			}
		case "ecl":
			{
				re := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
				matches := re.FindString(v)
				if len(matches) == 0 {
					log.Debugf("Bad ecl: %v\n", v)
					return false
				}
			}
		case "pid":
			{
				re := regexp.MustCompile(`^\d{9}$`)
				matches := re.FindString(v)
				if len(matches) == 0 {
					log.Debugf("Bad PID: '%v' '%v' in %v -- RE:%v\n", matches, v, p, re)
					return false
				}
			}
		case "cid":
			{
				// nop, ignored.
			}

		}
	}
	return true

}

func readInput() (data passports) {
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

	passportFields := map[string]string{}
	for _, line := range lines {
		if len(line) <= 1 {
			for pk, pf := range passportFields {
				passportFields[string(pk)] = strings.Trim(string(pf), "\n\r")
			}
			data = append(data, passport{fields: passportFields})
			passportFields = map[string]string{}
			continue
		}
		fields := bytes.Split(line, []byte(" "))
		for _, fld := range fields {
			v := bytes.Split(fld, []byte(":"))
			passportFields[string(v[0])] = string(v[1])
		}
	}

	data = append(data, passport{fields: passportFields})

	return
}
