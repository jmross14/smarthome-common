package smarthome_common

// DeviceType is meant to distinguish the type of sensor.
type DeviceType int

const (
	// TemperatureSensor distinguishes that a device is a temperature sensor.
	TemperatureSensor DeviceType = iota
	// GarageOpener distinguishes that a device is a garage opener.
	GarageOpener
)

// SensorModel is meant to distinguish the model of an Item.
type SensorModel string

const (
	// DHT22 distinguishes a DHT-22 Humidity/Temperature sensor.
	DHT22 SensorModel = "DHT-22"
)
