package sensors

import (
	"errors"
)

var (
	errUnexpectedAdvertType      = errors.New("unexpected advertisement type")
	errUnexpectedAdvertLength    = errors.New("unexpected advertisement length")
	errUnexpectedServiceDataUuid = errors.New("unexpected service data UUID")
)

// Advertisement types.
const (
	advertTypeServiceData = 0x16
)

// Service data UUIDs.
const (
	uuidThermometer = 0x181A
	uuidScale       = 0x181B
)
