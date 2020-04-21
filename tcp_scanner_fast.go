package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	// use a wait group for our goroutines
	var wg sync.WaitGroup

	// for loop for our port range
	for i := 1; i <= 1024; i++ {
		// increment the wait group
		wg.Add(1)

		// create goroutine
		go func(j int) {
			// tell the wait group its done when it is done
			defer wg.Done()
			// create address string
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			// create tcp connection
			conn, err := net.Dial("tcp", address)

			// port is closed or filtered
			if err != nil {
				return
			}

			conn.Close()
			fmt.Printf("Port %d open\n", j)
		}(i)
	}
	// block until all work is done
	wg.Wait()
	fmt.Println("Scan complete.")
}
