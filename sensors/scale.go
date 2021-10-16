package sensors

import (
	"encoding/binary"

	"github.com/marcv81/go-ble/ble"
	"github.com/marcv81/go-ble/point"
)

// Parses advertisements from a Mi body composition scale.
func ReadScale(advert *ble.Advert) ([]point.NamedValue, error) {
	if advert.Type != advertTypeServiceData {
		return nil, errUnexpectedAdvertType
	}
	if len(advert.Data) != 15 {
		return nil, errUnexpectedAdvertLength
	}
	uuid := binary.LittleEndian.Uint16(advert.Data[0:2])
	if uuid != uuidScale {
		return nil, errUnexpectedServiceDataUuid
	}

	control := binary.LittleEndian.Uint16(advert.Data[2:4])
	unitPound := control&(1<<0) != 0
	unitCatty := control&(1<<14) != 0
	unitKilogram := !(unitPound || unitCatty)
	weightReady := control&(1<<13) != 0
	impedanceReady := control&(1<<9) != 0

	fields := []point.NamedValue{}
	if weightReady && unitKilogram {
		weight := float32(binary.LittleEndian.Uint16(advert.Data[13:15])) / 200
		fields = append(fields, point.NamedValue{Name: "weight", Value: weight})
	}
	if impedanceReady {
		impedance := binary.LittleEndian.Uint16(advert.Data[11:13])
		fields = append(fields, point.NamedValue{Name: "impedance", Value: impedance})
	}
	return fields, nil
}
