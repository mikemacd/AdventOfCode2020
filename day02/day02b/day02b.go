package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type passwordList []password
type password struct {
	posA     int
	posB     int
	letter   byte
	password string
}

func main() {
	data := readInput()

	n := data.countValidPasswords()

	fmt.Printf("Number of valid passwords: %d\n", n)

}

func (p passwordList) countValidPasswords() (num int) {
	for _, pass := range p {
		if pass.isValid() {
			num++
		}
	}
	return
}

func (p password) isValid() bool {
	if len(p.password) < p.posA-1 || len(p.password) < p.posB-1 {
		return false
	}

	return ((p.password[p.posA-1] == p.letter && p.password[p.posB-1] != p.letter) ||
		(p.password[p.posA-1] != p.letter && p.password[p.posB-1] == p.letter))
}

func readInput() (passwords passwordList) {

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
	for i, line := range lines {
		pass := password{}

		_, err := fmt.Sscanf(string(line), "%d-%d %c: %s\n", &pass.posA, &pass.posB, &pass.letter, &pass.password)
		if err != nil {
			log.Fatalf("Error (%v) while parsing line %d: %v\n", err, i, line)
		}
		passwords = append(passwords, pass)
	}

	return
}
