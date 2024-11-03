package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	blockingExample()
	//goroutinesInParallel()
}

// Function simulating a task
func task(name string, delay time.Duration) {
	for i := 0; i <= 3; i++ {
		fmt.Printf("%s - step %d\n", name, i)
		time.Sleep(delay)
	}
}

func taskTest() {
	// Start two tasks as goroutines
	go task("Task 1", 1*time.Second)
	go task("Task 2", 2*time.Second)

	// Allow time for goroutines to complete
	time.Sleep(4 * time.Second)
	fmt.Println("All tasks finished")
}

func computeTask(name string) {
	sum := 0
	for i := 0; i < 1e9; i++ {
		sum += i
	}
	fmt.Printf("%s completed with sum: %d\n", name, sum)
}
func computeTaskTest() {
	// Set the number of CPU cores GO should use
	runtime.GOMAXPROCS(2)

	// Start two compute-intensive tasks as goroutines
	go computeTask("Task 1")
	go computeTask("Task 2")

	// Allow time for goroutines to complete
	time.Sleep(2 * time.Second)
	fmt.Println("All compute tasks finished")
}

// goroutines
/*
	Goroutine is a lightweight, independently running function managed by Go Runtime
	rather than the Operating System. When launching the goroutine, it creates a
	concurrent task which the Go runtime handles.
*/

func myFunc() {
	fmt.Println("Hello from goroutine!")
}

func myFuncTest() {
	go myFunc()
	time.Sleep(1 * time.Second)
	fmt.Println("This is the main function.")
}

func waitGroup() {
	var wg sync.WaitGroup
	wg.Add(1) // Add a task to the wait group

	go func() {
		defer wg.Done() // Mark this task as done
		fmt.Println("Hello from goroutine!")
	}()

	wg.Wait() // Wait until all tasks are done
	fmt.Println("Main function finished.")
}

func anonymousGoroutines() {
	go func() {
		fmt.Println("Hello, I'm anonymous goroutine!")
	}()

	// Wait for a bit to see the output
	time.Sleep(1 * time.Second)
}

func goroutinesInParallel() {
	runtime.GOMAXPROCS(4) // Use up to 4 CPU cores

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine 1: ", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for i := range 5 {
			fmt.Println("Goroutine 2: ", i)
			time.Sleep(150 * time.Millisecond)
		}
	}()

	// Wait for goroutines to complete
	time.Sleep(1 * time.Second)
}

func superReducedString(s string) string {
	count := 0
	result := ""
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && s[i] == s[i+1] {
			count++
			i += 1
			continue
		}
		result += string(s[i])
	}
	if count == 0 {
		if len(result) == 0 {
			return "Empty String"
		}
		return result
	}
	return superReducedString(result)
}

// channels

/*
	Channels are a way for go routines to communicate and
	synchronize without using shared memory, making it
	easier to manage concurrency in Go.

	Creating example: channel := make(chan int); channel of integers

	Sending and Receiving Data through channels
	sending and receiving are done by '<-' (arrow) operator

	channel <- 42; send value 42 to the channel

	To receive data from the channel, we use: =<- channel:
	value := <- channel
*/

func basicChannelCommunication() {
	channel := make(chan string)

	go func() {
		channel <- "Hello from goroutine!" // send a message
	}()

	message := <-channel // receive the message
	fmt.Println(message) // output: Hello from goroutine!
}

func blockingExample() {
	channel := make(chan int)

	go func() {
		channel <- 10 // send date; this will block until received
		fmt.Println("Data send!")
	}()
	value := <-channel              // receive data
	fmt.Println("Received:", value) // output: Received: 10
}

/*
							Buffered Channels
	By default channels in golang are unbuffered, each sending must wait for
a corresponding receiving.
*/
