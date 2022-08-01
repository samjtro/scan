package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	fmt.Println("|| Beginning Scan ||")
	fmt.Println(Scan("localhost"))
	fmt.Println(WideScan("localhost"))
}

type Log struct {
	Port    string
	State   string
	Service string
}

func PortCheck(protocol, hostname string, port int) Log {
	result := Log{Port: strconv.Itoa(port) + string("/") + protocol}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		result.State = "Closed"

		return result
	}

	defer conn.Close()

	result.State = "Open"

	return result
}

func Scan(hostname string) []Log {
	var results []Log

	for i := 0; i <= 1024; i++ {
		results = append(results, PortCheck("tcp", hostname, i))
	}

	for i := 0; i <= 1024; i++ {
		results = append(results, PortCheck("udp", hostname, i))
	}

	return results
}

func WideScan(hostname string) []Log {
	var results []Log

	for i := 0; i <= 49152; i++ {
		results = append(results, PortCheck("tcp", hostname, i))
	}

	for i := 0; i <= 49152; i++ {
		results = append(results, PortCheck("udp", hostname, i))
	}

	return results
}
