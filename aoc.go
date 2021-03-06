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
		day3Set      = flag.NewFlagSet("day3", flag.ExitOnError)
		day3InputPtr = day3Set.String("input", "", "input number")
		day4Set      = flag.NewFlagSet("day4", flag.ExitOnError)
		day4InputPtr = day4Set.String("input", "", "input file")
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
	case "day3":
		day3Set.Parse(os.Args[2:])
		day3(*day3InputPtr)
	case "day4":
		day4Set.Parse(os.Args[2:])
		day4(*day4InputPtr)
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
