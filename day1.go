package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func day1(input string) {
	str, err := ioutil.ReadFile(input)
	check(err)

	sum := 0
	lines := strings.Split(string(str), "\n")
	r := []rune(lines[0])
	for i := 0; i < len(r); i++ {
		next := i + 1
		if i == len(r)-1 {
			next = 0
		}
		if r[i] == r[next] {
			sum += int(r[i] - '0')
		}
	}
	fmt.Printf("%d\n", sum)
	sum = 0
	half := len(r) / 2
	for i := 0; i < len(r); i++ {
		next := (i + half) % len(r)
		if r[i] == r[next] {
			sum += int(r[i] - '0')
		}
	}
	fmt.Printf("%d\n", sum)
}
