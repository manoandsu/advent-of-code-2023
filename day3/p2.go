package day3

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

func P2(f *os.File) {
	scanner := bufio.NewScanner(f)

	row := 0
	symbols := []Symbol{}
	digits := []Number{}

	for scanner.Scan() {
		line := scanner.Text()

		tempDigit := ""
		for col, rune := range line {
			if unicode.IsDigit(rune) {
				tempDigit += string(rune)
			} else {
				if rune == '*' {
					symbol := Symbol{Value: string(rune), Row: row, Col: col}
					symbols = append(symbols, symbol)
				}

				if tempDigit != "" {
					digits = appendDigit(digits, row, col, tempDigit)
					tempDigit = ""
				}
			}

			if col == utf8.RuneCountInString(line)-1 && tempDigit != "" {
				digits = appendDigit(digits, row, col+1, tempDigit)
				tempDigit = ""
			}
		}

		row += 1
	}

	var ans int64 = 0
	for _, symbol := range symbols {
		var values int64 = 1
		count := 0
		for _, digit := range digits {
			if symbol.Col > digit.End || symbol.Col < digit.Start {
				continue
			}

			if abs(symbol.Row-digit.Row) > 1 {
				continue
			}

			values *= int64(digit.Value)
			count += 1

			if count > 2 {
				break
			}
		}

		if count == 2 {
			ans += values
		}
	}

	fmt.Println(ans)
}
