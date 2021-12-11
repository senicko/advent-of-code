package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	flashes    int
	board      [][]int
	neighbours = [][2]int{
		{0, 1}, {0, -1},
		{1, 0}, {-1, 0},
		{1, 1}, {1, -1},
		{-1, 1}, {-1, -1},
	}
)

func flash(x, y int, flashed [][]bool) {
	if !flashed[y][x] {
		board[y][x]++
	}

	if board[y][x] < 10 || flashed[y][x] {
		return
	}

	flashed[y][x] = true
	board[y][x] = 0
	flashes++

	for _, neighbour := range neighbours {
		nx := x + neighbour[0]
		ny := y + neighbour[1]

		if nx < len(board[y]) && nx >= 0 && ny < len(board) && ny >= 0 {
			flash(nx, ny, flashed)
		}
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

		for _, v := range strings.Split(scanner.Text(), "") {
			energy, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			row = append(row, energy)
		}

		board = append(board, row)
	}

	for i := 0; i < 100; i++ {
		flashed := make([][]bool, len(board))
		for i := range flashed {
			flashed[i] = make([]bool, len(board[i]))
		}

		for y := 0; y < len(board); y++ {
			for x := 0; x < len(board[y]); x++ {
				flash(x, y, flashed)
			}
		}
	}

	fmt.Print(flashes)
}
