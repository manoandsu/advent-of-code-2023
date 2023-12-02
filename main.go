package main

import (
	"adventofcode2023/day2"
	"os"
)

func main() {
	f, err := os.Open("./day2/input.txt")

	if err != nil {
		panic(err)
	}

	day2.P2(f)

	defer f.Close()
}
