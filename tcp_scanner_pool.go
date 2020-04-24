package main

import (
	"fmt"
	"net"
	"sort"
)

// create a worker pool
func worker(ports, results chan int) {
	// continuously receive from ports channel
	for p := range ports {
		// create address string
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		// create tcp connection
		conn, err := net.Dial("tcp", address)

		// port is closed or filtered
		if err != nil {
			// send 0 as result on results channel
			results <- 0
			continue
		}

		// close the connection
		conn.Close()
		// send port number as result on results channel
		results <- p
	}
}

func main() {
	// create buffered channel capped at 100
	ports := make(chan int, 400)
	// create channel to communicate results from worker
	// to main thread
	results := make(chan int)
	// declare open ports array slice
	var openports []int

	// create the desired number of workers
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	// send ports to workers in separate go routine
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	// result gathering loop
	for i := 1; i <= 1024; i++ {
		// receive port results
		port := <-results
		// add open ports to openports array slice
		if port != 0 {
			openports = append(openports, port)
		}
	}

	// close the channels
	close(ports)
	close(results)

	// sort the openports array slice for readability
	sort.Ints(openports)
	// print the open ports
	for _, port := range openports {
		fmt.Printf("Port %d open\n", port)
	}
	fmt.Println("Scan complete.")
}
