package main

import (
	"fmt"
	"net"
	"sort"
	"sync"
)

func worker(ports, results chan int, wg *sync.WaitGroup) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			// fmt.Println(p)
			wg.Done()
			continue
		}
		conn.Close()
		results <- p
		wg.Done()
	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int
	var wg sync.WaitGroup

	// Start a worker for each size? of ports
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results, &wg)
	}

	// Add ports via go routine so can start results loop
	go func() {
		for i := 1; i <= 100; i++ {
			wg.Add(1)
			ports <- i
		}
	}()

	// Read one result for each port added to ports
	for i := 0; i < 100; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	// Wait for all the work to finish; Clean up
	wg.Wait()
	close(ports)
	close(results)

	// Sort and print openports
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
