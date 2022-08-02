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

var udpCmd = &cobra.Command{
	Use:   "udp",
	Short: "UDP Scan",
	Long:  `Usage: scan udp [depth: int]`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("|| Beginning Scan ||")
		depth, err := strconv.Atoi(args[0])

		if err != nil {
			log.Fatalf("ERR: strconv.Atoi - %s", err.Error())
		}

		// goroutines (1)
		go fmt.Println(Scan("localhost", "udp", depth))

		fmt.Println("|| Scan Complete ||")
	},
}

func init() {
	rootCmd.AddCommand(udpCmd)
}
