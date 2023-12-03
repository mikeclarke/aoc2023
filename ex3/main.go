package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// Generate a function that multiples all integers of an array and returns the product
func multiply(numbers []int) int {
	product := 1
	for _, num := range numbers {
		product *= num
	}
	return product
}

func main() {
	//Open the file
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0
	lineNum := 0
	numbers := make([][4]int, 0)
	symbols := map[[2]int]string{}
	gearMap := map[[2]int][]int{}

	numberRegex := regexp.MustCompile(`\d+`)
	symbolRegex := regexp.MustCompile(`[^0-9\.]`)

	//Read the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		numberMatches := numberRegex.FindAllStringIndex(s, -1)
		symbolMatches := symbolRegex.FindAllStringIndex(s, -1)

		for _, match := range numberMatches {
			//Substring of a string at index
			num := s[match[0]:match[1]]
			//Convert string to int
			val, _ := strconv.Atoi(num)

			tuple := [4]int{val, lineNum, match[0], len(num)}
			numbers = append(numbers, tuple)
		}

		for _, match := range symbolMatches {
			symbols[[2]int{lineNum, match[0]}] = s[match[0]:match[1]]
		}
		lineNum += 1
	}

	for _, tuple := range numbers {
		val := tuple[0]
		lineNum := tuple[1]
		index := tuple[2]
		length := tuple[3]

		// Above number
		for i := index - 1; i <= index+length; i++ {
			for j := lineNum - 1; j <= lineNum+1; j++ {
				if symbols[[2]int{j, i}] == "*" {
					// Append the value to the gearMap
					gearMap[[2]int{j, i}] = append(gearMap[[2]int{j, i}], val)
				}
			}
		}
	}

	// Iterate the gearMap and find the values that have more than one value
	for _, values := range gearMap {
		if len(values) > 1 {
			total += multiply(values)
		}
	}

	//Handle errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(total)
}
