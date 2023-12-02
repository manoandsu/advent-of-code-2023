package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func P2(f *os.File) {
	scanner := bufio.NewScanner(f)

	var ans int64 = 0
	for scanner.Scan() {
		countedGems := []int64{0, 0, 0}
		line := scanner.Text()

		lineParts := strings.Split(line, ": ")

		for _, gemSetList := range strings.Split(lineParts[1], ";") {
			for _, gemSet := range strings.Split(gemSetList, ",") {
				gemData := strings.Split(strings.Trim(gemSet, " "), " ")

				gemCount, err := strconv.ParseInt(gemData[0], 10, 64)
				if err != nil {
					panic(err)
				}

				gemColor := gemData[1]
				idx := utf8.RuneCountInString(gemColor) - 3
				countedGems[idx] = max(countedGems[idx], gemCount)
			}
		}

		ans += countedGems[0] * countedGems[1] * countedGems[2]
	}

	fmt.Println(ans)
}
