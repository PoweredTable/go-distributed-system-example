package common

import (
	"encoding/json"
	"fmt"
)

// PacketType definition
type PacketType int

const (
	REQUEST PacketType = iota
	RESPONSE
)

// Packet struct definition
type Packet struct {
	ID      string                 `json:"id"`
	Type    PacketType             `json:"type"`
	Message string                 `json:"message"`
	Body    map[string]interface{} `json:"body"`
}

// ToBytes converts the Packet to a JSON-encoded byte slice.
func (p *Packet) ToBytes() ([]byte, error) {
	packetBytes, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("error marshalling packet: %w", err)
	}
	return packetBytes, nil
}

// FromBytes converts a JSON-encoded byte slice to a Packet.
func PacketFromBytes(data []byte) (*Packet, error) {
	var packet Packet
	err := json.Unmarshal(data, &packet)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling packet: %w", err)
	}
	return &packet, nil
}