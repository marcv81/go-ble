package ble

import (
	"errors"
)

var errInvalidAdvertLength = errors.New("invalid advertisement length")

// Device information obtained from a BLE scan.
type DeviceInfo struct {
	MacAddress string
	Rssi       int8
	Adverts    []Advert
}

// BLE advertisement.
type Advert struct {
	Type byte
	Data []byte
}

// Parses data for advertisements.
func CreateAdverts(data []byte) ([]Advert, error) {
	adverts := []Advert{}
	i := 0
	for i < len(data) {
		length := int(data[i])
		if (length < 1) || (i+1+length > len(data)) {
			return nil, errInvalidAdvertLength
		}
		adverts = append(adverts, Advert{
			Type: data[i+1],
			Data: data[i+2 : i+1+length],
		})
		i += length + 1
	}
	return adverts, nil
}
