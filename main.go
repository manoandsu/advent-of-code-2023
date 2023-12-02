package main

import (
	"adventofcode2023/day1"
	"os"
)

func main() {
	f, err := os.Open("./day1/input.txt")

	if err != nil {
		panic(err)
	}

	day1.Day1p1(f)

	defer f.Close()
}
