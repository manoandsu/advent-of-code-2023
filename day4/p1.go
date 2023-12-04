package day4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func P1(f *os.File) {
	scanner := bufio.NewScanner(f)

	var ans int64 = 0
	for scanner.Scan() {
		line := scanner.Text()

		splitString := strings.Split(line, ": ")

		cardSets := strings.Split(splitString[1], " | ")

		winningSet := strings.Split(cardSets[0], " ")
		yourSet := strings.Split(cardSets[1], " ")

		var tempAns int64 = 0
		for _, yourNumber := range yourSet {
			if yourNumber == "" {
				continue
			}

			for _, winningNumber := range winningSet {
				if winningNumber == "" {
					continue
				}

				if yourNumber == winningNumber {
					if tempAns == 0 {
						tempAns = 1
					} else {
						tempAns *= 2
					}
					break
				}
			}
		}

		ans += tempAns
	}

	fmt.Println(ans)
}
