package day4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func P2(f *os.File) {
	scanner := bufio.NewScanner(f)

	ans := 0
	counter := []int{1}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()

		splitString := strings.Split(line, ": ")

		cardSets := strings.Split(splitString[1], " | ")

		winningSet := strings.Split(cardSets[0], " ")
		yourSet := strings.Split(cardSets[1], " ")

		tempAns := 0
		for _, yourNumber := range yourSet {
			if yourNumber == "" {
				continue
			}

			for _, winningNumber := range winningSet {
				if winningNumber == "" {
					continue
				}

				if yourNumber == winningNumber {
					tempAns += 1
					break
				}
			}
		}

		for len(counter) <= i+tempAns {
			counter = append(counter, 1)
		}

		for j := i + 1; j <= i+tempAns; j += 1 {
			counter[j] += counter[i]
		}

		ans += counter[i]
		i += 1
	}

	fmt.Println(ans)
}
