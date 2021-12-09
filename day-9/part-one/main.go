package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var board [][]int

func isLowest(x, y int) bool {
	checks := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	ok := true
	for _, check := range checks {
		if y+check[0] >= 0 &&
			y+check[0] < len(board) &&
			x+check[1] >= 0 &&
			x+check[1] < len(board[y]) &&
			board[y][x] >= board[y+check[0]][x+check[1]] {
			ok = false
		}
	}

	return ok
}

func main() {
	file, err := os.Open("../in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var row []int

		for _, v := range strings.Split(scanner.Text(), "") {
			num, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			row = append(row, num)
		}

		board = append(board, row)
	}

	var points []int
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			lowest := isLowest(x, y)
			if lowest {
				points = append(points, board[y][x])
			}
		}
	}

	var answer int
	for _, p := range points {
		answer += p + 1
	}

	fmt.Println(answer)
}
