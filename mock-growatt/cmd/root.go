package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var (
	ip       net.IP
	port     int
	delaySec int
	rootCmd  = &cobra.Command{
		Use:   "mock-growatt",
		Short: "Generates mock growatt messages over a socket connection",
		Run:   mockGrowatt,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func mockGrowatt(cmd *cobra.Command, args []string) {
	log.Println("Starting mock growatt")
	log.Printf("IP = %v, Port = %v, DelaySecs = %v", ip, port, delaySec)

	messages, err := loadMessages()
	if err != nil {
		log.Fatalf("Error loading messages: %v", err.Error())
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip.String(), port))
	if err != nil {
		log.Fatalf("Error on dial: %v", err.Error())
	}
	defer conn.Close()

	for {
		err := emitMessages(conn, messages)
		if err != nil {
			log.Fatalf("Error emitting messages: %v", err.Error())
		}
	}
}

func emitMessages(conn net.Conn, messages []string) error {
	for _, message := range messages {
		count, err := conn.Write([]byte(message))
		if err != nil {
			return err
		}
		log.Printf("Mock service sent [%d] bytes.", count)
		time.Sleep(time.Duration(delaySec) * time.Second)
	}

	return nil
}

func loadMessages() (messages []string, err error) {
	log.Println("Loading messages")
	messageFiles, err := filepath.Glob("./messages/*.txt")
	if err != nil {
		return nil, err
	}

	if messageFiles == nil {
		return nil, errors.New("No files match ./messages/*.txt")
	}

	for _, messageFile := range messageFiles {
		file, err := os.Open(messageFile)
		if err != nil {
			return nil, err
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			messages = append(messages, scanner.Text())
		}
		file.Close()
	}

	if len(messages) == 0 {
		return nil, errors.New("No messages found")
	}

	return
}

func init() {
	// Default is localhost
	defaultIP := net.IPv4(127, 0, 0, 1)
	rootCmd.Flags().IPVar(&ip, "ip", defaultIP, "IP Address to send mock messages.")

	// Default to Growatt port
	defaultPort := 5279
	rootCmd.Flags().IntVar(&port, "port", defaultPort, "The port of the address.")

	defaultDelaySecs := 1
	rootCmd.Flags().IntVar(&delaySec, "delaySec", defaultDelaySecs, "How many seconds to pause between messages.")
}