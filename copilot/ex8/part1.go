package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func stepsToZZZ(instructions string, network map[string][2]string) int {
    currentNode := "AAA"
    steps := 0
    instructionIndex := 0

    for currentNode != "ZZZ" {
        steps++
        if instructions[instructionIndex] == 'L' {
            currentNode = network[currentNode][0]
        } else { // instructions[instructionIndex] == 'R'
            currentNode = network[currentNode][1]
        }

        instructionIndex = (instructionIndex + 1) % len(instructions)
    }

    return steps
}

func main() {
    network := make(map[string][2]string)

    file, err := os.Open("data.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    // Read the instructions from the first line
    scanner.Scan()
    instructions := scanner.Text()

    // Skip the empty line
    scanner.Scan()

    // Read the network lines
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, " = ")
        nodes := strings.Split(parts[1][1:len(parts[1])-1], ", ")
        network[parts[0]] = [2]string{nodes[0], nodes[1]}
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }

    fmt.Println(stepsToZZZ(instructions, network)) // Output: 6
}
