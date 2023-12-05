package day5

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type MapItem struct {
	Src int64
	Dst int64
	Rng int64
}

func P1(f *os.File) {
	scanner := bufio.NewScanner(f)

	// seeds
	scanner.Scan()
	seedLine := scanner.Text()
	seeds := strings.Split(seedLine, " ")[1:]

	maps := [7][]MapItem{}
	i := -1
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if !unicode.IsDigit(rune(line[0])) {
			i += 1
			continue
		}

		values := strings.Split(line, " ")
		int_values := MapItem{}

		for j, val := range values {
			tmpInt, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				panic(err)
			}

			if j == 1 {
				int_values.Src = tmpInt
			} else if j == 0 {
				int_values.Dst = tmpInt
			} else if j == 2 {
				int_values.Rng = tmpInt
			}
		}

		maps[i] = append(maps[i], int_values)

	}

	var finalAns int64 = math.MaxInt64
	for _, strSeed := range seeds {
		seed, err := strconv.ParseInt(strSeed, 10, 64)

		if err != nil {
			panic(err)
		}

		var ans int64 = seed
		for _, mapy := range maps {

			for _, v := range mapy {
				if ans < v.Src || ans >= (v.Src+v.Rng) {
					continue
				}

				ans = (ans - v.Src) + v.Dst
				break
			}
		}

		finalAns = min(ans, finalAns)
	}

	fmt.Println(finalAns)
}
