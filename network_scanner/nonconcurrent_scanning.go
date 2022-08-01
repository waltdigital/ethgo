package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// testing for loop
	//for i := 1; i <= 1024; i++ {
	//	address := fmt.Sprintf("scanme.nmap.org:%d", i)
	//	fmt.Println(address)
	//}

	for i := 1; i <= 1024; i++ {
		host := "192.168.56.101"
		address := fmt.Sprintf("%s:%d", host, i)
		// fmt.Printf("Scanning host: %s\n", address)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// port is closed or filtered if err != nil
			continue
		}
		err = conn.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("PORT %d: open\n", i)
	}

}
