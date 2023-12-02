package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func Day1p1(f *os.File) {
	scanner := bufio.NewScanner(f)

	ans := 0
	for scanner.Scan() {
		line := scanner.Text()

		sumChar := ""
		lastSeen := ' '
		for _, char := range line {
			if unicode.IsDigit(char) {
				if utf8.RuneCountInString(sumChar) == 0 {
					sumChar += string(char)
				}

				lastSeen = char
			}
		}

		if utf8.RuneCountInString(sumChar) == 1 {
			sumChar += string(lastSeen)
		}

		num, err := strconv.Atoi(sumChar)

		if err != nil {
			panic(err)
		}

		ans += num
	}

	fmt.Printf("%d", ans)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file", err)
	}
}
