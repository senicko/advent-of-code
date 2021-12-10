package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("../in_test.txt")
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
	}

	scores := map[int32]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	var score int
	var stack []int32

	for scanner.Scan() {
		line := scanner.Text()

		for _, bracket := range line {
			found := false

			switch bracket {
			case '(', '[', '{', '<':
				stack = append(stack, bracket)
			default:
				if stack[len(stack)-1] == brackets[bracket] {
					stack = stack[:len(stack)-1]
				} else {
					found = true
					score += scores[bracket]
				}
			}

			if found {
				break
			}
		}
	}

	fmt.Println(score)
}
