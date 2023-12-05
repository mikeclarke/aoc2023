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

// Generate a function that chunks an array of integers into chunks of a given size
func chunk(size int) func([]int) [][]int {
	return func(values []int) [][]int {
		chunked := [][]int{}
		for i := 0; i < len(values); i += size {
			chunked = append(chunked, values[i:i+size])
		}
		return chunked
	}
}

func applyRules(rules [][][]int, seed int) int {
	value := seed
	//Iterate the list of rules
	for _, rule := range rules {
		//Iterate the list of rule chunks
		for _, chunk := range rule {
			//If the seed is within the range of the chunk
			if value >= chunk[1] && value < chunk[1]+chunk[2] {
				// Update the seed value
				value = (value - (chunk[1] - chunk[0]))
				break
			}
		}
	}
	return value
}

func main() {
	//Open the file
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := ""

	//Read a file into a string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input += scanner.Text() + "\n"
	}

	//Split a string on empty lines
	split := strings.Split(input, "\n\n")

	//Parse a list of integers from a string
	seeds := parseValues(split[0], `\d+`)

	rules := make([][][]int, 7)
	rules[0] = chunk(3)(parseValues(split[1], `\d+`))
	rules[1] = chunk(3)(parseValues(split[2], `\d+`))
	rules[2] = chunk(3)(parseValues(split[3], `\d+`))
	rules[3] = chunk(3)(parseValues(split[4], `\d+`))
	rules[4] = chunk(3)(parseValues(split[5], `\d+`))
	rules[5] = chunk(3)(parseValues(split[6], `\d+`))
	rules[6] = chunk(3)(parseValues(split[7], `\d+`))

	//Iterate the list of seeds
	minimumLocation := seeds[0]
	for _, seed := range seeds {
		seedCandidate := applyRules(rules, seed)
		if seedCandidate < minimumLocation {
			minimumLocation = seedCandidate
		}
	}

	//Handle errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(minimumLocation)
}
