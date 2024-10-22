package main

import (
	"fmt"
	"net"
	"time"
)

func Check(domain string, port string) string {
	address := domain + ":" + port
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", address, timeout)
	status := ""

	if err != nil {
		status = fmt.Sprintf("[Down] %v is unreachable,\n Error: %v", domain, err)
	} else {
		status = fmt.Sprintf("[Up] %v is reachable\n From %v\n To: %v", domain, conn.LocalAddr(), conn.RemoteAddr())
		conn.Close()
	}
	return status
}
