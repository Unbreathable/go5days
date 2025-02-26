package main

import (
	"bufio"
	"fmt"
	"maps"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {

	// Get the grammar.txt file
	file, err := os.Open("grammar.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	// Parse all the rules and set the start rule
	current := ""
	rules := map[string][]string{}
	for scanner.Scan() {

		// Read the line and parse the format
		rule := strings.Split(scanner.Text(), "->")

		// Set the starting point (if not set yet)
		if current == "" {
			current = rule[0]
		}

		// Parse the other side of the arrow
		without := []string{}
		for _, out := range strings.Split(rule[1], "|") {
			without = append(without, strings.Trim(out, "\""))
		}
		rules[rule[0]] = without
	}

	// Evaluate grammar randomly
	for {
		found := false
		for k := range maps.Keys(rules) {
			// Get random result
			results := rules[k]
			randomResult := results[rand.Intn(len(results))]

			// Apply pattern
			replaced := strings.Replace(current, k, randomResult, 1)
			if replaced != current {
				current = replaced
				fmt.Println(current)
				found = true
				break
			}
		}

		// Sleep for cool animation
		time.Sleep(10 * time.Millisecond)

		if !found {
			break
		}
	}
}
