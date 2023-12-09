## Day 8

### Prompt

> mikeclarke: Implement the following with Golang.
>
> It seems like you're meant to use the left/right instructions to navigate the network. Perhaps if you have the camel follow the same instructions, you can escape the haunted wasteland!
>
> After examining the maps for a bit, two nodes stick out: AAA and ZZZ. You feel like AAA is where you are now, and you have to follow the left/right instructions until you reach ZZZ.
>
> This format defines each node of the network individually. For example:
>
> RL
>
> AAA = (BBB, CCC)
> BBB = (DDD, EEE)
> CCC = (ZZZ, GGG)
> DDD = (DDD, DDD)
> EEE = (EEE, EEE)
> GGG = (GGG, GGG)
> ZZZ = (ZZZ, ZZZ)
> Starting with AAA, you need to look up the next element based on the next left/right instruction in your input. In this example, start with AAA and go right (R) by choosing the right element of AAA, CCC. Then, L means to choose the left element of CCC, ZZZ. By following the left/right instructions, you reach ZZZ in 2 steps.
>
> Of course, you might not find ZZZ right away. If you run out of left/right instructions, repeat the whole sequence of instructions as necessary: RL really means RLRLRLRLRLRLRLRL... and so on. For example, here is a situation that takes 6 steps to reach ZZZ:
>
> LLR
>
> AAA = (BBB, BBB)
> BBB = (AAA, ZZZ)
> ZZZ = (ZZZ, ZZZ)
> Starting at AAA, follow the left/right instructions. How many steps are required to reach ZZZ?

GitHub Copilot: Here's how you can implement this in Go:

1. Define a map to represent the network, where each key is a node and its value is a slice representing the left and right nodes.
2. Initialize a variable to keep track of the current node, starting with 'AAA'.
3. Initialize a variable to keep track of the number of steps taken, starting with 0.
4. Create a loop that continues until the current node is 'ZZZ'.
5. Inside the loop, increment the steps variable by 1.
6. Determine the next node based on the current instruction (either 'L' or 'R'), and update the current node to this new node.
7. If the end of the instructions is reached, start over from the beginning.
8. Once the loop ends, return the number of steps taken.

Here's the Go code that implements this plan:

```go
package main

import "fmt"

func stepsToZZZ(instructions string, network map[string][2]string) int {
    currentNode := "AAA"
    steps := 0
    instructionIndex := 0

    for currentNode != "ZZZ" {
        steps++
        if instructions[instructionIndex] == 'L' {
            currentNode = network[currentNode][0]
        } else {  // instructions[instructionIndex] == 'R'
            currentNode = network[currentNode][1]
        }

        instructionIndex = (instructionIndex + 1) % len(instructions)
    }

    return steps
}

func main() {
    instructions := "LLR"
    network := map[string][2]string{
        "AAA": {"BBB", "BBB"},
        "BBB": {"AAA", "ZZZ"},
        "ZZZ": {"ZZZ", "ZZZ"},
    }
    fmt.Println(stepsToZZZ(instructions, network))  // Output: 6
}
```

This code defines a function `stepsToZZZ` that takes the instructions and network as arguments and returns the number of steps required to reach 'ZZZ'. The `main` function demonstrates how to call this function with a specific set of instructions and network.

> mikeclarke: Update the implementation to read the network from a file data.txt instead.

GitHub Copilot: To read the network from a file in Go, you can use the `os` and `bufio` packages to open the file and read it line by line. Each line should be split into the node and its left and right nodes, which can be added to the network map. Here's how you can do it:

```go
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
	instructions := "LLR"
	network := make(map[string][2]string)

	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
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
```

This code assumes that the file `data.txt` is in the format:

```
AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
```

Each line is split into the node and its left and right nodes, which are added to the network map. The `stepsToZZZ` function is then called with the instructions and the network map.

> mikeclarke: Update the implementation if the format of data.txt includes the instructions as the first line of the file, followed by an empty line, followed by the network lines.

GitHub Copilot: If the instructions are included as the first line of the file, you can simply read the first line before the loop that reads the network lines. Here's how you can do it:

```go
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
```

This code assumes that the file `data.txt` is in the format:

```
LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
```

The first line is read before the loop that reads the network lines, and the instructions are stored in the `instructions` variable. The second line is skipped because it's an empty line. The rest of the file is read as before.

> mikeclarke: Great! What if the map isn't for people - what if the map is for ghosts? Are ghosts even bound by the laws of spacetime? Only one way to find out.
>
> After examining the maps a bit longer, your attention is drawn to a curious fact: the number of nodes with names ending in A is equal to the number ending in Z! If you were a ghost, you'd probably just start at every node that ends with A and follow all of the paths at the same time until they all simultaneously end up at nodes that end with Z.
>
> For example:
>
> LR
>
> 11A = (11B, XXX)
> 11B = (XXX, 11Z)
> 11Z = (11B, XXX)
> 22A = (22B, XXX)
> 22B = (22C, 22C)
> 22C = (22Z, 22Z)
> 22Z = (22B, 22B)
> XXX = (XXX, XXX)
> Here, there are two starting nodes, 11A and 22A (because they both end with A). As you follow each left/right instruction, use that instruction to simultaneously navigate away from both nodes you're currently on. Repeat this process until all of the nodes you're currently on end with Z. (If only some of the nodes you're on end with Z, they act like any other node and you continue as normal.) In this example, you would proceed as follows:
>
> Step 0: You are at 11A and 22A.
> Step 1: You choose all of the left paths, leading you to 11B and 22B.
> Step 2: You choose all of the right paths, leading you to 11Z and 22C.
> Step 3: You choose all of the left paths, leading you to 11B and 22Z.
> Step 4: You choose all of the right paths, leading you to 11Z and 22B.
> Step 5: You choose all of the left paths, leading you to 11B and 22C.
> Step 6: You choose all of the right paths, leading you to 11Z and 22Z.
> So, in this example, you end up entirely on nodes that end in Z after 6 steps.
>
> Simultaneously start on every node that ends with A. How many steps does it take before you're only on nodes that end with Z?

GitHub Copilot: To solve this problem, you need to keep track of multiple current nodes instead of just one. You can use a slice to store the current nodes, and for each step, you update all nodes in the slice based on the current instruction. You also need to check if all current nodes end with 'Z' to determine when to stop the loop. Here's how you can do it:

```go
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

	steps := 0
	instructionIndex := 0

	for {
		steps++
		var nextNodes []string
		for _, node := range currentNodes {
			if instructions[instructionIndex] == 'L' {
				nextNodes = append(nextNodes, network[node][0])
			} else { // instructions[instructionIndex] == 'R'
				nextNodes = append(nextNodes, network[node][1])
			}
		}

		// Check if all nextNodes end with 'Z'
		allEndWithZ := true
		for _, node := range nextNodes {
			if !strings.HasSuffix(node, "Z") {
				allEndWithZ = false
				break
			}
		}
		if allEndWithZ {
			break
		}

		currentNodes = nextNodes
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
```

This code initializes `currentNodes` with all nodes that end with 'A'. In each step, it updates all nodes in `currentNodes` based on the current instruction. It then checks if all nodes in `nextNodes` end with 'Z', and if so, it breaks the loop. Otherwise, it updates `currentNodes` with `nextNodes` and continues to the next step.

> mikeclarke: This implementation is too slow. Propose an alternative implementation that returns the number of steps faster.

GitHub Copilot: The current implementation is slow because it's repeatedly visiting the same nodes. We can optimize this by using a breadth-first search (BFS) algorithm, which is a common approach for traversing or searching graph data structures. 

In this case, we can use BFS to find the shortest path from all nodes ending with 'A' to all nodes ending with 'Z'. We'll use a queue to keep track of the nodes to visit next, and a set to keep track of the nodes we've already visited. 

Here's the updated Go code:

```go
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
		for i := 0; i < queueSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if strings.HasSuffix(node, "Z") && len(queue) == 0 {
				return steps
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
```

This optimized implementation should return the number of steps much faster than the previous one.

> mikeclarke: This implementation returns the first path to a node that ends with the letter Z. Update this implementation to return the number of steps for all initial nodes that end with A.

GitHub Copilot: To ensure that all paths from nodes ending with 'A' reach nodes ending with 'Z', we need to modify the condition for returning the number of steps. Instead of returning the steps when the first node ending with 'Z' is found, we should continue the search until all nodes in the queue end with 'Z'. Here's the updated Go code:

```go
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
```

In this updated implementation, we check if all nodes in the queue end with 'Z' in each step. If they do, we return the number of steps. Otherwise, we continue to the next step.
