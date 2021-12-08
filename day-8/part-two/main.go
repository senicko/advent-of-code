package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func findInt(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func findUint8(slice []uint8, value uint8) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func getIndex(slice []uint8, value uint8) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func digitFromLength(length int) int {
	switch length {
	case 2:
		return 1
	case 3:
		return 7
	case 4:
		return 4
	case 7:
		return 8
	}
	return -1
}

func main() {
	file, err := os.Open("../in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewScanner(file)

	unique := []int{2, 3, 4, 7}
	mappings := map[int][]int{
		0: {0, 1, 2, 4, 5, 6},
		1: {2, 5},
		2: {0, 2, 3, 4, 6},
		3: {0, 2, 3, 5, 6},
		4: {1, 2, 3, 5},
		5: {0, 1, 3, 5, 6},
		6: {0, 1, 3, 4, 5, 6},
		7: {0, 2, 5},
		8: {0, 1, 2, 3, 4, 5, 6},
		9: {0, 1, 2, 3, 5, 6},
	}

	var answer int

	for reader.Scan() {
		var patterns, outputs []string
		digits := make([]uint8, 7)
		note := strings.Split(reader.Text(), " | ")

		for _, pattern := range strings.Split(note[0], " ") {
			patterns = append(patterns, pattern)
		}

		for _, digit := range strings.Split(note[1], " ") {
			outputs = append(outputs, digit)
		}

		sort.Slice(patterns, func(i, j int) bool {
			return len(patterns[i]) < len(patterns[j])
		})

		for _, pattern := range patterns {
			if findInt(unique, len(pattern)) {
				var v, i int
				digit := digitFromLength(len(pattern))

				for v < len(mappings[digit]) && i < len(pattern) {
					if digits[mappings[digit][v]] == 0 && !findUint8(digits, pattern[i]) {
						digits[mappings[digit][v]] = pattern[i]
						v++
						i++
					} else if findUint8(digits, pattern[i]) && digits[mappings[digit][v]] == 0 {
						i++
					} else if !findUint8(digits, pattern[i]) && digits[mappings[digit][v]] != 0 {
						v++
					} else {
						i++
						v++
					}
				}
			}
		}

		var leftCorner, midLeftCorner, right [2]int

		for _, pattern := range patterns {
			for i := 0; i < len(pattern); i++ {
				if pattern[i] == digits[4] {
					leftCorner[0]++
				} else if pattern[i] == digits[6] {
					leftCorner[1]++
				} else if pattern[i] == digits[1] {
					midLeftCorner[0]++
				} else if pattern[i] == digits[3] {
					midLeftCorner[1]++
				} else if pattern[i] == digits[2] {
					right[0]++
				} else if pattern[i] == digits[5] {
					right[1]++
				}
			}
		}

		if leftCorner[0] > leftCorner[1] {
			tmp := digits[4]
			digits[4] = digits[6]
			digits[6] = tmp
		}

		if midLeftCorner[0] > midLeftCorner[1] {
			tmp := digits[1]
			digits[1] = digits[3]
			digits[3] = tmp
		}

		if right[0] > right[1] {
			tmp := digits[2]
			digits[2] = digits[5]
			digits[5] = tmp
		}

		var result string

		//for _, c := range digits {
		//	fmt.Print(string(c))
		//}
		//fmt.Println()

		for _, output := range outputs {
			var number []int

			for i := 0; i < len(output); i++ {
				index := getIndex(digits, output[i])
				number = append(number, index)
			}

			sort.Ints(number)

			//fmt.Print(number)

			for num, mapping := range mappings {
				if len(mapping) < len(number) {
					continue
				}

				ok := true
				for i := 0; i < len(number); i++ {
					if mapping[i] != number[i] {
						ok = false
					}
				}

				if ok {
					result += strconv.Itoa(num)
					//fmt.Println(num)
					break
				}
			}

		}

		a, err := strconv.Atoi(result)
		if err != nil {
			panic(err)
		}
		answer += a
	}

	fmt.Println(answer)
}
