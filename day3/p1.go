package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
	"unicode/utf8"
)

type Symbol struct {
	Value string
	Row   int
	Col   int
}

type Number struct {
	Value int
	Start int
	End   int
	Row   int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func appendDigit(digits []Number, row int, col int, tempDigit string) []Number {
	digit, err := strconv.Atoi(tempDigit)

	if err != nil {
		panic(err)
	}

	dx, dy := max(0, col-utf8.RuneCountInString(tempDigit)-1), col
	return append(digits, Number{Value: digit, Start: dx, End: dy, Row: row})
}

func P1(f *os.File) {
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
				if rune != '.' {
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

	ans := 0
	for _, digit := range digits {
		for _, symbol := range symbols {
			if symbol.Col > digit.End || symbol.Col < digit.Start {
				continue
			}

			if abs(symbol.Row-digit.Row) > 1 {
				continue
			}

			ans += digit.Value
			break
		}
	}

	fmt.Println(ans)
}
