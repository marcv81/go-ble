package sensors

import (
	"reflect"
	"testing"

	"github.com/marcv81/go-ble/ble"
)

// Both weight and impedance are available.
func TestReadScale1(t *testing.T) {
	input := &ble.Advert{
		Type: 0x16,
		Data: []byte{
			0x1B, 0x18, 0x02, 0x26, 0xB2, 0x07, 0x01, 0x01,
			0x00, 0x36, 0x2B, 0xBA, 0x01, 0xEC, 0x31,
		},
	}
	expected := Fields{
		"weight":    float32(63.9),
		"impedance": uint16(442),
	}
	actual, err := ReadScale(input)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected: %v, actual: %v", expected, actual)
	}
}

// Weight is availble, impedance is not available.
func TestReadScale2(t *testing.T) {
	input := &ble.Advert{
		Type: 0x16,
		Data: []byte{
			0x1B, 0x18, 0x02, 0x24, 0xB2, 0x07, 0x01, 0x01,
			0x00, 0x36, 0x2B, 0xFE, 0xFF, 0xEC, 0x31,
		},
	}
	expected := Fields{
		"weight": float32(63.9),
	}
	actual, err := ReadScale(input)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected: %v, actual: %v", expected, actual)
	}
}

// Neither weight nor impedance is available.
func TestReadScale3(t *testing.T) {
	input := &ble.Advert{
		Type: 0x16,
		Data: []byte{
			0x1B, 0x18, 0x02, 0x84, 0xB2, 0x07, 0x01, 0x01,
			0x00, 0x36, 0x39, 0x00, 0x00, 0x14, 0x00,
		},
	}
	expected := Fields{}
	actual, err := ReadScale(input)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected: %v, actual: %v", expected, actual)
	}
}
