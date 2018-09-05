package portknock

import (
	"fmt"
	"net"
	"time"
)

func connectTo(host string, port int) {
	fmt.Printf("knocking on %s:%d..\n", host, port)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), time.Millisecond*10)
	if err != nil {
		// ignore error
		return
	}
	defer conn.Close()
}
