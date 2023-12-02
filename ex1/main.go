package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func getFirstNumber(str string) int {
	re := regexp.MustCompile(`^.*?(\d).*?$`)
	match := re.FindStringSubmatch(str)
	first, _ := strconv.Atoi(match[1])
	return first
}

// Generate a function that reverses a string
func reverseString(str string) string {
	//Convert the string to a slice of runes
	runes := []rune(str)

	//Reverse the slice
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		//Swap the elements
		runes[i], runes[j] = runes[j], runes[i]
	}

	//Return the reversed string
	return string(runes)
}

// Generate a function that replaces numeric words with numbers
func replaceNumericWords(str string) string {
	re := regexp.MustCompile(`one`)
	str = re.ReplaceAllString(str, "1")
	re = regexp.MustCompile(`two`)
	str = re.ReplaceAllString(str, "2")
	re = regexp.MustCompile(`three`)
	str = re.ReplaceAllString(str, "3")
	re = regexp.MustCompile(`four`)
	str = re.ReplaceAllString(str, "4")
	re = regexp.MustCompile(`five`)
	str = re.ReplaceAllString(str, "5")
	re = regexp.MustCompile(`six`)
	str = re.ReplaceAllString(str, "6")
	re = regexp.MustCompile(`seven`)
	str = re.ReplaceAllString(str, "7")
	re = regexp.MustCompile(`eight`)
	str = re.ReplaceAllString(str, "8")
	re = regexp.MustCompile(`nine`)
	str = re.ReplaceAllString(str, "9")
	return str
}

// Generate a function that matches a regex and returns the parts
func replaceFirst(str string) string {
	re := regexp.MustCompile(`^(.*?)(one|two|three|four|five|six|seven|eight|nine)(.*)$`)
	match := re.FindStringSubmatch(str)
	if len(match) == 0 {
		return str
	}
	return match[1] + replaceNumericWords(match[2]) + match[3]
}

func replaceLast(str string) string {
	re := regexp.MustCompile(`^(.*)(one|two|three|four|five|six|seven|eight|nine)(.*)$`)
	match := re.FindStringSubmatch(str)
	if len(match) == 0 {
		return str
	}
	return match[1] + replaceNumericWords(match[2]) + match[3]
}

func main() {
	//Open the file
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	calibration := 0

	//Read the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		first := getFirstNumber(replaceFirst(s)) * 10
		second := getFirstNumber(reverseString(replaceLast(s)))
		calibration += first + second
		fmt.Println(s, first+second)
	}

	//Handle errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(calibration)
}
