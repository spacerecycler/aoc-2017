package main

import (
	"flag"
	"fmt"
	"os"
)

// Check if the passed in error is not empty and panic
func check(e error) {
	if e != nil {
		panic(e)
	}
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
