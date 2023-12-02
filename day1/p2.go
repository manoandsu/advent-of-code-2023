package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func Day1p2(f *os.File) {
	scanner := bufio.NewScanner(f)

	ans := 0
	words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for scanner.Scan() {
		line := scanner.Text()

		sumChar := ""
		lastSeen := ' '
		spelled := ""

		for _, char := range line {
			spelled += string(char)
			if unicode.IsDigit(char) {
				lastSeen = char
			}

			for i, word := range words {
				idx := strings.Index(spelled, word)

				if idx != -1 {
					lastSeen = rune(strconv.Itoa(i + 1)[0])
					spelled = spelled[idx+utf8.RuneCountInString(word)-1:]
				}
			}

			if utf8.RuneCountInString(sumChar) == 0 && lastSeen != ' ' {
				sumChar += string(lastSeen)
			}
		}

		for utf8.RuneCountInString(sumChar) < 2 {
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
