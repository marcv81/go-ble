package cc2540

import (
	"errors"
)

var errUnexpectedPacketType = errors.New("unexpected packet type")

const (
	packetTypeCmd   = 1
	packetTypeEvent = 4
)

// Vendor-specific command opcodes.
const (
	opcodeInit     = 0xFE00
	opcodeDiscover = 0xFE04
)

// Vendor-specific event code.
const eventCodeVendor = 0xFF

// Vendor-specific event opcodes.
const (
	opcodeInitDone     = 0x0600
	opcodeDiscoverDone = 0x0601
	opcodeDeviceInfo   = 0x060D
)

// HCI "GAP device init" command.
var initCommand = &command{
	opcode: opcodeInit,
	params: []byte{
		// Profile role, 8 = central
		8,
		// Maximum scan responses
		10,
		// Identity resolving key, 0 = generated
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		// Connection signature resolving key, 0 = generated
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		// Initial signature counter
		1, 0, 0, 0,
	},
}

// HCI "GAP device discovery request" command
var discoverCommand = &command{
	opcode: opcodeDiscover,
	params: []byte{
		// Mode, 3 = scan for all devices
		3,
		// Active scan, 0 = off
		0,
		// White list, 0 = do not use the white list
		0,
	},
}
