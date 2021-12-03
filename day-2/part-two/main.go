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

	var horizontal int
	var depth int
	var aim int

	for scanner.Scan() {
		row := strings.Split(scanner.Text(), " ")

		direction := row[0]

		value, err := strconv.Atoi(row[1])
		if err != nil {
			panic(err)
		}

		switch direction {
		case "forward":
			horizontal += value
			depth += value * aim
		case "up":
			aim -= value
		case "down":
			aim += value
		}
	}

	fmt.Println(horizontal * depth)
}
