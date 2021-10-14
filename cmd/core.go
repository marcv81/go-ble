package main

import (
	"fmt"
	"strings"

	"github.com/marcv81/go-ble/ble"
	"github.com/marcv81/go-ble/sensors"
)

type Tags map[string]string

// BLE advertisement processor.
type processor struct {
	read func(*ble.Advert) (sensors.Fields, error)
	tags Tags
}

// Creates the processors from the devices configuration.
// Maps each device MAC address to the associated processor.
func indexProcessors(devices []DeviceConfig) map[string]processor {
	processors := map[string]processor{}
	for _, device := range devices {
		tags := make(Tags, len(device.Tags)+2)
		for key, value := range device.Tags {
			tags[key] = value
		}
		tags["device"] = device.Type
		tags["addr"] = device.MacAddress

		var read func(*ble.Advert) (sensors.Fields, error)
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

// Processes the scanned BLE devices information.
// Prints the discovered fields and tags in InfluxDB line format.
func process(processors map[string]processor, info ble.DeviceInfo) {
	if _, ok := processors[info.MacAddress]; !ok {
		return
	}
	processor := processors[info.MacAddress]
	for _, advert := range info.Adverts {
		fields, err := processor.read(&advert)
		if err != nil {
			continue
		}
		fields["rssi"] = info.Rssi
		fmt.Println(format(fields, processor.tags))
	}
}

// Formats fields and tags in InfluxDB line format.
func format(fields sensors.Fields, tags Tags) string {
	f := make([]string, 0, len(fields))
	for key, value := range fields {
		f = append(f, fmt.Sprintf("%s=%v", key, value))
	}
	t := make([]string, 0, len(tags))
	for key, value := range tags {
		t = append(t, fmt.Sprintf("%s=%s", key, value))
	}
	return fmt.Sprintf(
		"%s %s",
		strings.Join(append([]string{"bluetooth"}, t...), ","),
		strings.Join(f, ","),
	)
}
