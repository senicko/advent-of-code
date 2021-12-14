package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("../in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	template := scanner.Text()
	scanner.Scan()

	rules := map[string]string{}
	for scanner.Scan() {
		rule := strings.Split(scanner.Text(), " -> ")
		rules[rule[0]] = rule[1]
	}

	pairs := map[string]int{}
	for i := 0; i < len(template)-1; i++ {
		pair := fmt.Sprintf("%s%s", string(template[i]), string(template[i+1]))
		pairs[pair]++
	}

	for i := 0; i < 40; i++ {
		newPairs := map[string]int{}

		for k, v := range pairs {
			a := string(k[0])
			b := string(k[1])

			if r, ok := rules[k]; ok {
				newPairs[fmt.Sprintf("%s%s", a, r)] += v
				newPairs[fmt.Sprintf("%s%s", r, b)] += v
			} else {
				newPairs[fmt.Sprintf("%s%s", a, b)] += v
			}
		}

		pairs = newPairs
	}

	count := map[string]int64{}
	for k, v := range pairs {
		count[string(k[0])] += int64(v)
		count[string(k[1])] += int64(v)
	}

	var min, max int64
	for _, v := range count {
		if min == 0 || v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}

	fmt.Println(int64(math.Ceil(float64(max)/2)) - int64(math.Ceil(float64(min)/2)))
}
