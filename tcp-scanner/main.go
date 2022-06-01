package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports, results chan int) {
	// Range to continuously receive from the ports channel, looping until the channel is closed
	for p := range ports {
		address := fmt.Sprintf("127.0.0.1:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// Port is closed or filtered
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	// Create a buffered channel with a capacity of 100
	// This will be used to create 100 workers
	ports := make(chan int, 100)

	// Use a separate channel to pass the result of the port scan back to the main thread
	results := make(chan int)

	// Use a slice to store the results to sort it later
	var openPorts []int

	fmt.Println("Start scanning...")

	// Start a desired number of workers, in this case it's 100
	for i := 0; i <= cap(ports); i++ {
		go worker(ports, results)
	}

	// Send all port values to the workers as a separate goroutine so the result-gathering loop below can start concurrently
	go func() {
		for i := 1; i <= 65535; i++ {
			ports <- i
		}
	}()

	// Because the number of ports sent and the number of results received are the same, the program can know when to close the channels and subsequently shut down the workers
	for i := 0; i < 65535; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	close(ports)
	close(results)

	sort.Ints(openPorts)
	for _, port := range openPorts {
		fmt.Printf("%d open\n", port)
	}

	fmt.Println("Scan finished.")
}
