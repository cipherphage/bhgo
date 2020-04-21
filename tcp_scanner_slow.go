package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("starting main")
	for i := 1; i <= 100; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", address)

		// port is closed or filtered
		if err != nil {
			fmt.Printf("Port %d closed\n", i)
			continue
		}

		conn.Close()
		fmt.Printf("Port %d open\n", i)
	}
	fmt.Println("main ended")
}
