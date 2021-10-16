package point

import (
	"errors"
	"fmt"
	"strings"
)

var (
	errInvalidMeasurement = errors.New("invalid measurement")
	errInvalidName        = errors.New("invalid name")
	errNoFields           = errors.New("at least 1 field is required")
)

// Name and value pair.
type NamedValue struct {
	Name  string
	Value interface{}
}

// Data point.
type Point struct {
	Measurement string
	Fields      []NamedValue
	Tags        []NamedValue
}

// Converts a point to a string in InfluxDB line format.
func (o *Point) String() (string, error) {
	if len(o.Measurement) == 0 {
		return "", errInvalidMeasurement
	}
	if len(o.Fields) == 0 {
		return "", errNoFields
	}

	// Measurement and tags
	list1 := make([]string, 0, len(o.Tags)+1)
	list1 = append(list1, o.Measurement)
	for _, nv := range o.Tags {
		if len(nv.Name) == 0 {
			return "", errInvalidName
		}
		element := fmt.Sprintf("%s=%v", nv.Name, nv.Value)
		list1 = append(list1, element)
	}

	// Fields
	list2 := make([]string, 0, len(o.Fields))
	for _, nv := range o.Fields {
		if len(nv.Name) == 0 {
			return "", errInvalidName
		}
		element := fmt.Sprintf("%s=%v", nv.Name, nv.Value)
		list2 = append(list2, element)
	}

	s := fmt.Sprintf(
		"%s %s",
		strings.Join(list1, ","),
		strings.Join(list2, ","),
	)
	return s, nil
}
