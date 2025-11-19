package raindrop

import (
	"encoding/binary"
	"errors"
	"io"
)

// Protocol and Software versions
const (
	ProtocolVersion = 1
	SoftwareVersion = "v0.0.2" //TODO Update versioning scheme
)

// Command types
const (
	// CmdHello is the initial handshake packet.
	// Payload: HelloPayload
	CmdHello = 0x01

	// CmdAuth is used for authentication.
	// Payload: AuthPayload
	CmdAuth = 0x02

	// CmdList requests a list of files or returns a list.
	// Payload: TBD (e.g., path string or file list struct)
	CmdList = 0x03

	// CmdChunk represents a file data chunk.
	// Payload: ChunkPayload
	CmdChunk = 0x04

	// CmdAck acknowledges receipt of a packet or operation.
	// Payload: TBD (e.g., ID of acknowledged item)
	CmdAck = 0x05

	// CmdError reports an error.
	// Payload: ErrorPayload
	CmdError = 0x06

	// CmdUpload initiates a file upload.
	// Payload: TBD (e.g., file metadata)
	CmdUpload = 0x07

	// CmdDownload initiates a file download.
	// Payload: TBD (e.g., file path/ID)
	CmdDownload = 0x08

	// 0x09 - 0xFF are reserved for future use.
)

// HeaderSize is the size of the packet header in bytes (1 byte version + 1 byte type + 4 bytes length)
const HeaderSize = 6

// Packet represents a RainDrop protocol packet
type Packet struct {
	Version byte
	Type    byte
	Payload []byte
}

// ReadPacket reads a packet from the reader
func ReadPacket(r io.Reader) (*Packet, error) {
	header := make([]byte, HeaderSize)
	if _, err := io.ReadFull(r, header); err != nil {
		return nil, err
	}

	version := header[0]
	cmdType := header[1]
	payloadLen := binary.BigEndian.Uint32(header[2:])

	payload := make([]byte, payloadLen)
	if _, err := io.ReadFull(r, payload); err != nil {
		return nil, err
	}

	return &Packet{
		Version: version,
		Type:    cmdType,
		Payload: payload,
	}, nil
}

// WritePacket writes a packet to the writer
func WritePacket(w io.Writer, cmdType byte, payload []byte) error {
	if len(payload) > 4294967295 { // Max uint32
		return errors.New("payload too large")
	}

	header := make([]byte, HeaderSize)
	header[0] = ProtocolVersion
	header[1] = cmdType
	binary.BigEndian.PutUint32(header[2:], uint32(len(payload)))

	if _, err := w.Write(header); err != nil {
		return err
	}

	if len(payload) > 0 {
		if _, err := w.Write(payload); err != nil {
			return err
		}
	}

	return nil
}
