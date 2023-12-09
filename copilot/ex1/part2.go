package main

import (
    "bufio"
    "log"
    "os"
    "strconv"
    "strings"
)

var digitWords = map[string]int{
    "one":   1,
    "two":   2,
    "three": 3,
    "four":  4,
    "five":  5,
    "six":   6,
    "seven": 7,
    "eight": 8,
    "nine":  9,
}

func extractDigits(line string) int {
    firstDigit := -1
    lastDigit := -1

    // Find first digit
    for i := 0; i < len(line); i++ {
        for word, digit := range digitWords {
            if strings.HasPrefix(line[i:], word) {
                firstDigit = digit
                break
            }
        }
        if firstDigit != -1 || (line[i] >= '1' && line[i] <= '9') {
            if firstDigit == -1 {
                firstDigit = int(line[i] - '0')
            }
            break
        }
    }

    // Find last digit
    for i := len(line) - 1; i >= 0; i-- {
        for word, digit := range digitWords {
            if strings.HasSuffix(line[:i+1], word) {
                lastDigit = digit
                break
            }
        }
        if lastDigit != -1 || (line[i] >= '1' && line[i] <= '9') {
            if lastDigit == -1 {
                lastDigit = int(line[i] - '0')
            }
            break
        }
    }

    if firstDigit != -1 && lastDigit != -1 {
        result, _ := strconv.Atoi(strconv.Itoa(firstDigit) + strconv.Itoa(lastDigit))
        return result
    }

    return 0
}

func main() {
    file, err := os.Open("data.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    total := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        total += extractDigits(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    log.Printf("Total calibration value: %d", total)
}
