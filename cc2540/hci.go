package cc2540

import (
	"encoding/binary"
	"io"
)

// HCI command.
type command struct {
	opcode uint16
	params []byte
}

// HCI event.
type event struct {
	eventCode byte
	data      []byte
}

// Sends a HCI command to the HostTest firmware.
func (o *command) send(writer io.Writer) error {
	data := make([]byte, 4+len(o.params))
	data[0] = packetTypeCmd
	binary.LittleEndian.PutUint16(data[1:3], o.opcode)
	data[3] = byte(len(o.params))
	copy(data[4:], o.params)
	_, err := writer.Write(data)
	return err
}

// Receives an HCI event from the HostTest firmware.
func receiveEvent(reader io.Reader) (*event, error) {
	// Read the packet type.
	data := make([]byte, 1)
	_, err := io.ReadFull(reader, data)
	if err != nil {
		return nil, err
	}
	packetType := data[0]
	if packetType != packetTypeEvent {
		return nil, errUnexpectedPacketType
	}
	// Read the event code and data length.
	data = make([]byte, 2)
	_, err = io.ReadFull(reader, data)
	if err != nil {
		return nil, err
	}
	eventCode := data[0]
	length := int(data[1])
	// Read the event data.
	data = make([]byte, length)
	_, err = io.ReadFull(reader, data)
	if err != nil {
		return nil, err
	}
	return &event{
		eventCode: eventCode,
		data:      data,
	}, nil
}
