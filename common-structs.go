package smarthome_common

import "net/http"

// TemperatureReading contains a temperature reading.
type TemperatureReading struct {
	Temperature float64 `json:"temperature"`
}

type HumidityReading struct {
	Humidity float64 `json:"humidity"`
}

// SensorConnectMessage is sent from a sensor to a hub so that the hub can make periodic calls for readings.
type SensorConnectMessage struct {
	SensorType DeviceType
	Model      SensorModel
	IP         string
	Port       string
}

type ActuatorConnectMessage struct {
	ActuatorType DeviceType
	IP           string
	Port         string
}

type PingMessage struct {
	Response *http.Response
	Address  string
	Body     []byte
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
