package main

func day3(input string) {

	// 1 + 8 = 9
	// 1 + 8 + 16 = 25
	// 1 + 8 + 16 + 32 = 57
	// 1 + 8 + 16 + 32 + 64 = 121
	// 1 (0, 0)
	// 2 (1, 0)
	// 3 (1, 1)
	// 4 (0, 1)
	// 5 (-1, 1)
	// 6 (-1, 0)
	// 7 (-1, -1)
	// 8 (0, -1)
	// 9 (1, -1)
	// 10 (2, -1)
	// 11 (2, 0)

	// first find ring
	//n := 8
	//ringNum := (((n - 1) / 8) + 1) / 2
	//mod := (n - 1) % 8
	// 8*((2^n)-1)
	// ring 0 8*0 max
	// ring 1 8*1 max
	// ring 2 8*3 max
	// ring 3 8*7 max
	// ring 4 8*15 max
	// then find location in ring

	//fmt.Printf("coords:  (%d, %d)", x, y)
	//fmt.Printf("distance: %d", d)
}
