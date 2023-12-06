package day6

import (
	"fmt"
	"math"
)

func P1() {
	// races := [][]float64{{7, 9}, {15, 40}, {30, 200}}
	// races := [][]float64{{46, 347}, {82, 1522}, {84, 1406}, {79, 1471}}
	races := [][]float64{{46828479, 347152214061471}}

	ans := 1.0
	for _, race := range races {

		b := race[0]
		c := -race[1]

		delta := math.Sqrt(b*b + 4*c)
		minValue := (b - delta) * .5
		maxValue := (b + delta) * .5

		ans *= math.Ceil(maxValue) - math.Floor(minValue) - 1
	}

	fmt.Printf("%.0f", ans)

}
