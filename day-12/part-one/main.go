package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var smallCaveName = regexp.MustCompile(`[a-z]`)

type Node struct {
	ID          string
	connections []*Node
}

func copyMap(m map[string]bool) map[string]bool {
	copy := map[string]bool{}
	for key, value := range m {
		copy[key] = value
	}
	return copy
}

func getAllPaths(node *Node, visited map[string]bool) int {
	if node.ID == "end" {
		return 1
	}

	if smallCaveName.MatchString(node.ID) {
		visited[node.ID] = true
	}

	var paths int
	for _, n := range node.connections {
		if !visited[n.ID] {
			paths += getAllPaths(n, copyMap(visited))
		}
	}

	return paths
}

func main() {
	file, err := os.Open("../in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	caves := map[string]*Node{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		connection := strings.Split(scanner.Text(), "-")

		var from, to *Node

		if node, ok := caves[connection[0]]; ok {
			from = node
		} else {
			from = &Node{ID: connection[0]}
			caves[connection[0]] = from
		}

		if node, ok := caves[connection[1]]; ok {
			to = node
		} else {
			to = &Node{ID: connection[1]}
			caves[connection[1]] = to
		}

		from.connections = append(from.connections, to)
		to.connections = append(to.connections, from)
	}

	visited := map[string]bool{}
	for id := range caves {
		visited[id] = false
	}
	visited["start"] = true

	var paths int
	for _, node := range caves["start"].connections {
		paths += getAllPaths(node, copyMap(visited))
	}
	fmt.Println(paths)
}
