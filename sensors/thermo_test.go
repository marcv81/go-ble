package sensors

import (
	"reflect"
	"testing"

	"github.com/marcv81/go-ble/ble"
	"github.com/marcv81/go-ble/point"
)

func TestReadThermometer(t *testing.T) {
	input := &ble.Advert{
		Type: 0x16,
		Data: []byte{
			0x1A, 0x18, 0x57, 0xED, 0x8F, 0x38, 0xC1, 0xA4,
			0xEC, 0x09, 0x03, 0x11, 0x2B, 0x0C, 0x64, 0x0A,
			0x04,
		},
	}
	expected := []point.NamedValue{
		{Name: "temperature", Value: float32(25.4)},
		{Name: "humidity", Value: float32(43.55)},
		{Name: "battery_volt", Value: float32(3.115)},
		{Name: "battery_percent", Value: uint8(100)},
	}
	actual, err := ReadThermometer(input)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected: %v, actual: %v", expected, actual)
	}
}
