package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// Generate a function that returns an array of integers given a start and length
func rangeOf(start int, length int) []int {
	result := make([]int, length)
	for i := 0; i < length; i++ {
		result[i] = start + i
	}
	return result
}

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

// Generate a function that removes whitespace from a string
func removeWhitespace(str string) string {
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(str, "")
}

func main() {
	//Open the file
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	waysToWin := 1

	//Read a file into a string
	scanner := bufio.NewScanner(file)

	scanner.Scan()

	durations := parseValues(removeWhitespace(scanner.Text()), `\d+`)

	scanner.Scan()
	records := parseValues(removeWhitespace(scanner.Text()), `\d+`)

	//Iterate over the duration values
	for i, duration := range durations {
		wins := 0
		//Iterate over rangeOf values from 0 to duration
		for _, hold := range rangeOf(0, duration) {
			//If the record is less than or equal to the duration, print the record
			distance := hold * (duration - hold)
			if distance > records[i] {
				wins++
			}
		}
		waysToWin = waysToWin * wins
	}

	fmt.Println(waysToWin)

	//Handle errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
