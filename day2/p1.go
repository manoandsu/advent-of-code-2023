package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func P1(f *os.File) {
	scanner := bufio.NewScanner(f)

	// red, blue, green
	gameIdx := 1
	ans := 0
	maxGems := []int{12, 14, 13}
	for scanner.Scan() {
		countedGems := []int{0, 0, 0}
		line := scanner.Text()

		lineParts := strings.Split(line, ": ")

		for _, gemSetList := range strings.Split(lineParts[1], ";") {
			for _, gemSet := range strings.Split(gemSetList, ",") {
				gemData := strings.Split(strings.Trim(gemSet, " "), " ")

				gemCount, err := strconv.Atoi(gemData[0])
				if err != nil {
					panic(err)
				}

				gemColor := gemData[1]
				idx := utf8.RuneCountInString(gemColor) - 3
				countedGems[idx] = max(countedGems[idx], gemCount)
			}
		}

		possible := true
		for i, v := range countedGems {
			if v > maxGems[i] {
				possible = false
				break
			}
		}

		if possible {
			ans += gameIdx
		}

		gameIdx += 1
	}

	fmt.Println(ans)
}
