package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseValues(str string, regex string) []int {
	re := regexp.MustCompile(regex)
	matches := re.FindAllString(str, -1)
	result := make([]int, 0, len(matches))
	for _, match := range matches {
		value, _ := strconv.Atoi(match)
		result = append(result, value)
	}
	return result
}

// Generate a function that intersects two sets
func intersect(a []int, b []int) []int {
	m := make(map[int]bool)

	for _, item := range a {
		m[item] = true
	}

	result := make([]int, 0, len(m))

	for _, item := range b {
		if _, ok := m[item]; ok {
			result = append(result, item)
		}
	}

	return result
}

// 1: 2, 3, 4, 5
// 2: 3, 4
// 3: 4, 5

// Generate a function that recursively counts the number of paths from a given node to the end
func countPaths(graph map[int][]int, node int) int {
	if len(graph[node]) == 0 {
		return 1
	}

	total := 1
	for _, child := range graph[node] {
		total += countPaths(graph, child)
	}

	return total
}

// Generate a function that creates a range based on an offset and a length
func integerRange(start int, length int) []int {
	result := make([]int, length)
	for i := 0; i < length; i++ {
		result[i] = start + i + 1
	}
	return result
}

func main() {
	//Open the file
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0
	lineNum := 1
	results := make(map[int][]int)

	//Read the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		input := strings.Split(s, ":")
		values := strings.Split(input[1], "|")
		winnerNumbers := parseValues(values[0], `\d+`)
		candidateNumbers := parseValues(values[1], `\d+`)
		winners := len(intersect(winnerNumbers, candidateNumbers))
		results[lineNum] = integerRange(lineNum, winners)
		lineNum += 1
	}

	// Iterate over the keys of a map
	for key := range results {
		paths := countPaths(results, key)
		total += paths
	}

	//Handle errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(total)
}
