package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findMostCommon(bits []string, bit int) byte {
	var result [2]int

	for i := 0; i < len(bits); i++ {
		if bits[i][bit] == '0' {
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

func findOxygenNumber(numbers []string) string {
	for i := 0; i < len(numbers[0]); i++ {
		mostCommon := findMostCommon(numbers, i)

		var filtered []string

		for j := 0; j < len(numbers); j++ {
			if numbers[j][i] == mostCommon {
				filtered = append(filtered, numbers[j])
			}
		}

		numbers = filtered

		if len(numbers) == 1 {
			break
		}
	}

	return numbers[0]
}

func findCO2Number(numbers []string) string {
	for i := 0; i < len(numbers[0]); i++ {
		leastCommon := flip(findMostCommon(numbers, i))

		var filtered []string

		for j := 0; j < len(numbers); j++ {
			if numbers[j][i] == leastCommon {
				filtered = append(filtered, numbers[j])
			}
		}

		numbers = filtered

		if len(numbers) == 1 {
			break
		}
	}

	return numbers[0]
}

func main() {
	raw, err := os.ReadFile("../in.txt")
	if err != nil {
		panic(err)
	}

	oxygenBinaryNumbers := strings.Split(string(raw), "\n")
	CO2BinaryNumbers := strings.Split(string(raw), "\n")

	oxygenNumber, err := strconv.ParseInt(findOxygenNumber(oxygenBinaryNumbers), 2, 64)
	if err != nil {
		panic(err)
	}

	CO2Number, err := strconv.ParseInt(findCO2Number(CO2BinaryNumbers), 2, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println(oxygenNumber * CO2Number)
}
