package cmd

import (
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"
)

var (
	ip       net.IP
	port     int
	delaySec int
	rootCmd  = &cobra.Command{
		Use:   "growatt-to-iot",
		Short: "Sends messages to AWS IOT",
		Run:   growattToIOT,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func growattToIOT(cmd *cobra.Command, args []string) {
	// TODO
}