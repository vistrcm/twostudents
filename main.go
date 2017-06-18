package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// student generates random number from 1 to 4 and sends back to channel
func student(c chan<- int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		c <- r.Intn(4)
	}
}

// main function
func main() {

	c1 := make(chan int) // channel for student 1
	c2 := make(chan int) // channel for student 2

	// "start" students
	go student(c1)
	go student(c2)

	iteration := 0              // le'ts count iterations.
	divider := 1000000          // how often to report results
	hits := 0                   // increase if both students suggested the same
	previousProbability := 10.0 // store previously calculated probability for stopping condition
	epsilon := 0.000001
	noChangeLimit := 3 // stop if probability was not changed noChangeLimit times
	noChangeFor := 0   // count amount of changes

	for {
		iteration++
		if <-c1 == <-c2 {
			hits++
		}

		if iteration%divider == 0 {
			probability := float64(hits) / float64(iteration)
			fmt.Printf("iteration: %d. Probability: %.6f.\n", iteration, probability)

			if math.Abs(probability-previousProbability) <= epsilon {
				noChangeFor++ // increase if probability does not changed
			} else {
				noChangeFor = 0 // reset if there is a change
			}

			// stop condition
			if noChangeFor >= noChangeLimit {
				fmt.Printf("looks like probability does not changing for about %d iterations.\n", noChangeFor*divider)
				break
			}
			previousProbability = probability

		}
	}
}
