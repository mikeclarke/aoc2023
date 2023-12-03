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
	matches := re.FindAllStringSubmatch(str, -1)
	result := make([]int, 0, len(matches))
	for _, match := range matches {
		value, _ := strconv.Atoi(match[1])
		result = append(result, value)
	}
	return result
}

// Generate a function that maps sum over string of ints
func sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// Generate a function that maps sum over string of ints
func compare(nums []int, val int) bool {
	for _, num := range nums {
		if num > val {
			return false
		}
	}
	return true
}

// Generate a function that reduces an array of ints picking the max value
func max(nums []int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
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
		gameId := sum(parseValues(s, `Game\s+(\d+):`))
		// blue := compare(parseValues(s, `(\d+)\s+blue`), 14)
		// green := compare(parseValues(s, `(\d+)\s+green`), 13)
		// red := compare(parseValues(s, `(\d+)\s+red`), 12)
		blue := max(parseValues(s, `(\d+)\s+blue`))
		green := max(parseValues(s, `(\d+)\s+green`))
		red := max(parseValues(s, `(\d+)\s+red`))
		fmt.Println(s, gameId, blue, green, red)
		total += blue * green * red
	}

	//Handle errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(total)
}
