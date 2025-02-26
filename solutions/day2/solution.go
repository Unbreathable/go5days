package main

import (
	"bufio"
	"fmt"
	"maps"
	"math/rand"
	"os"
	"strings"
	"sync"
)

func main() {

	// Get the grammar.txt file
	file, err := os.Open("grammar.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)

	// Parse all the rules and set the start rule
	start := ""
	rules := map[string][]string{}
	finished := false
	for !finished {

		// Read the line and parse the format
		l, _, err := reader.ReadLine()
		if err != nil {
			finished = true
			continue
		}
		rule := strings.Split(string(l), "->")

		// Set the starting point (if not set yet)
		if start == "" {
			start = rule[0]
		}

		// Parse the other side of the arrow
		without := []string{}
		for _, out := range strings.Split(rule[1], "|") {
			without = append(without, strings.Trim(out, "\""))
		}
		rules[rule[0]] = without
	}

	// Collect using a mutex locked list
	resultMap := map[int]string{}
	mutex := &sync.Mutex{}

	// Start all the goroutines
	resultChan := make(chan string)
	for i := range 10000 {
		go func() {
			current := start

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
						found = true
						break
					}
				}

				if !found {

					mutex.Lock()
					resultMap[i] = current
					mutex.Unlock()

					// Send it to the channel
					resultChan <- current

					// Insert it into the list with the mutex in front of it

					break
				}
			}
		}()
	}

	// Collect all the results
	i := 10000
	for {
		result := <-resultChan
		i--
		fmt.Println("result came in:", result)
		if i == 0 {
			break
		}
	}

	fmt.Println("All workers are done!")
}
