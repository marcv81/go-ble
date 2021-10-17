package sensors

import (
	"reflect"
	"testing"

	"github.com/marcv81/go-ble/ble"
	"github.com/marcv81/go-ble/point"
)

func TestReadThermometer(t *testing.T) {
	testCases := []struct {
		in  ble.Advert
		out []point.NamedValue
		err error
	}{
		{
			in: ble.Advert{
				Type: 0x16,
				Data: []byte{
					0x1A, 0x18, 0x57, 0xED, 0x8F, 0x38, 0xC1, 0xA4,
					0xEC, 0x09, 0x03, 0x11, 0x2B, 0x0C, 0x64, 0x0A,
					0x04,
				},
			},
			out: []point.NamedValue{
				{Name: "temperature", Value: float32(25.4)},
				{Name: "humidity", Value: float32(43.55)},
				{Name: "battery_volt", Value: float32(3.115)},
				{Name: "battery_percent", Value: uint8(100)},
			},
			err: nil,
		},
	}
	for i, tc := range testCases {
		out, err := ReadThermometer(&tc.in)
		if err != tc.err {
			s := "test case %d error: expected %+v, actual %+v"
			t.Fatalf(s, i, tc.err, err)
		}
		if !reflect.DeepEqual(out, tc.out) {
			s := "test case %d output: expected %+v, actual %+v"
			t.Fatalf(s, i, tc.out, out)
		}
	}
}
