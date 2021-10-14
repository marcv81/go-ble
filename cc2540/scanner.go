package cc2540

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/marcv81/go-ble/ble"
)

// BLE scanner.
type Scanner struct {
	ReadWriter io.ReadWriter         // ReadWriter to communicate with the HostTest firmware
	Callback   func(*ble.DeviceInfo) // Callback to process BLE scan data
}

// Listens to BLE advertisements.
// Runs forever unless an error occurs.
func (o *Scanner) Scan() error {
	err := initCommand.send(o.ReadWriter)
	if err != nil {
		return err
	}
	for {
		ev, err := receiveEvent(o.ReadWriter)
		if err != nil {
			return err
		}
		err = o.handleEvent(ev)
		if err != nil {
			return err
		}
	}
}

// Handles the HCI events received as we listen to BLE advertisements.
func (o *Scanner) handleEvent(ev *event) error {
	if ev.eventCode != eventCodeVendor {
		return nil
	}
	opcode := binary.LittleEndian.Uint16(ev.data[0:2])
	switch opcode {
	// Discover devices after "GAP device init" completes.
	case opcodeInitDone:
		discoverCommand.send(o.ReadWriter)
	// Discover devices again after "GAP device discovery request" completes.
	case opcodeDiscoverDone:
		discoverCommand.send(o.ReadWriter)
	// Call the callback when we discover a device.
	// Silently ignore malformatted advertisements.
	case opcodeDeviceInfo:
		info, _ := createDeviceInfo(ev.data)
		o.Callback(info)
	}
	return nil
}

// Creates a DeviceInfo from a HCI "GAP device information" event data.
func createDeviceInfo(data []byte) (*ble.DeviceInfo, error) {
	length := data[12]
	adverts, err := ble.CreateAdverts(data[13 : 13+length])
	if err != nil {
		return nil, err
	}
	return &ble.DeviceInfo{
		MacAddress: formatMacAddress(data[5:11]),
		Rssi:       int8(data[11]),
		Adverts:    adverts,
	}, nil
}

// Converts a 6-bytes little-endian MAC address to a string.
func formatMacAddress(data []byte) string {
	return fmt.Sprintf(
		"%02x:%02x:%02x:%02x:%02x:%02x",
		data[5], data[4], data[3], data[2], data[1], data[0],
	)
}
