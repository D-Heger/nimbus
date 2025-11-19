package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"path/filepath"

	"github.com/D-Heger/nimbus/raindrop"
)

func main() {
	certDir := os.Getenv("NIMBUS_CERT_DIR")
	if certDir == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home directory:", err)
			os.Exit(1)
		}
		certDir = filepath.Join(homeDir, ".nimbus", "certs")
	}

	certFile := filepath.Join(certDir, "cert.pem")
	keyFile := filepath.Join(certDir, "key.pem")

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		fmt.Println("Error loading certificates:", err)
		os.Exit(1)
	}

	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	listener, err := tls.Listen("tcp", ":8080", config)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Printf("Welcome to the Nimbus! Running %s (Protocol v%d)\n", raindrop.SoftwareVersion, raindrop.ProtocolVersion)
	fmt.Println("Loaded certificates from", certDir)
	fmt.Println("Nimbus Server listening on :8080 (TLS)")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	packet, err := raindrop.ReadPacket(conn)
	if err != nil {
		fmt.Println("Error reading packet:", err)
		return
	}

	if packet.Version != raindrop.ProtocolVersion {
		fmt.Printf("Received packet with incompatible version: %d (expected %d)\n", packet.Version, raindrop.ProtocolVersion)
		return
	}

	if packet.Type == raindrop.CmdChunk {
		fmt.Printf("Received: %s", string(packet.Payload))
	} else {
		fmt.Printf("Received unknown packet type: %d\n", packet.Type)
	}
}
