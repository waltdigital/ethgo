package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	// scanning too quickly
	//for i := 1; i <= 1024; i++ {
	//	go func(j int) {
	//		address := fmt.Sprintf("192.168.56.101:%d", j)
	//		conn, err := net.Dial("tcp", address)
	//		if err != nil {
	//			return
	//		}
	//		err = conn.Close()
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//		fmt.Println("%d open\n", j)
	//	}(i)
	//}

	host := "scanme.nmap.org"
	var wg sync.WaitGroup
	for i := 1; i < 1024; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", host, j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			err = conn.Close()
			if err != nil {
				fmt.Println("[-] Error closing connection!")
				return
			}
			fmt.Printf("[+] %d open\n", j)
		}(i)
	}
	wg.Wait()
}
