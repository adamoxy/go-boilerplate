package main

import (
	"fmt"
	"time"
)

func main() {
	// Use Case 1: Basic Channel Communication
	// Creating a channel to communicate integers
	messages := make(chan string)

	// Starting a new goroutine
	go func() {
		// Sending a message to the channel
		messages <- "Hello, World!"
	}()

	// Receiving the message from the channel
	msg := <-messages
	fmt.Println(msg) // Expected output: "Hello, World!"

	// Use Case 2: Buffered Channels
	// Creating a buffered channel with a capacity of 2
	bufferChannel := make(chan int, 2)

	// Sending values to the channel without a corresponding concurrent receive
	bufferChannel <- 1
	bufferChannel <- 2

	// Receiving the values from the channel
	fmt.Println(<-bufferChannel) // Expected output: 1
	fmt.Println(<-bufferChannel) // Expected output: 2

	// Use Case 3: Range and Close
	// Creating a channel to communicate integers
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Starting a worker goroutine
	go func() {
		for {
			// Receiving jobs from the channel
			j, more := <-jobs
			if more {
				fmt.Println("Received job", j)
			} else {
				// Notifying that all jobs are done
				fmt.Println("Received all jobs")
				done <- true
				return
			}
		}
	}()

	// Sending jobs to the worker
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("Sent job", j)
	}
	close(jobs) // Closing the jobs channel
	fmt.Println("Sent all jobs")

	// Waiting for the worker to finish
	<-done

	// Use Case 4: Select
	// Creating two channels
	c1 := make(chan string)
	c2 := make(chan string)

	// Starting two goroutines to send data to channels after a delay
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// Using select to wait on multiple channel operations
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received", msg1)
		case msg2 := <-c2:
			fmt.Println("Received", msg2)
		}
	}

	// Use Case 5: Timeouts
	// Creating a channel for simulating a long-running task
	result := make(chan string)

	// Simulating a task in a goroutine
	go func() {
		time.Sleep(2 * time.Second)
		result <- "result"
	}()

	// Implementing a timeout for the task
	select {
	case res := <-result:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}
