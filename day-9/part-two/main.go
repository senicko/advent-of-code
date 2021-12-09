package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var board [][]int
var visited [][]bool

func findBasin(x, y int, basin *[]int) {
	visited[y][x] = true

	point := board[y][x]
	if point == 9 {
		return
	}
	*basin = append(*basin, point)

	if x+1 < len(board[y]) && !visited[y][x+1] {
		findBasin(x+1, y, basin)
	}

	if x-1 >= 0 && !visited[y][x-1] {
		findBasin(x-1, y, basin)
	}

	if y+1 < len(board) && !visited[y+1][x] {
		findBasin(x, y+1, basin)
	}

	if y-1 >= 0 && !visited[y-1][x] {
		findBasin(x, y-1, basin)
	}
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
		var visits []bool

		for _, v := range strings.Split(scanner.Text(), "") {
			num, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			row = append(row, num)
			visits = append(visits, false)
		}

		board = append(board, row)
		visited = append(visited, visits)
	}

	var basins [][]int

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			if !visited[y][x] {
				var basin []int
				findBasin(x, y, &basin)
				basins = append(basins, basin)
			}
		}
	}

	sort.Slice(basins, func(i, j int) bool {
		return len(basins[i]) > len(basins[j])
	})

	answer := 1
	for i := 0; i < 3; i++ {
		answer *= len(basins[i])
	}

	fmt.Println(answer)
}
