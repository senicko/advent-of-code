package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var scans []int

	raw, err := os.ReadFile("../in.txt")
	if err != nil {
		panic(err)
	}
	data := string(raw)

	for _, v := range strings.Split(data, "\n") {
		scan, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		scans = append(scans, scan)
	}

	if len(scans) < 3 {
		panic(err)
	}

	prev := scans[0] + scans[1] + scans[2]
	var result int

	for i := 1; i < len(scans)-2; i++ {
		curr := scans[i] + scans[i+1] + scans[i+2]
		if curr > prev {
			result++
		}
		prev = curr
	}

	fmt.Println(result)
}
