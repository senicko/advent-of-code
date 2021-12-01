package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var scans []int

	// read input file
	raw, err := os.ReadFile("./in.txt")
	if err != nil {
		log.Fatal("failed to read input file")
	}
	data := string(raw)

	// parse input file data to numbers
	for _, v := range strings.Split(data, "\n") {
		scan, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal("failed to convert to int: ", err)
		}
		scans = append(scans, scan)
	}

	// check if we can create at least one measurement sum
	if len(scans) < 3 {
		log.Fatal("too few scans in order to produce the result!")
	}

	// find all sums that are bigger then previous sum
	prev := scans[0] + scans[1] + scans[2]
	res := 0

	for i := 1; i < len(scans)-2; i++ {
		curr := scans[i] + scans[i+1] + scans[i+2]
		if curr > prev {
			res++
		}
		prev = curr
	}

	// print the result
	fmt.Println(res)
}
