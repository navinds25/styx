package utils

import (
	"fmt"
	"net"
	"time"
)

// PortKnock for sending a single request to a given host and port
func PortKnock(host string, port int) {
	fmt.Printf("knocking on %s:%d..\n", host, port)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), time.Millisecond*10)
	defer conn.Close()
	if err != nil {
		// ignore error
		return
	}
}
