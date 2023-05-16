package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)
//Scan Function called in the main func which receives the IP address and prints out whether the port is open or closed
func scanPort(host string, port int, wg *sync.WaitGroup) {
	defer wg.Done()
	target := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", target, 2*time.Second)
	if err != nil {
		fmt.Printf("Port %d closed\n", port)
		return
	}
	conn.Close()
	fmt.Printf("Port %d open\n", port)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the host to scan: ")
	host, _ := reader.ReadString('\n')
	host = strings.TrimSpace(host)

	fmt.Printf("Scanning ports on host: %s\n", host)

	var wg sync.WaitGroup

	for port := 1; port <= 65535; port++ {
		wg.Add(1)
		go scanPort(host, port, &wg)
	}

	wg.Wait()
}
