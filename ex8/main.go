package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime/debug"
	"strconv"
	"strings"
)

// Calculate gcdTwo of two numbers
func gcdTwo(a, b int) int {
	if a == 0 {
		return b
	}
	return gcdTwo(b%a, a)
}

// Calculate lcmTwo of two numbers
func lcmTwo(a, b int) int {
	return a * b / gcdTwo(a, b)
}

// Calculate the least common multiple of N numbers
func lcm(numbers []int) int {
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = lcmTwo(result, numbers[i])
	}
	return result
}

func main() {
	//Set max stack to 4GB
	debug.SetMaxStack(4000000000)
	//Open the file
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//Read a file into a string
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	letters := strings.Split(strings.ReplaceAll(strings.ReplaceAll(scanner.Text(), "L", "0"), "R", "1"), "")
	instructions := make([]int, 0)

	//Iterate over instructions and parse into ints
	for _, instruction := range letters {
		value, _ := strconv.Atoi(instruction)
		instructions = append(instructions, value)
	}

	scanner.Scan()
	inputs := []int{20777, 19199, 18673, 16043, 12361, 15517}
	fmt.Println(lcm(inputs))
	os.Exit(0)

	graph := map[string][]string{}

	for scanner.Scan() {
		input := scanner.Text()
		split := strings.Split(input, "=")
		node := strings.Trim(split[0], " ")
		re := regexp.MustCompile(`\w+`)
		matches := re.FindAllStringSubmatch(split[1], -1)
		graph[node] = []string{matches[0][0], matches[1][0]}
	}

	starts := make([]string, 0)
	for key := range graph {
		// If key ends with A
		if strings.HasSuffix(key, "A") {
			starts = append(starts, key)
		}
	}

	//Recursively traverse graph until node is found
	traverse := func(nodes []string, i int) []string {
		newNodes := make([]string, 0)
		// Iterate over current nodes and find the next
		for _, node := range nodes {
			newNodes = append(newNodes, graph[node][instructions[i]])
		}
		return newNodes
	}

	idx := 0
	current := []string{"XQA"}
	steps := 0

	for x := 0; x > -1; x = 0 {
		steps++
		current = traverse(current, idx)
		allZ := true
		for _, node := range current {
			if !strings.HasSuffix(node, "Z") {
				allZ = false
				break
			} else {
				fmt.Println(steps)
			}
		}
		if allZ {
			break
		}
		idx = (idx + 1) % len(instructions)
	}

	//Handle errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
