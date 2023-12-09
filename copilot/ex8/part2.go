package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func stepsToZZZ(instructions string, network map[string][2]string) int {
    // Initialize currentNodes with nodes that end with 'A'
    var currentNodes []string
    for node := range network {
        if strings.HasSuffix(node, "A") {
            currentNodes = append(currentNodes, node)
        }
    }

    visited := make(map[string]bool)
    queue := make([]string, len(currentNodes))
    copy(queue, currentNodes)

    steps := 0
    instructionIndex := 0

    for len(queue) > 0 {
        queueSize := len(queue)
        allEndWithZ := true

        for i := 0; i < queueSize; i++ {
            node := queue[0]
            queue = queue[1:]

            if !strings.HasSuffix(node, "Z") {
                allEndWithZ = false
            }

            nextNode := ""
            if instructions[instructionIndex] == 'L' {
                nextNode = network[node][0]
            } else { // instructions[instructionIndex] == 'R'
                nextNode = network[node][1]
            }

            if !visited[nextNode] {
                visited[nextNode] = true
                queue = append(queue, nextNode)
            }
        }

        if allEndWithZ {
            return steps
        }

        instructionIndex = (instructionIndex + 1) % len(instructions)
        steps++
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
