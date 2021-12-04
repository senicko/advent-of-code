package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	numbers []int
	boards  [][5][5]int
	marked  [][5][5]bool
	scanner *bufio.Scanner
)

func loadNumbers() {
	scanner.Scan()
	for _, v := range strings.Split(scanner.Text(), ",") {
		number, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, number)
	}
}

func loadBoards() {
	scanner.Scan()

	var (
		row   int
		board [5][5]int
	)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			row = 0
			boards = append(boards, board)
			marked = append(marked, [5][5]bool{})
			continue
		}

		for i := 0; i < len(line); i += 3 {
			number, err := strconv.Atoi(strings.TrimLeft(string(line[i])+string(line[i+1]), " "))
			if err != nil {
				panic(err)
			}
			board[row][i/3] = number
		}

		row++
	}

	boards = append(boards, board)
	marked = append(marked, [5][5]bool{})
}

func simulate() (int, error) {
	for _, number := range numbers {
		for i := 0; i < len(boards); i++ {
			var unmarkedSum int

			for j := 0; j < 5; j++ {
				for k := 0; k < 5; k++ {
					if boards[i][j][k] == number {
						marked[i][j][k] = true
					} else if marked[i][j][k] == false {
						unmarkedSum += boards[i][j][k]
					}
				}
			}

			for j := 0; j < 5; j++ {
				winX := true
				winY := true

				for k := 0; k < 5; k++ {
					if marked[i][j][k] == false {
						winX = false
					}

					if marked[i][k][j] == false {
						winY = false
					}
				}

				if winX || winY {
					return unmarkedSum * number, nil
				}
			}

		}
	}

	return 0, nil
}

func main() {
	file, err := os.Open("../in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)

	loadNumbers()
	loadBoards()

	result, err := simulate()
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
