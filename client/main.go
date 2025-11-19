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
		fmt.Println("Usage: nimbus -send <message>")
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
	default:
		fmt.Println("Welcome to the Nimbus! Running v0.0.2")
		fmt.Println("Usage: nimbus -send <message>")
		os.Exit(1)
	}
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

	err = raindrop.WritePacket(conn, raindrop.CmdData, []byte(message))
	if err != nil {
		fmt.Println("Error sending message:", err)
		os.Exit(1)
	}

	fmt.Println("Message sent successfully")
}
