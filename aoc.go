package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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
}

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
}

func main() {
	var (
		day1Set      = flag.NewFlagSet("day1", flag.ExitOnError)
		day1InputPtr = day1Set.String("input", "", "input file")
		day2Set      = flag.NewFlagSet("day2", flag.ExitOnError)
		day2InputPtr = day2Set.String("input", "", "input file")
	)
	if len(os.Args) < 2 {
		fmt.Printf("enter a solution day\n")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "day1":
		day1Set.Parse(os.Args[2:])
		day1(*day1InputPtr)
	case "day2":
		day2Set.Parse(os.Args[2:])
		day2(*day2InputPtr)
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
