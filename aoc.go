package main

import (
	"flag"
	"fmt"
	"strings"
	"io/ioutil"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	inputPtr := flag.String("input", "", "input file")
	flag.Parse()
	fmt.Printf("Input: %v\n", *inputPtr)
	str, err := ioutil.ReadFile(*inputPtr)
	check(err)

	sum := 0
	lines := strings.Split(string(str),"\n")
	r := []rune(lines[0])
	for i := 0; i < len(r); i++ {
		next := i+1;
		if(i == len(r) - 1) {
			next = 0
		}
		if(r[i] == r[next]) {
			sum += int(r[i] - '0')
		}
	}
	fmt.Printf("%d\n", sum)
}