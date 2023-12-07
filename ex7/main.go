package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
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

func mapLetters(hand string) string {
	values := strings.Split(hand, "")
	mapped := make([]string, 0, len(values))
	for _, value := range values {
		switch value {
		case "A":
			value = "M"
		case "K":
			value = "L"
		case "Q":
			value = "K"
		case "J":
			value = "0"
		case "T":
			value = "I"
		case "9":
			value = "H"
		case "8":
			value = "G"
		case "7":
			value = "F"
		case "6":
			value = "E"
		case "5":
			value = "D"
		case "4":
			value = "C"
		case "3":
			value = "B"
		case "2":
			value = "A"
		}
		mapped = append(mapped, value)
	}
	// Return a string of array values joined by an empty string
	return strings.Join(mapped, "")
}

// Generate a function that scores a poker hand
func scoreHand(hand string) string {
	//Split the hand into cards
	values := strings.Split(hand, "")
	//Sort array of strings
	sort.Strings(values)

	//Count the number of J characters in the hand
	jCount := strings.Count(hand, "J")

	prefix := "00"

	//Check for a pair
	pair := false
	if values[0] == values[1] || values[1] == values[2] || values[2] == values[3] || values[3] == values[4] {
		pair = true
		prefix = "30"
	}

	//Check for two pairs
	twoPair := false
	if values[0] == values[1] && values[2] == values[3] {
		twoPair = true
		prefix = "40"
	} else if values[0] == values[1] && values[3] == values[4] {
		twoPair = true
		prefix = "40"
	} else if values[1] == values[2] && values[3] == values[4] {
		twoPair = true
		prefix = "40"
	}

	//Check for a three of a kind
	threeOfAKind := false
	if values[0] == values[1] && values[1] == values[2] {
		threeOfAKind = true
		prefix = "50"
	} else if values[1] == values[2] && values[2] == values[3] {
		threeOfAKind = true
		prefix = "50"
	} else if values[2] == values[3] && values[3] == values[4] {
		threeOfAKind = true
		prefix = "50"
	}

	//Check for a full house
	fullHouse := false
	if values[0] == values[1] && values[1] == values[2] && values[3] == values[4] {
		fullHouse = true
		prefix = "60"
	} else if values[0] == values[1] && values[2] == values[3] && values[3] == values[4] {
		fullHouse = true
		prefix = "60"
	}

	//Check for a four of a kind
	fourOfAKind := false
	if values[0] == values[1] && values[1] == values[2] && values[2] == values[3] {
		fourOfAKind = true
		prefix = "70"
	} else if values[1] == values[2] && values[2] == values[3] && values[3] == values[4] {
		fourOfAKind = true
		prefix = "70"
	}

	//Check for a five of a kind
	fiveOfAKind := false
	if values[0] == values[1] && values[1] == values[2] && values[2] == values[3] && values[3] == values[4] {
		fiveOfAKind = true
		prefix = "80"
	}

	//Check for a high card
	if !(fiveOfAKind || fullHouse || fourOfAKind || threeOfAKind || pair || twoPair) {
		prefix = "20"
	}

	// Apply wildcards
	if jCount == 1 {
		prefix = "30"
		if pair {
			prefix = "50"
		}
		if twoPair {
			prefix = "60"
		}
		if threeOfAKind {
			prefix = "70"
		}
		if fourOfAKind {
			prefix = "80"
		}
	} else if jCount == 2 {
		if pair {
			prefix = "50"
		}
		if twoPair {
			prefix = "70"
		}
		if threeOfAKind {
			prefix = "80"
		}
	} else if jCount == 3 {
		if threeOfAKind {
			prefix = "70"
		}
		if fullHouse {
			prefix = "80"
		}
	} else if jCount == 4 {
		prefix = "80"
	}

	return prefix + mapLetters(hand)
}

// Sort a map by its values
func sortMapByValue(basket map[int]string) []int {
	keys := make([]int, 0, len(basket))

	for key := range basket {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return basket[keys[i]] < basket[keys[j]]
	})

	return keys
}

func main() {
	//Open the file
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lineNum := 0
	bids := map[int]int{}
	hands := map[int]string{}
	outcomes := map[int]string{}

	//Read a file into a string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		split := strings.Split(input, " ")
		hands[lineNum] = split[0]
		outcomes[lineNum] = scoreHand(split[0])
		bids[lineNum] = parseValues(split[1], `\d+`)[0]
		lineNum++
	}

	sorted := sortMapByValue(outcomes)
	winnings := 0

	//Iterate over sorted keys
	for idx, key := range sorted {
		fmt.Println(idx, hands[key], outcomes[key])
		winnings += bids[key] * (idx + 1)
	}

	fmt.Println(winnings)

	//Handle errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
