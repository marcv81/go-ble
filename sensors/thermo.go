package sensors

import (
	"encoding/binary"

	"github.com/marcv81/go-ble/ble"
	"github.com/marcv81/go-ble/point"
)

// Parses advertisements from a Mi thermometer flashed with the pvvx firmware.
func ReadThermometer(advert *ble.Advert) ([]point.NamedValue, error) {
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

	temperature := float32(binary.LittleEndian.Uint16(advert.Data[8:10])) / 100
	humidity := float32(binary.LittleEndian.Uint16(advert.Data[10:12])) / 100
	batteryVolt := float32(binary.LittleEndian.Uint16(advert.Data[12:14])) / 1000
	batteryPercent := advert.Data[14]
	fields := []point.NamedValue{
		{Name: "temperature", Value: temperature},
		{Name: "humidity", Value: humidity},
		{Name: "battery_volt", Value: batteryVolt},
		{Name: "battery_percent", Value: batteryPercent},
	}
	return fields, nil
}
