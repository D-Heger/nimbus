package raindrop

import (
	"encoding/binary"
	"errors"
	"io"
)

// Command types
const (
	CmdUpload   = 0x01
	CmdDownload = 0x02
	CmdData     = 0x03
	CmdSuccess  = 0x04
	CmdError    = 0x05
)

// HeaderSize is the size of the packet header in bytes (1 byte type + 4 bytes length)
const HeaderSize = 5

// Packet represents a RainDrop protocol packet
type Packet struct {
	Type    byte
	Payload []byte
}

// ReadPacket reads a packet from the reader
func ReadPacket(r io.Reader) (*Packet, error) {
	header := make([]byte, HeaderSize)
	if _, err := io.ReadFull(r, header); err != nil {
		return nil, err
	}

	cmdType := header[0]
	payloadLen := binary.BigEndian.Uint32(header[1:])

	payload := make([]byte, payloadLen)
	if _, err := io.ReadFull(r, payload); err != nil {
		return nil, err
	}

	return &Packet{
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
	header[0] = cmdType
	binary.BigEndian.PutUint32(header[1:], uint32(len(payload)))

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
