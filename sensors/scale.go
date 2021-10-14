package sensors

import (
	"encoding/binary"

	"github.com/marcv81/go-ble/ble"
)

// Parses advertisements from a Mi body composition scale.
func ReadScale(advert *ble.Advert) (Fields, error) {
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

	fields := Fields{}
	if weightReady && unitKilogram {
		fields["weight"] = float32(binary.LittleEndian.Uint16(advert.Data[13:15])) / 200
	}
	if impedanceReady {
		fields["impedance"] = binary.LittleEndian.Uint16(advert.Data[11:13])
	}
	return fields, nil
}
