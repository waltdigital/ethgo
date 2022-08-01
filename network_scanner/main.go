package main

import (
	"fmt"
	"net"
	"sync"
)

func scanPort(port int, wg *sync.WaitGroup) {
	defer wg.Done()
	IP := "192.168.56.101"
	address := fmt.Sprintf(IP+":%d", port)
	connection, err := net.Dial("tcp", address)
	if err == nil {
		fmt.Printf("[+] Connection established... PORT %v\n", port)
	} else {
		fmt.Printf("[-] No connection established... PORT %v\n", port)
		return
	}
	connection.Close()
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i < 100; i++ {
		wg.Add(1)
		go scanPort(i, &wg)
	}
	wg.Wait()
	fmt.Println("[+] Execution completed.")
}
