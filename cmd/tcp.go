/*
Copyright © 2022 samjtro
*/

package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var tcpCmd = &cobra.Command{
	Use:   "tcp",
	Short: "TCP Scan",
	Long:  `Usage: scan tcp [depth: int]`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("|| Beginning Scan ||")
		depth, err := strconv.Atoi(args[0])

		if err != nil {
			log.Fatalf("ERR: strconv.Atoi - %s", err.Error())
		}

		// goroutines (1)
		go fmt.Println(Scan("localhost", "tcp", depth))

		fmt.Println("|| Scan Complete ||")
	},
}

func init() {
	rootCmd.AddCommand(tcpCmd)
}
