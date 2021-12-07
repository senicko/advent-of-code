package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getMean(positions []int) int {
	var sum int
	for _, v := range positions {
		sum += v
	}
	return int(math.Round(float64(sum) / float64(len(positions))))
}

func getUsedFuel(steps int) int {
	var fuel int
	for i := 1; i <= steps; i++ {
		fuel += i
	}
	return fuel
}

func main() {
	in, err := os.ReadFile("../in.txt")
	if err != nil {
		panic(err)
	}

	var positions []int
	for _, v := range strings.Split(string(in), ",") {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		positions = append(positions, num)
	}

	mean := getMean(positions)

	var cases [7]int
	cases[0] = int(math.Round(float64(mean) - 1))
	cases[2] = mean
	cases[4] = int(math.Round(float64(mean) + 1))

	fuel := -1
	for _, c := range cases {
		var candidate int

		for _, v := range positions {
			steps := int(math.Abs(float64(v - c)))
			candidate += getUsedFuel(steps)
		}

		if candidate < fuel || fuel == -1 {
			fuel = candidate
		}
	}

	fmt.Println(fuel)
}
