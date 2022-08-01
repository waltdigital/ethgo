package main

import (
	"fmt"
	"net"
)

// main contains a basic port scanner
func main() {
	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err == nil {
		fmt.Println("[+] Connection successful.")
	}
}
