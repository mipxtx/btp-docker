package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	p := make([]byte, 4096)

	addr := net.UDPAddr{
		Port: 38001,
		IP:   net.ParseIP("0.0.0.0"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}

	i := 0

	conn, err := net.Dial("tcp", "btpd:38001")
	for {
		if err == nil {
			fmt.Println("btpd connected")
			break
		}
		fmt.Printf("Some error %v\n", err)
		i++
		if i == 10 {
			fmt.Println("connection fail. berak")
			return
		}
		time.Sleep(1 * time.Second)
		conn, err = net.Dial("tcp", "btpd:38001")
	}

	for {
		_, addr, err := ser.ReadFromUDP(p)

		if err != nil {
			fmt.Printf("Some error  %v\n", err)
			continue
		}

		s := strings.Trim(strings.Split(string(p), "\n")[0], "\r\n")
		fmt.Printf("income %s %s \n", addr, s)
		_, err = fmt.Fprintf(conn, s+"\r\n")
		if err != nil {
			fmt.Printf("Some error  %v\n", err)
			continue
		}
	}
}
