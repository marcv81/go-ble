# Intro

This project outputs Bluetooth Low Energy (BLE) sensors data in InfluxDB line protocol format. It integrates nicely with the Telegraf execd input plugin.

Supported BLE adapters
- CC2540 USB dongles flashed with the HostTest firmware.

Supported sensors
- Xiaomi Thermometers running the pvvx firmware (https://github.com/pvvx/ATC_MiThermometer)
- Xiaomi Body Composition Scale 2

# Output

A sample follows.

    bluetooth,addr=a4:c1:38:d3:c0:6a,device=mi_thermometer temperature=26.51,humidity=82.69,battery_volt=3.005,battery_percent=89,rssi=-42
    bluetooth,addr=5c:ca:d3:ed:7f:27,device=mi_scale weight=64.7,impedance=403,rssi=-57

# Installation

Run the unit tests.

    go test ./...

Compile for Raspberry Pi.

    GOOS=linux GOARCH=arm64 go build -o go-ble ./cmd

Copy the program and the configuration file to the Raspberry Pi.

Add the Telegraf user to the dialout group.

    sudo usermod -a -G dialout telegraf

Create /etc/telegraf/telegraf.d/ble-stats.conf.

    [[inputs.execd]]
      command = ["/home/ubuntu/go-ble/go-ble", "/home/ubuntu/go-ble/config.yaml"]
      signal = "none"
      data_format = "influx"
