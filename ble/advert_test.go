package ble

import (
	"reflect"
	"testing"
)

func TestCreateAdverts(t *testing.T) {
	input := []byte{
		1, 123,
		4, 100, 1, 2, 3,
	}
	expected := []Advert{
		{Type: 123, Data: []byte{}},
		{Type: 100, Data: []byte{1, 2, 3}},
	}
	actual, err := CreateAdverts(input)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected: %v, actual: %v", expected, actual)
	}
}

func TestCreateAdvertsErrors(t *testing.T) {
	inputs := [][]byte{
		{0},          // Length = 0
		{1},          // Length = 1 but no data
		{4, 1, 2, 3}, // Length = 4 but not enough data
		{1, 0, 1},    // Valid advertisement, then length = 1 but no data
	}
	for _, input := range inputs {
		_, err := CreateAdverts(input)
		if err != errInvalidAdvertLength {
			t.Fatalf("expected: %s, actual: %s", errInvalidAdvertLength, err)
		}
	}
}
