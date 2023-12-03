package main

import (
	"adventofcode2023/day3"
	"os"
)

func main() {
	f, err := os.Open("./day3/input.txt")

	if err != nil {
		panic(err)
	}

	day3.P2(f)

	defer f.Close()
}
