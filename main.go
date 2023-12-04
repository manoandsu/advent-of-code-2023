package main

import (
	"adventofcode2023/day4"
	"os"
)

func main() {
	f, err := os.Open("./day4/input.txt")

	if err != nil {
		panic(err)
	}

	day4.P2(f)

	defer f.Close()
}
