package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

var (
	wg    sync.WaitGroup
	mutex sync.Mutex
)

type Log struct {
	Port  string
	State string
}

func main() {
	fmt.Println("|| Beginning Scan ||")
	wg.Add(1)

	// goroutines (1)
	go fmt.Println(Scan("localhost", "tcp", 49152, &wg))

	wg.Wait()
	fmt.Scanln()
	fmt.Println("|| Scan Complete ||")
}

func PortCheck(hostname, protocol string, port int) Log {
	result := Log{Port: strconv.Itoa(port) + string("/") + protocol}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		result.State = "Closed"

		return result
	}

	conn.Close()
	result.State = "Open"

	return result
}

func Scan(hostname, protocol string, depth int, wg *sync.WaitGroup) []Log {
	var results []Log
	mutex.Lock()

	for i := 0; i <= depth; i++ {
		r := PortCheck(protocol, hostname, i)

		switch r.State {
		case "Open":
			results = append(results, r)
		default:
			continue
		}
	}

	mutex.Unlock()
	wg.Done()
	return results
}
