package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"os"

	"github.com/D-Heger/nimbus/raindrop"
)

func main() {
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "-send":
		sendCmd.Parse(os.Args[2:])
		if sendCmd.NArg() < 1 {
			fmt.Println("Usage: nimbus -send <message>")
			os.Exit(1)
		}
		message := sendCmd.Args()[0]
		for i := 1; i < sendCmd.NArg(); i++ {
			message += " " + sendCmd.Args()[i]
		}
		sendMessage(message)
	case "-version", "version":
		fmt.Printf("Nimbus Client %s (Protocol v%d)\n", raindrop.SoftwareVersion, raindrop.ProtocolVersion)
	case "-help", "help":
		printUsage()
	default:
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Printf("Nimbus Client %s (Protocol v%d)\n", raindrop.SoftwareVersion, raindrop.ProtocolVersion)
	fmt.Println("Usage:")
	fmt.Println("  nimbus -send <message>   Send a message to the server")
	fmt.Println("  nimbus -version          Show version information")
	fmt.Println("  nimbus -help             Show this help message")
}

func sendMessage(message string) {
	if message[len(message)-1] != '\n' {
		message += "\n"
	}

	config := &tls.Config{
		InsecureSkipVerify: true, // Skip certificate verification for self-signed cert
	}

	server := os.Getenv("NIMBUS_SERVER")
	if server == "" {
		server = "localhost:8080"
	}

	conn, err := tls.Dial("tcp", server, config)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	err = raindrop.WritePacket(conn, raindrop.CmdChunk, []byte(message))
	if err != nil {
		fmt.Println("Error sending message:", err)
		os.Exit(1)
	}

	fmt.Println("Message sent successfully")
}
