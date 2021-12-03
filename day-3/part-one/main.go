package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numbers []string

func findMostCommon(bit int) byte {
	var result [2]int

	for i := 0; i < len(numbers); i++ {
		if numbers[i][bit] == '0' {
			result[0]++
		} else {
			result[1]++
		}
	}

	if result[0] > result[1] {
		return '0'
	}

	return '1'
}

func flip(bit byte) byte {
	if bit == '1' {
		return '0'
	}

	return '1'
}

func main() {
	raw, err := os.ReadFile("../in.txt")
	if err != nil {
		panic(err)
	}

	numbers = strings.Split(string(raw), "\n")

	var gammaBinary string
	var epsilonBinary string

	for i := 0; i < len(numbers[0]); i++ {
		mostCommon := findMostCommon(i)

		gammaBinary += string(mostCommon)
		epsilonBinary += string(flip(mostCommon))
	}

	gammaRate, err := strconv.ParseInt(gammaBinary, 2, 64)
	if err != nil {
		panic(err)
	}

	epsilonRate, err := strconv.ParseInt(epsilonBinary, 2, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println(gammaRate * epsilonRate)
}
