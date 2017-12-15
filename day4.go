package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

func day4(input string) {
	str, err := ioutil.ReadFile(input)
	check(err)

	sum := 0
	lines := strings.Split(string(str), "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		valid := true
		words := strings.Split(line, " ")
	loop:
		for i, w1 := range words {
			for j, w2 := range words {
				if i == j {
					continue
				}
				if strings.Compare(w1, w2) == 0 {
					valid = false
					break loop
				}
			}
		}
		if valid {
			sum++
		}
	}
	fmt.Printf("%d\n", sum)
	sum = 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		valid := true
		words := strings.Split(line, " ")
	loop2:
		for i, w1 := range words {
			for j, w2 := range words {
				if i == j {
					continue
				}
				r1 := []rune(w1)
				r2 := []rune(w2)
				if len(r1) == len(r2) {
					m1 := make(map[rune]int)
					m2 := make(map[rune]int)
					for k := 0; k < len(r1); k++ {
						m1[r1[k]]++
						m2[r2[k]]++
					}
					if reflect.DeepEqual(m1, m2) {
						valid = false
						break loop2
					}
				}
			}
		}
		if valid {
			sum++
		}
	}
	fmt.Printf("%d\n", sum)
}
