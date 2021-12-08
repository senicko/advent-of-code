package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("../in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewScanner(file)

	var count int
	unique := [4]int{2, 4, 3, 7}

	for reader.Scan() {
		observation := strings.Split(reader.Text(), " | ")
		for _, digit := range strings.Split(observation[1], " ") {
			for _, u := range unique {
				if len(digit) == u {
					count++
					break
				}
			}
		}
	}

	fmt.Println(count)
}
