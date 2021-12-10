package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("../in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	brackets := map[int32]int32{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}

	scores := map[int32]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	var results []int

	for scanner.Scan() {
		var (
			stack []int32
			score int
		)

		line := scanner.Text()
		corrupted := false

		for _, bracket := range line {

			switch bracket {
			case '(', '[', '{', '<':
				stack = append(stack, bracket)
			default:
				if stack[len(stack)-1] == brackets[bracket] {
					stack = stack[:len(stack)-1]
				} else {
					corrupted = true
				}
			}

			if corrupted {
				break
			}
		}

		if !corrupted {
			for i := len(stack) - 1; i >= 0; i-- {
				score *= 5
				score += scores[brackets[stack[i]]]
			}
			results = append(results, score)
		}
	}

	sort.Ints(results)
	fmt.Println(results[len(results)/2])
}
