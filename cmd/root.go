/*
Copyright © 2022 samjtro
*/

package cmd

import (
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var (
	mutex sync.Mutex
)

type Log struct {
	Port  string
	State string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "scan",
	Short: "Simple Port Scanning",
	Long:  `Pure Go Port Scanning, & other Network Security Utilities`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
