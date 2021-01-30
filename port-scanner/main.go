package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			// fmt.Println(p)
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	// How many workers to run? Must be >= 1
	const NumWorkers int = 50
	// Define range of ports to scan
	const MinPort int = 1
	const MaxPort int = 100
	// Set size of ports channel
	const MaxChan int = 10

	ports := make(chan int, MaxChan)
	results := make(chan int)
	var openports []int

	// Start a worker for each size? of ports
	for i := 0; i < NumWorkers; i++ {
		go worker(ports, results)
	}

	// Add ports via go routine so can start results loop
	go func() {
		for i := MinPort; i <= MaxPort; i++ {
			ports <- i
		}
	}()

	// Read one result for each port added to ports
	for i := MinPort; i <= MaxPort; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	// Wait for all the work to finish; Clean up
	close(ports)
	close(results)

	// Sort and print openports
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
