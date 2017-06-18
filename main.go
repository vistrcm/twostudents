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

	iteration := 0                            // le'ts count iterations.
	divider := 1000000                        // how offten to report results
	hits := 0                                 // increase if both students suggested the same
	pre_probability := 10.0                   // store previously calculated probability for stopping condition
	epsilon := math.Nextafter(1.0, 2.0) - 1.0 // calculate machine epsilon

	for {
		iteration++
		if <-c1 == <-c2 {
			hits++
		}

		if iteration%divider == 0 {
			probability := float64(hits) / float64(iteration)
			fmt.Printf("iteration: %d. Probability: %.16f.\n", iteration, probability)

			// stop condition
			if math.Abs(probability-pre_probability) <= epsilon {
				fmt.Printf("looks like probability does not changing a lot. %.16f, -> %.16f\n", pre_probability, pre_probability)
				break
			}
			pre_probability = probability

		}
	}
}
