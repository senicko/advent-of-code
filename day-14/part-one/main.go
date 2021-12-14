package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("../in_test.txt")
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

	counts := map[rune]int{}
	for _, c := range template {
		counts[c] = 0
	}

	for s := 0; s < 10; s++ {
		var product string

		for i := 0; i < len(template)-1; i++ {
			var match string

			match += string(template[i])
			match += string(template[i+1])

			for rule, element := range rules {
				if rule == match {
					match = match[:1] + element + match[1:]
					break
				}
			}

			product += match[:2]
		}

		product += string(template[len(template)-1])
		template = product
	}

	for _, c := range template {
		counts[c]++
	}

	var min, max int
	for _, v := range counts {
		if v > max {
			max = v
		}

		if min == 0 || min > v {
			min = v
		}
	}

	fmt.Println(max - min)
}
