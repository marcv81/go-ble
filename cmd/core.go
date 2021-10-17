package main

import (
	"github.com/marcv81/go-ble/ble"
	"github.com/marcv81/go-ble/point"
	"github.com/marcv81/go-ble/sensors"
)

// BLE device information router.
type router struct {
	processors map[string]processor
}

// Creates a router from device configurations.
func newRouter(devices []DeviceConfig) *router {
	processors := make(map[string]processor, len(devices))
	for _, device := range devices {
		processors[device.MacAddress] = *newProcessor(&device)
	}
	return &router{processors: processors}
}

// Routes the BLE advertisements in a BLE device information
// to the appropriate processor.
func (o *router) route(info *ble.DeviceInfo, cb func(*point.Point)) {
	if _, ok := o.processors[info.MacAddress]; !ok {
		return
	}
	processor := o.processors[info.MacAddress]
	for _, advert := range info.Adverts {
		processor.process(&advert, func(p *point.Point) {
			p.Fields = append(p.Fields, point.NamedValue{
				Name: "rssi", Value: info.Rssi,
			})
			cb(p)
		})
	}
}

// BLE advertisement processor.
type processor struct {
	read func(*ble.Advert) ([]point.NamedValue, error)
	tags []point.NamedValue
}

// Creates a processor from a device configuration.
func newProcessor(device *DeviceConfig) *processor {
	// Read function.
	var read func(*ble.Advert) ([]point.NamedValue, error)
	switch device.Type {
	case "mi_thermometer":
		read = sensors.ReadThermometer
	case "mi_scale":
		read = sensors.ReadScale
	}
	// Tags.
	tags := make([]point.NamedValue, 0, len(device.Tags)+2)
	tags = append(tags, point.NamedValue{Name: "device", Value: device.Type})
	tags = append(tags, point.NamedValue{Name: "addr", Value: device.MacAddress})
	for _, tag := range device.Tags {
		tags = append(tags, point.NamedValue{Name: tag.Name, Value: tag.Value})
	}
	return &processor{read: read, tags: tags}
}

// Attempts to create and process a data point from a BLE advertisement.
func (o *processor) process(advert *ble.Advert, cb func(*point.Point)) {
	fields, err := o.read(advert)
	if err != nil {
		return
	}
	cb(&point.Point{
		Measurement: "bluetooth",
		Fields:      fields,
		Tags:        o.tags,
	})
}
