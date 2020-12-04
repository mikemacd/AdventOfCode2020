package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

//	"github.com/davecgh/go-spew/spew"
)

type passports []passport

type passport struct {
	fields map[string]string
}

func main() {
	data := readInput()

	n := data.countValidPassports()
	fmt.Printf("Valid passports:%d out of %d \n",n,len(data))

//	spew.Dump(data)
}
func (p passports) countValidPassports() (num int) {
	for _,passport := range p {
		if passport.isValid() {
			num++
		}
	}
	return
}

func (p passport) isValid() (bool) {
	requiredFields := []string{"byr","iyr","eyr","hgt","hcl","ecl","pid"}
	permittedFields := []string{"byr","iyr","eyr","hgt","hcl","ecl","pid","cid"}

	for k,_ := range p.fields {
		permitted := false
		for _,f := range permittedFields {
			if k==f {
				permitted = true
			} 
		}
		if !permitted {
			// non permitted field
			// fmt.Printf("Non permitted field %s in %v\n", k, p )
			return false
		}
	}

	for _,f := range requiredFields {
		if _,ok := p.fields[f]; !ok {
			// fmt.Printf("Missing required field %s in %v\n", f, p )
			return false
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
				passportFields[string(pk)]=strings.Trim(string(pf), "\n\r")
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
