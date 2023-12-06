package main

import (
	"adventofcode2023/day6"
	"os"
)

func main() {
	f, err := os.Open("./day5/input.txt")

	if err != nil {
		panic(err)
	}

	// day5.P2(f)
	day6.P1()

	defer f.Close()
}
