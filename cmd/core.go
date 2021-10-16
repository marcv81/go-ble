package main

import (
	"github.com/marcv81/go-ble/ble"
	"github.com/marcv81/go-ble/point"
	"github.com/marcv81/go-ble/sensors"
)

// BLE advertisement processor.
type processor struct {
	read func(*ble.Advert) ([]point.NamedValue, error)
	tags []point.NamedValue
}

// Creates the processors from the devices configuration.
// Maps each device MAC address to the associated processor.
func indexProcessors(devices []DeviceConfig) map[string]processor {
	processors := map[string]processor{}
	for _, device := range devices {
		tags := make([]point.NamedValue, 0, len(device.Tags)+2)
		tags = append(tags, point.NamedValue{Name: "device", Value: device.Type})
		tags = append(tags, point.NamedValue{Name: "addr", Value: device.MacAddress})
		for name, value := range device.Tags {
			tags = append(tags, point.NamedValue{Name: name, Value: value})
		}

		var read func(*ble.Advert) ([]point.NamedValue, error)
		switch device.Type {
		case "mi_thermometer":
			read = sensors.ReadThermometer
		case "mi_scale":
			read = sensors.ReadScale
		}

		processors[device.MacAddress] = processor{
			read: read,
			tags: tags,
		}
	}
	return processors
}

// Attempts to convert BLE scan data into a data point.
// If successful, calls a function on the data point.
func process(processors map[string]processor, info ble.DeviceInfo, callback func(point.Point)) {
	if _, ok := processors[info.MacAddress]; !ok {
		return
	}
	processor := processors[info.MacAddress]
	for _, advert := range info.Adverts {
		fields, err := processor.read(&advert)
		if err != nil {
			continue
		}
		fields = append(fields, point.NamedValue{Name: "rssi", Value: info.Rssi})

		callback(point.Point{
			Measurement: "bluetooth",
			Fields:      fields,
			Tags:        processor.tags,
		})
	}
}
