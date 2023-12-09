package main

import (
    "bufio"
    "log"
    "os"
    "strconv"
    "unicode"
)

func extractDigits(line string) int {
    firstDigit := -1
    lastDigit := -1

    for _, r := range line {
        if unicode.IsDigit(r) {
            if firstDigit == -1 {
                firstDigit = int(r - '0')
            }
            lastDigit = int(r - '0')
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
