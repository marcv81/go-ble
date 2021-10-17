package ble

import (
	"reflect"
	"testing"
)

func TestCreateAdverts(t *testing.T) {
	testCases := []struct {
		in  []byte
		out []Advert
		err error
	}{
		// Valid advertisements.
		{
			in: []byte{
				1, 123,
				4, 100, 1, 2, 3,
			},
			out: []Advert{
				{Type: 123, Data: []byte{}},
				{Type: 100, Data: []byte{1, 2, 3}},
			},
			err: nil,
		},
		// Length = 0.
		{
			in:  []byte{0},
			out: nil,
			err: errInvalidAdvertLength,
		},
		// Length = 1 but no data.
		{
			in:  []byte{1},
			out: nil,
			err: errInvalidAdvertLength,
		},
		// Length = 4 but not enough data.
		{
			in:  []byte{4, 1, 2, 3},
			out: nil,
			err: errInvalidAdvertLength,
		},
		// Valid advertisement, then length = 1 but no data.
		{
			in:  []byte{1, 0, 1},
			out: nil,
			err: errInvalidAdvertLength,
		},
	}
	for i, tc := range testCases {
		out, err := CreateAdverts(tc.in)
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
