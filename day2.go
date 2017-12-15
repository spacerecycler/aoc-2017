package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func day2(input string) {
	str, err := ioutil.ReadFile(input)
	check(err)

	sum := 0
	lines := strings.Split(string(str), "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		numbers := strings.Split(line, "\t")
		fn, err := strconv.Atoi(numbers[0])
		check(err)
		min := fn
		max := fn
		for _, s := range numbers[1:] {
			n, err := strconv.Atoi(s)
			check(err)
			if n < min {
				min = n
			}
			if n > max {
				max = n
			}
		}
		sum += max - min
	}
	fmt.Printf("%d\n", sum)
	sum = 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		strings := strings.Split(line, "\t")
		numbers := make([]int, 0, len(strings))
		for _, s := range strings {
			n, err := strconv.Atoi(s)
			check(err)
			numbers = append(numbers, n)
		}
		for i, n1 := range numbers {
			for j, n2 := range numbers {
				if i == j {
					continue
				}
				if n1%n2 == 0 {
					sum += n1 / n2
				}
			}
		}
	}
	fmt.Printf("%d\n", sum)
}
