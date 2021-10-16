package sensors

import (
	"reflect"
	"testing"

	"github.com/marcv81/go-ble/ble"
	"github.com/marcv81/go-ble/point"
)

func TestReadScale(t *testing.T) {
	testCases := []struct {
		in  ble.Advert
		out []point.NamedValue
		err error
	}{
		// Both weight and impedance are available.
		{
			in: ble.Advert{
				Type: 0x16,
				Data: []byte{
					0x1B, 0x18, 0x02, 0x26, 0xB2, 0x07, 0x01, 0x01,
					0x00, 0x36, 0x2B, 0xBA, 0x01, 0xEC, 0x31,
				},
			},
			out: []point.NamedValue{
				{Name: "weight", Value: float32(63.9)},
				{Name: "impedance", Value: uint16(442)},
			},
			err: nil,
		},
		// Weight is availble, impedance is not available.
		{
			in: ble.Advert{
				Type: 0x16,
				Data: []byte{
					0x1B, 0x18, 0x02, 0x24, 0xB2, 0x07, 0x01, 0x01,
					0x00, 0x36, 0x2B, 0xFE, 0xFF, 0xEC, 0x31,
				},
			},
			out: []point.NamedValue{
				{Name: "weight", Value: float32(63.9)},
			},
			err: nil,
		},
		// Neither weight nor impedance is available.
		{
			in: ble.Advert{
				Type: 0x16,
				Data: []byte{
					0x1B, 0x18, 0x02, 0x84, 0xB2, 0x07, 0x01, 0x01,
					0x00, 0x36, 0x39, 0x00, 0x00, 0x14, 0x00,
				},
			},
			out: []point.NamedValue{},
			err: nil,
		},
	}
	for i, tc := range testCases {
		out, err := ReadScale(&tc.in)
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
