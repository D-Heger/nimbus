package raindrop

import (
	"bytes"
	"encoding/gob"
)

// Payload defines an interface for protocol payloads
type Payload interface {
	Encode() ([]byte, error)
}

// EncodePayload is a generic helper to encode any struct using gob
func EncodePayload(v any) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// DecodePayload is a generic helper to decode data into a pointer v
func DecodePayload(data []byte, v any) error {
	dec := gob.NewDecoder(bytes.NewReader(data))
	return dec.Decode(v)
}

// ChunkPayload represents the payload for a CmdChunk packet
type ChunkPayload struct {
	FileID  string
	ChunkID uint64
	Offset  int64
	Data    []byte
	IsLast  bool
}

func (c *ChunkPayload) Encode() ([]byte, error) {
	return EncodePayload(c)
}

// ErrorPayload represents the payload for a CmdError packet
type ErrorPayload struct {
	Code    uint32
	Message string
}

func (e *ErrorPayload) Encode() ([]byte, error) {
	return EncodePayload(e)
}

// HelloPayload represents the payload for a CmdHello packet
type HelloPayload struct {
	ClientVersion string
	ProtocolVer   uint32
}

func (h *HelloPayload) Encode() ([]byte, error) {
	return EncodePayload(h)
}

// AuthPayload represents the payload for a CmdAuth packet
type AuthPayload struct {
	Token string
}

func (a *AuthPayload) Encode() ([]byte, error) {
	return EncodePayload(a)
}
