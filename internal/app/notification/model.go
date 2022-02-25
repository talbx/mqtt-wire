package notification

type InputData struct {
	DataPoints []DeviceData
}

type DeviceData struct {
	Message Message `json:"message"`
	Topic string `json:"topic"`
}

type Message struct {
	BatteryStatus int `json:"battery"`
}
