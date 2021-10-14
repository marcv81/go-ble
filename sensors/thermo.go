package sensors

import (
	"encoding/binary"

	"github.com/marcv81/go-ble/ble"
)

// Parses advertisements from a Mi thermometer flashed with the pvvx firmware.
func ReadThermometer(advert *ble.Advert) (Fields, error) {
	if advert.Type != advertTypeServiceData {
		return nil, errUnexpectedAdvertType
	}
	if len(advert.Data) != 17 {
		return nil, errUnexpectedAdvertLength
	}
	uuid := binary.LittleEndian.Uint16(advert.Data[0:2])
	if uuid != uuidThermometer {
		return nil, errUnexpectedServiceDataUuid
	}

	fields := Fields{
		"temperature":     float32(binary.LittleEndian.Uint16(advert.Data[8:10])) / 100,
		"humidity":        float32(binary.LittleEndian.Uint16(advert.Data[10:12])) / 100,
		"battery_volt":    float32(binary.LittleEndian.Uint16(advert.Data[12:14])) / 1000,
		"battery_percent": advert.Data[14],
	}
	return fields, nil
}
