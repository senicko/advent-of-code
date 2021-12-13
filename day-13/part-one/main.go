package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var points [][2]int
	var folds [][2]int

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		coords := strings.Split(scanner.Text(), ",")

		x, err := strconv.Atoi(coords[0])
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(coords[1])
		if err != nil {
			panic(err)
		}

		points = append(points, [2]int{x, y})
	}

	for scanner.Scan() {
		fold := strings.Split(strings.Split(scanner.Text(), " ")[2], "=")

		coord, err := strconv.Atoi(fold[1])
		if err != nil {
			panic(err)
		}

		switch fold[0] {
		case "x":
			folds = append(folds, [2]int{coord, 0})
		case "y":
			folds = append(folds, [2]int{0, coord})
		}
	}

	var maxX, maxY int
	for _, point := range points {
		if point[0] > maxX {
			maxX = point[0]
		}

		if point[1] > maxY {
			maxY = point[1]
		}
	}

	var board [][]bool
	for i := 0; i < maxY+1; i++ {
		var row []bool
		for j := 0; j < maxX+1; j++ {
			row = append(row, false)
		}
		board = append(board, row)
	}

	for _, point := range points {
		board[point[1]][point[0]] = true
	}

	for y := folds[0][1]; y < maxY+1; y++ {
		for x := folds[0][0]; x < maxX+1; x++ {
			fY := y
			fX := x

			if folds[0][1] > 0 {
				fY = folds[0][1] - (y - folds[0][1])
			}

			if folds[0][0] > 0 {
				fX = folds[0][0] - (x - folds[0][0])
			}

			board[fY][fX] = board[y][x] || board[fY][fX]
		}
	}

	if folds[0][1] > 0 {
		board = board[:folds[0][1]]
	}

	if folds[0][0] > 0 {
		for i, row := range board {
			board[i] = row[:folds[0][0]]
		}
	}

	var answer int
	for _, row := range board {
		for _, dot := range row {
			if dot {
				answer++
			}
		}
	}

	fmt.Println(answer)
}
