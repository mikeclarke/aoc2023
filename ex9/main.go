package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

// Find differences in sequence of numbers
func findDifferences(numbers []int) []int {
	differences := make([]int, 0, len(numbers))
	for i := 0; i < len(numbers)-1; i++ {
		difference := numbers[i+1] - numbers[i]
		differences = append(differences, difference)
	}
	return differences
}

func findAllDifferences(input []int) [][]int {
	tree := [][]int{input}

	for i := 1; i < len(input); i++ {
		tree = append(tree, findDifferences(tree[i-1]))
	}
	return tree
}

// Function that returns the last value of an array of integers
func lastValue(input []int) int {
	return input[len(input)-1]
}

func buildPredictions(input [][]int) []int {
	predictions := []int{0}
	for i := len(input) - 1; i >= 0; i-- {
		predictions = append(predictions, lastValue(input[i])+lastValue(predictions))
	}
	return predictions
}

func buildPredictions2(input [][]int) []int {
	predictions := []int{0}
	for i := len(input) - 1; i >= 0; i-- {
		predictions = append(predictions, input[i][0]-lastValue(predictions))
	}
	return predictions
}

func main() {
	//Open the file
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0

	//Read the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		numberMatches := parseValues(s, `-?\d+`)
		total += lastValue(buildPredictions2(findAllDifferences(numberMatches)))
	}

	fmt.Println(total)

	//Handle errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
