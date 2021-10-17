package main

import (
	"reflect"
	"testing"

	"github.com/marcv81/go-ble/ble"
	"github.com/marcv81/go-ble/point"
)

func TestCore(t *testing.T) {
	router := newRouter([]DeviceConfig{
		{
			Type:       "mi_thermometer",
			MacAddress: "11:22:33:44:55:66",
			Tags:       []TagConfig{{Name: "location", Value: "bedroom"}},
		},
		{
			Type:       "mi_scale",
			MacAddress: "aa:bb:cc:dd:ee:ff",
			Tags:       []TagConfig{},
		},
	})
	testCases := []struct {
		in    ble.DeviceInfo
		out   point.Point
		calls int
	}{
		// Thermometer.
		{
			in: ble.DeviceInfo{
				MacAddress: "11:22:33:44:55:66",
				Rssi:       -49,
				Adverts: []ble.Advert{
					{
						Type: 0x16,
						Data: []byte{
							0x1A, 0x18, 0x57, 0xED, 0x8F, 0x38, 0xC1, 0xA4,
							0xEC, 0x09, 0x03, 0x11, 0x2B, 0x0C, 0x64, 0x0A,
							0x04,
						},
					},
				},
			},
			out: point.Point{
				Measurement: "bluetooth",
				Fields: []point.NamedValue{
					{Name: "temperature", Value: float32(25.4)},
					{Name: "humidity", Value: float32(43.55)},
					{Name: "battery_volt", Value: float32(3.115)},
					{Name: "battery_percent", Value: uint8(100)},
					{Name: "rssi", Value: int8(-49)},
				},
				Tags: []point.NamedValue{
					{Name: "device", Value: "mi_thermometer"},
					{Name: "addr", Value: "11:22:33:44:55:66"},
					{Name: "location", Value: "bedroom"},
				},
			},
			calls: 1,
		},
		// Scale.
		{
			in: ble.DeviceInfo{
				MacAddress: "aa:bb:cc:dd:ee:ff",
				Rssi:       -37,
				Adverts: []ble.Advert{
					{
						Type: 0x16,
						Data: []byte{
							0x1B, 0x18, 0x02, 0x26, 0xB2, 0x07, 0x01, 0x01,
							0x00, 0x36, 0x2B, 0xBA, 0x01, 0xEC, 0x31,
						},
					},
				},
			},
			out: point.Point{
				Measurement: "bluetooth",
				Fields: []point.NamedValue{
					{Name: "weight", Value: float32(63.9)},
					{Name: "impedance", Value: uint16(442)},
					{Name: "rssi", Value: int8(-37)},
				},
				Tags: []point.NamedValue{
					{Name: "device", Value: "mi_scale"},
					{Name: "addr", Value: "aa:bb:cc:dd:ee:ff"},
				},
			},
			calls: 1,
		},
	}
	for i, tc := range testCases {
		calls := 0
		router.route(&tc.in, func(out *point.Point) {
			calls += 1
			if !reflect.DeepEqual(*out, tc.out) {
				s := "test case %d output: expected %+v, actual %+v"
				t.Fatalf(s, i, tc.out, *out)
			}
		})
		if calls != tc.calls {
			s := "test case %d calls: expected %d, actual %d"
			t.Fatalf(s, i, tc.calls, calls)
		}
	}
}
