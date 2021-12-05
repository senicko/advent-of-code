package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMinMax(v [2][2]int) (int, int, int, int) {
	minX := v[1][0]
	maxX := v[0][0]

	if minX > maxX {
		minX, maxX = maxX, minX
	}

	minY := v[1][1]
	maxY := v[0][1]

	if minY > maxY {
		minY, maxY = maxY, minY
	}

	return minX, maxX, minY, maxY
}

func getPoints(p string) []string {
	var (
		points []string
		line   [2][2]int
	)

	for i, point := range strings.Split(p, " -> ") {
		coords := strings.Split(point, ",")

		x, err := strconv.Atoi(coords[0])
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(coords[1])
		if err != nil {
			panic(err)
		}

		line[i][0] = x
		line[i][1] = y
	}

	minX, maxX, minY, maxY := getMinMax(line)

	var point string

	if minX == maxX {
		for i := minY; i <= maxY; i++ {
			point = fmt.Sprintf("%d;%d", minX, i)
			points = append(points, point)
		}
	} else if minY == maxY {
		for i := minX; i <= maxX; i++ {
			point = fmt.Sprintf("%d;%d", i, minY)
			points = append(points, point)
		}
	} else {
		base := line[0]
		if base[0] > minX {
			base = line[1]
		}

		slope := -1
		if base[1]-maxY < 0 {
			slope = 1
		}

		for i := 0; i <= maxX-minX; i++ {
			point = fmt.Sprintf("%d;%d", base[0]+i, base[1]+(slope*i))
			points = append(points, point)
		}
	}

	return points
}

func main() {
	file, err := os.Open("../in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result int
	set := map[string]int{}

	for scanner.Scan() {
		points := getPoints(scanner.Text())

		for _, point := range points {
			set[point]++
			if set[point] == 2 {
				result++
			}
		}
	}

	fmt.Println(result)
}
