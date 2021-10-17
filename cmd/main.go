package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jacobsa/go-serial/serial"
	"github.com/marcv81/go-ble/ble"
	"github.com/marcv81/go-ble/cc2540"
	"github.com/marcv81/go-ble/point"
)

var errParameter = errors.New("mandatory first parameter: config filename")

func main() {
	// Read the command line parameters.
	if len(os.Args) < 2 {
		log.Fatal(errParameter)
	}
	configFile := os.Args[1]

	// Read the config file.
	config, err := readAppConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}

	// Open the serial port to communicate with the CC2540.
	options := serial.OpenOptions{
		PortName:        config.Port,
		BaudRate:        115200,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 1,
	}
	readWriter, err := serial.Open(options)
	if err != nil {
		log.Fatal(err)
	}

	// Scan for BLE devices and process the advertisements.
	router := newRouter(config.Devices)
	scanner := cc2540.Scanner{ReadWriter: readWriter}
	scanner.Scan(func(info *ble.DeviceInfo) {
		router.route(info, func(p *point.Point) {
			s, err := p.String()
			if err != nil {
				return
			}
			fmt.Println(s)
		})
	})
}
