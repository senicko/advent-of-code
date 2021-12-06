package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("../in.txt")
	if err != nil {
		panic(err)
	}

	cycles := make([]int, 9)
	for _, v := range strings.Split(string(file), ",") {
		fish, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		cycles[fish]++
	}

	for i := 0; i < 256; i++ {
		tmp := make([]int, 9)
		copy(tmp, cycles[1:])

		tmp[6] += cycles[0]
		tmp[8] += cycles[0]

		cycles = tmp
	}

	var result int
	for _, cycle := range cycles {
		result += cycle
	}

	fmt.Println(result)
}
