package point

import (
	"testing"
)

func TestString(t *testing.T) {
	testCases := []struct {
		in  Point
		out string
		err error
	}{
		// Data point with 2 fields and 2 tags.
		{
			in: Point{
				Measurement: "abc",
				Fields: []NamedValue{
					{Name: "x", Value: 12.3},
					{Name: "y", Value: 456},
				},
				Tags: []NamedValue{
					{Name: "foo", Value: "bar"},
					{Name: "baz", Value: "qux"},
				},
			},
			out: "abc,foo=bar,baz=qux x=12.3,y=456",
			err: nil,
		},
		// Data point with 1 field and no tags.
		{
			in: Point{
				Measurement: "abc",
				Fields:      []NamedValue{{Name: "x", Value: 12.3}},
				Tags:        []NamedValue{},
			},
			out: "abc x=12.3",
			err: nil,
		},
		// Invalid data point with no measurement.
		{
			in: Point{
				Measurement: "",
				Fields:      []NamedValue{{Name: "x", Value: 12.3}},
				Tags:        []NamedValue{},
			},
			out: "",
			err: errInvalidMeasurement,
		},
		// Invalid data point with no fields.
		{
			in: Point{
				Measurement: "abc",
				Fields:      []NamedValue{},
				Tags:        []NamedValue{},
			},
			out: "",
			err: errNoFields,
		},
		// Invalid data point with an unnamed field.
		{
			in: Point{
				Measurement: "abc",
				Fields:      []NamedValue{{Name: "", Value: 12.3}},
				Tags:        []NamedValue{},
			},
			out: "",
			err: errInvalidName,
		},
		// Invalid data point with an unnamed tag.
		{
			in: Point{
				Measurement: "abc",
				Fields:      []NamedValue{{Name: "x", Value: 12.3}},
				Tags:        []NamedValue{{Name: "", Value: "bar"}},
			},
			out: "",
			err: errInvalidName,
		},
	}
	for i, tc := range testCases {
		out, err := tc.in.String()
		if err != tc.err {
			s := "test case %d error: expected %+v, actual %+v"
			t.Fatalf(s, i, tc.err, err)
		}
		if out != tc.out {
			s := "test case %d output: expected %+v, actual %+v"
			t.Fatalf(s, i, tc.out, out)
		}
	}
}
